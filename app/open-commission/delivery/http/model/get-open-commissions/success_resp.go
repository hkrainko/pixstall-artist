package get_open_commissions

import "pixstall-artist/domain/open-commission/model"

type Response struct {
	ArtistID        string                 `json:"artistId"`
	OpenCommissions []model.OpenCommission `json:"openCommissions"`
	Offset          int                    `json:"offset"`
	Count           int                    `json:"count"`
	Total           int                    `json:"total"`
}

func NewResponse(result model.GetOpenCommissionResult, artistID string) *Response {
	return &Response{
		ArtistID:        artistID,
		OpenCommissions: result.OpenCommissions,
		Offset:          result.Offset,
		Count:           len(result.OpenCommissions),
		Total:           result.Total,
	}
}
