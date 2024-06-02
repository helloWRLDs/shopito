package categoryrepository

import (
	"database/sql"
	productproto "shopito/pkg/protobuf/products"
)

type Repository interface {
	Insert(category *productproto.Category) (int, error)
	Delete(id int) error
	Update(id int, category *productproto.Category) error
	Exist(id int) bool
	Get(id int) (*productproto.Category, error)
	List() ([]*productproto.Category, error)
}

type CategoryRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) Insert(category *productproto.Category) (int, error) {
	var id int
	stmt := `INSERT INTO categories(name) VALUES($1) RETURNING id`
	err := r.db.QueryRow(stmt, category.Name).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (r *CategoryRepository) Delete(id int) error {
	stmt := `DELETE from categories WHERE id=$1`
	_, err := r.db.Exec(stmt, id)
	return err
}

func (r *CategoryRepository) Update(id int, category *productproto.Category) error {
	stmt := `UPDATE categories SET name=$1 WHERE id=$2`
	_, err := r.db.Exec(stmt, category.Name, id)
	return err
}

func (r *CategoryRepository) Exist(id int) bool {
	var exist bool
	stmt := `SELECT EXISTS(SELECT TRUE FROM categories WHERE id=$1)`
	if err := r.db.QueryRow(stmt, id).Scan(&exist); err != nil {
		return false
	}
	return exist
}

func (r *CategoryRepository) Get(id int) (*productproto.Category, error) {
	var c productproto.Category
	stmt := `SELECT * FROM categories WHERE id=$1`
	err := r.db.QueryRow(stmt, id).Scan(&c.Id, &c.Name)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *CategoryRepository) List() ([]*productproto.Category, error) {
	var cs []*productproto.Category
	stmt := `SELECT * FROM categories`
	rows, err := r.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c productproto.Category
		if err := rows.Scan(&c.Id, &c.Name); err != nil {
			return nil, err
		}
		cs = append(cs, &c)
	}
	return cs, nil
}
