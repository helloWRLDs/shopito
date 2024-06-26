// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package repository

import (
	"database/sql"
)

type Category struct {
	ID   int64
	Name string
}

type Product struct {
	ID         int64
	Name       string
	ImgUrl     sql.NullString
	Price      int32
	Stock      int32
	CategoryID sql.NullInt64
}
