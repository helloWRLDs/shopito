package productrepository

import (
	"database/sql"
	productdomain "shopito/api/internal/domain/product"
)

type ProductRepository interface {
	GetAll() (*[]productdomain.Product, error)
	GetById(id int) (*productdomain.Product, error)
	Insert(product *productdomain.Product) (int, error)
	Delete(id int) error
	Update(id int, product *productdomain.Product) error
	Exist(id int) bool
}

type ProductRepositoryImpl struct {
	db *sql.DB
}

func New(db *sql.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		db: db,
	}
}

func (r *ProductRepositoryImpl) GetAll() (*[]productdomain.Product, error) {
	var ps []productdomain.Product
	stmt := `SELECT * FROM products`
	rows, err := r.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p productdomain.Product
		err := rows.Scan(&p.ID, &p.Name, &p.ImageURL, &p.Price, &p.Stock, &p.CategoryID)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	return &ps, nil
}

func (r *ProductRepositoryImpl) GetById(id int) (*productdomain.Product, error) {
	var p productdomain.Product
	stmt := `SELECT * FROM products WHERE id=$1`
	err := r.db.QueryRow(stmt, id).Scan(&p.ID, &p.Name, &p.ImageURL, &p.Price, &p.Stock, &p.CategoryID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepositoryImpl) Insert(product *productdomain.Product) (int, error) {
	var id int
	stmt := `INSERT INTO products(name, img_url, price, stock, category_id) VALUES($1, $2, $3, $4, $5) RETURNING id`
	err := r.db.QueryRow(stmt, product.Name, product.ImageURL, product.Price, product.Stock, product.CategoryID).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (r *ProductRepositoryImpl) Delete(id int) error {
	stmt := `DELETE FROM products WHERE id=$1`
	_, err := r.db.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryImpl) Update(id int, product *productdomain.Product) error {
	stmt := `UPDATE products SET name=$1, img_url=$2, price=$3, stock=$4, category_id=$5 WHERE id=$6`
	_, err := r.db.Exec(stmt, product.Name, product.ImageURL, product.Price, product.Stock, product.CategoryID, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryImpl) Exist(id int) bool {
	var exist bool
	stmt := `SELECT EXISTS(SELECT TRUE FROM products WHERE id=$1)`
	if err := r.db.QueryRow(stmt, id).Scan(&exist); err != nil {
		return false
	}
	return exist
}
