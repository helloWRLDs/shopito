package errors

import (
	"database/sql"
	"errors"
)

func WrapSqlErr(err error) *Error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return ErrNotFound
	case errors.Is(err, sql.ErrTxDone):
		return ErrConflict
	case errors.Is(err, sql.ErrConnDone):
		return ErrInternal
	default:
		return ErrInternal
	}
}
