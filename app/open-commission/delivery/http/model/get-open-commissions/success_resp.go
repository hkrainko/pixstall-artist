package get_open_commissions

import "pixstall-artist/domain/open-commission/model"

type Response struct {
	ArtistID        string                 `json:"artistId"`
	OpenCommissions []model.OpenCommission `json:"openCommissions"`
	Offset          int                    `json:"offset"`
	FetchCount      int                    `json:"fetchCount"`
	Total           int                    `json:"total"`
}

func NewResponse(result model.GetOpenCommissionsResult, artistID string, fetchCount int) *Response {
	return &Response{
		ArtistID:        artistID,
		OpenCommissions: result.OpenCommissions,
		Offset:          result.Offset,
		FetchCount:      fetchCount,
		Total:           result.Total,
	}
}
