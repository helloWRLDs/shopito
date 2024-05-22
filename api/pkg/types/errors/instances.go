package errors

var (
	ErrBadRequest         = New("Bad Request", 400)
	ErrNotAuthorized      = New("Not Authorized", 401)
	ErrForbidden          = New("Forbidden", 403)
	ErrNotFound           = New("Not Found", 404)
	ErrConflict           = New("Conflict: Resource Already Exist", 409)
	ErrUnpocessableEntity = New("Unprocessable Entity", 422)
	ErrTooManyRequests    = New("Too Many Requests", 429)
	ErrInternal           = New("Internal Server Error", 500)
)
