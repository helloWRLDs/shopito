package errors

type Error struct {
	StatusCode int    `json:"status"`
	ErrorMsg   string `json:"err"`
	Msg        string `json:"msg"`
}

func New(err string, status int) *Error {
	return &Error{
		ErrorMsg:   err,
		StatusCode: status,
	}
}

func (e *Error) SetMessage(msg string) *Error {
	e.Msg = msg
	return e
}

func (e *Error) Status() int {
	return e.StatusCode
}

func (e *Error) Error() string {
	return e.ErrorMsg
}
