package response

type JsonMessage struct {
	Code int    `json:"status"`
	Msg  string `json:"msg"`
}

func NewJsonMessage(code int, msg string) *JsonMessage {
	return &JsonMessage{
		Code: code,
		Msg:  msg,
	}
}
