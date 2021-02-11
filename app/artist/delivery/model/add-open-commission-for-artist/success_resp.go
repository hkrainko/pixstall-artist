package add_open_commission_for_artist

type Response struct {
	OpenCommID string `json:"openCommissionId"`
}

func NewResponse(openCommID string) *Response {
	return &Response{
		OpenCommID: openCommID,
	}
}
