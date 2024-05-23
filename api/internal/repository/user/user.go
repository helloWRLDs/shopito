package userrepository

import (
	"database/sql"
	userdomain "shopito/api/internal/domain/user"
	"time"
)



type UserRepositoryImpl struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) GetByEmail(email string) (*userdomain.User, error) {
	var (
		u    userdomain.User
		stmt string = `SELECT * FROM users WHERE email=$1`
	)
	row := r.db.QueryRow(stmt, email)
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.IsAdmin, &u.IsVerified, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepositoryImpl) GetById(id int) (*userdomain.User, error) {
	var (
		u    userdomain.User
		stmt string = `SELECT * FROM users WHERE id=$1`
	)
	row := r.db.QueryRow(stmt, id)
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.IsAdmin, &u.IsVerified, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepositoryImpl) Insert(user *userdomain.User) (int, error) {
	var (
		id   int
		stmt = `INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id`
	)
	if err := r.db.QueryRow(stmt, user.Name, user.Email, user.Password).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (r *UserRepositoryImpl) ExistById(id int) bool {
	var (
		exist bool
		stmt  string = `SELECT EXISTS(SELECT TRUE FROM users WHERE id=$1)`
	)
	if err := r.db.QueryRow(stmt, id).Scan(&exist); err != nil {
		return false
	}
	return exist
}

func (r *UserRepositoryImpl) Update(id int, user *userdomain.User) error {
	stmt := `UPDATE users SET name=$1, email=$2, password=$3, is_verified=$4, is_admin=$5, updated_at=$6 WHERE id=$7`
	_, err := r.db.Exec(stmt, user.Name, user.Email, user.Password, user.IsVerified, user.IsAdmin, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) Delete(id int) error {
	stmt := `DELETE FROM users WHERE id=$1`
	_, err := r.db.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) IsVerified(id int) bool {
	var verified bool
	stmt := `SELECT is_verified FROM users WHERE id=$1`
	if err := r.db.QueryRow(stmt, id).Scan(&verified); err != nil {
		return false
	}
	return verified
}

func (r *UserRepositoryImpl) GetAll() (*[]userdomain.User, error) {
	var us []userdomain.User
	stmt := `SELECT * FROM users`
	rows, err := r.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u userdomain.User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.IsAdmin, &u.IsVerified, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	return &us, nil
}

func (r *UserRepositoryImpl) ExistByEmail(email string) bool {
	var exist bool
	stmt := `SELECT EXISTS(SELECT TRUE FROM users WHERE email=$1)`
	if err := r.db.QueryRow(stmt, email).Scan(&exist); err != nil {
		return false
	}
	return exist
}
