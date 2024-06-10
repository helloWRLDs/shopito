package errors

import (
	"net/http"
	jsonutil "shopito/pkg/util/json"
)

func SendErr(w http.ResponseWriter, err error) {
	e, ok := err.(*HTTPError)
	if !ok {
		jsonutil.EncodeJson(w, ErrInternal.StatusCode, ErrInternal.SetMessage(err.Error()))
		return
	}
	jsonutil.EncodeJson(w, e.StatusCode, e)
}

func GetHTTPErrorByCode(code int, message string) *HTTPError {
	var err *HTTPError

	switch code {
	case 400:
		err = ErrBadRequest
	case 401:
		err = ErrNotAuthorized
	case 403:
		err = ErrForbidden
	case 404:
		err = ErrNotFound
	case 409:
		err = ErrConflict
	case 422:
		err = ErrUnpocessableEntity
	case 429:
		err = ErrTooManyRequests
	case 500:
		err = ErrInternal
	default:
		err = ErrInternal
	}

	err.SetMessage(message)
	return err
}
