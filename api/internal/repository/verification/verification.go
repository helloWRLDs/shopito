package verifyrepository

import (
	"database/sql"
)

type VerifyRepository interface {
	GetCodeByUser(user_id int) (string, error)
	Insert(user_id int, code string) error
	DeleteByUser(user_id int) error
	Update(user_id int, code string) error
}

type VerifyRepositoryImpl struct {
	db *sql.DB
}

func New(db *sql.DB) *VerifyRepositoryImpl {
	return &VerifyRepositoryImpl{
		db: db,
	}
}

func (r *VerifyRepositoryImpl) Insert(user_id int, code string) error {
	stmt := `INSERT INTO verification(code, user_id) VALUES($1, $2)`
	_, err := r.db.Exec(stmt, code, user_id)
	if err != nil {
		return err
	}
	return nil
}

func (r *VerifyRepositoryImpl) Update(user_id int, code string) error {
	stmt := `UPDATE verification SET code=$1 WHERE user_id=$2`
	_, err := r.db.Exec(stmt, code, user_id)
	if err != nil {
		return err
	}
	return nil
}

func (r *VerifyRepositoryImpl) DeleteByUser(user_id int) error {
	var stmt string = `DELETE FROM verification WHERE user_id=$1`
	_, err := r.db.Exec(stmt, user_id)
	if err != nil {
		return err
	}
	return nil
}

func (r *VerifyRepositoryImpl) GetCodeByUser(user_id int) (string, error) {
	var (
		code string
		stmt string = `SELECT code FROM verification WHERE user_id=$1`
	)
	if err := r.db.QueryRow(stmt, user_id).Scan(&code); err != nil {
		return "", err
	}
	return code, nil
}
