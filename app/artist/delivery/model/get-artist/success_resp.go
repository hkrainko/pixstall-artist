package get_artist

import (
	domain "pixstall-artist/domain/artist/model"
	domainOpenCommissionModel "pixstall-artist/domain/open-commission/model"
	model2 "pixstall-artist/domain/user/model"
)

type Response struct {
	model2.User       `json:",inline"`
	ArtistID          string                                     `json:"artistId"`
	ArtistIntro       domain.ArtistIntro                         `json:"artistIntro"`
	ArtistBoard       domain.ArtistBoard                         `json:"artistBoard"`
	PaymentMethods    []string                                   `json:"paymentMethods"`
	CommissionDetails domain.CommissionDetails                   `json:"commissionDetails"`
	OpenCommissions   []domainOpenCommissionModel.OpenCommission `json:"openCommissions"`
	BookmarkCount     int                                        `json:"bookmarkCount"`
}

func NewResponse(dArtist domain.Artist) *Response {
	return &Response{
		User:              dArtist.User,
		ArtistID:          dArtist.ArtistID,
		ArtistIntro:       dArtist.ArtistIntro,
		ArtistBoard:       dArtist.ArtistBoard,
		PaymentMethods:    dArtist.PaymentMethods,
		CommissionDetails: dArtist.CommissionDetails,
		OpenCommissions:   dArtist.OpenCommissions,
		BookmarkCount:     dArtist.BookmarkCount,
	}
}
