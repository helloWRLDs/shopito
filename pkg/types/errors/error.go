package errors

import "fmt"

type HTTPError struct {
	StatusCode  int    `json:"status"`
	ErrorString string `json:"err"`
	Message     string `json:"msg,omitempty"`
}

func New(err string, status int) *HTTPError {
	return &HTTPError{
		ErrorString: err,
		StatusCode:  status,
	}
}

func (e *HTTPError) SetMessage(msg string) error {
	e.Message = msg
	return e
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%s[%v]: %s", e.ErrorString, e.StatusCode, e.Message)
}
