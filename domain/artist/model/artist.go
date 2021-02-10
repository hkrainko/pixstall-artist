package model

import (
	domainArtworkModel "pixstall-artist/domain/artwork/model"
	domainOpenCommissionModel "pixstall-artist/domain/open-commission/model"
	model2 "pixstall-artist/domain/user/model"
	"time"
)

type Artist struct {
	model2.User
	ArtistID          string                                     `json:"artistId"`
	Fans              Fans                                       `json:"fans,omitempty"`
	ArtistIntro       ArtistIntro                                `json:"artistIntro"`
	ArtistBoard       ArtistBoard                                `json:"artistBoard"`
	PaymentMethods    []string                                   `json:"paymentMethods"`
	CommissionDetails CommissionDetails                          `json:"commissionDetails"`
	OpenCommissions   []domainOpenCommissionModel.OpenCommission `json:"openCommissions"`
	Artworks          []domainArtworkModel.Artwork               `json:"artworks"`
}

type ArtistUpdater struct {
	ArtistID    string
	UserName    *string
	Email       *string
	Birthday    *string
	Gender      *string
	ProfilePath *string
	State       *model2.UserState
	//Fans            *map[string]domainFanModel.Fan
	RegTime     *time.Time
	ArtistIntro *ArtistIntroUpdater
	ArtistBoard *ArtistBoardUpdater
	//CommissionDetails   *CommissionDetailsUpdater
	//OpenCommissions *[]domainOpenCommissionModel.OpenCommissionUpdater
	//Artworks        *[]domainArtworkModel.ArtworkUpdater
}
