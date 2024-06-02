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