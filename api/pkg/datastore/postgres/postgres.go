package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Open(user, password, name string) (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		user, password, name,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
