package productrepository

import (
	"database/sql"
	"shopito/pkg/protobuf/products"
	"strconv"
)

type Repository interface {
	Insert(product *productproto.Product) (int, error)
	Delete(id int) error
	Update(id int, product *productproto.Product) error
	Get(id int) (*productproto.Product, error)
	Exist(id int) bool
}

type ProductRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) Insert(product *productproto.Product) (int, error) {
	var id int
	stmt := `INSERT INTO products(name, img_url, price, stock, category_id) VALUES($1, $2, $3, $4, $5) RETURNING id`
	categoryId, err := strconv.Atoi(product.CategoryId)
	if err != nil {
		categoryId = 0
	}
	err = r.db.QueryRow(stmt, product.Name, product.ImgUrl, product.Price, product.Stock, categoryId).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (r *ProductRepository) Delete(id int) error {
	stmt := `DELETE FROM products WHERE id=$1`
	_, err := r.db.Exec(stmt, id)
	return err
}

func (r *ProductRepository) Update(id int, product *productproto.Product) error {
	categoryId, err := strconv.Atoi(product.GetCategoryId())
	if err != nil {
		categoryId = 0
	}
	stmt := `UPDATE products SET name=$1, img_url=$2, price=$3, stock=$4, category_id=$5 WHERE id=$6`
	_, err = r.db.Exec(stmt, product.Name, product.ImgUrl, product.Price, product.Stock, categoryId, id)
	return err
}

func (r *ProductRepository) Exist(id int) bool {
	var exist bool
	stmt := `SELECT EXISTS(SELECT TRUE FROM products WHERE id=$1)`
	if err := r.db.QueryRow(stmt, id).Scan(&exist); err != nil {
		return false
	}
	return exist
}

func (r *ProductRepository) Get(id int) (*productproto.Product, error) {
	var p productproto.Product
	stmt := `SELECT * FROM products WHERE id=$1`
	err := r.db.QueryRow(stmt, id).Scan(&p.Id, &p.Name, &p.ImgUrl, &p.Price, &p.Stock, &p.CategoryId)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
