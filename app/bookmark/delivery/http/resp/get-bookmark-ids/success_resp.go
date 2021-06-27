package get_bookmark_ids

type Response struct {
	IDs []string `json:"ids"`
}

func NewResponse(ids []string) *Response {
	return &Response {
		IDs: ids,
	}
}
