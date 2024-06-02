package repository

import (
	"database/sql"
	"shopito/services/users/protobuf"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetByEmail(email string) (*protobuf.User, error) {
	var (
		u    protobuf.User
		stmt string = `SELECT id, name, email, password, is_admin, is_verified FROM users WHERE email=$1`
	)
	row := r.db.QueryRow(stmt, email)
	err := row.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.IsAdmin, &u.IsVerified)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetById(id int64) (*protobuf.User, error) {
	var (
		u    protobuf.User
		stmt string = `SELECT id, name, email, password, is_admin, is_verified FROM users WHERE id=$1`
	)
	row := r.db.QueryRow(stmt, id)
	err := row.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.IsAdmin, &u.IsVerified)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) Insert(user *protobuf.User) (int64, error) {
	var (
		id   int64
		stmt = `INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id`
	)
	if err := r.db.QueryRow(stmt, user.Name, user.Email, user.Password).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (r *UserRepository) ExistById(id int64) bool {
	var (
		exist bool
		stmt  string = `SELECT EXISTS(SELECT TRUE FROM users WHERE id=$1)`
	)
	if err := r.db.QueryRow(stmt, id).Scan(&exist); err != nil {
		return false
	}
	return exist
}

func (r *UserRepository) Update(id int64, user *protobuf.User) error {
	stmt := `UPDATE users SET name=$1, email=$2, password=$3, is_verified=$4, is_admin=$5, updated_at=$6 WHERE id=$7`
	_, err := r.db.Exec(stmt, user.Name, user.Email, user.Password, user.IsVerified, user.IsAdmin, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Delete(id int64) error {
	stmt := `DELETE FROM users WHERE id=$1`
	_, err := r.db.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) IsVerified(id int64) bool {
	var verified bool
	stmt := `SELECT is_verified FROM users WHERE id=$1`
	if err := r.db.QueryRow(stmt, id).Scan(&verified); err != nil {
		return false
	}
	return verified
}

func (r *UserRepository) GetAll() ([]*protobuf.User, error) {
	var us []*protobuf.User
	stmt := `SELECT id, name, email, password, is_admin, is_verified FROM users`
	rows, err := r.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u protobuf.User
		err := rows.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.IsAdmin, &u.IsVerified)
		if err != nil {
			return nil, err
		}
		us = append(us, &u)
	}
	return us, nil
}

func (r *UserRepository) ExistByEmail(email string) bool {
	var exist bool
	stmt := `SELECT EXISTS(SELECT TRUE FROM users WHERE email=$1)`
	if err := r.db.QueryRow(stmt, email).Scan(&exist); err != nil {
		return false
	}
	return exist
}
