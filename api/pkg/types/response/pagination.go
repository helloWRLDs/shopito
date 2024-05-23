package response

type PaginationResponse struct {
	Data     interface{} `json:"data"`
	PrevPage string      `json:"prev"`
	NextPage string      `json:"next"`
}

func NewPagination(data interface{}, prev, next string) *PaginationResponse {
	return &PaginationResponse{
		Data:     data,
		PrevPage: prev,
		NextPage: next,
	}
}
