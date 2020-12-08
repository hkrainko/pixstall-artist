package model

import (
	domainArtworkModel "pixstall-artist/app/domain/artwork/model"
	"pixstall-artist/app/domain/fan/model"
	domainOpenCommissionModel "pixstall-artist/app/domain/open-commission/model"
	"time"
)

type Artist struct {
	ArtistID         string
	UserID           string
	UserName         string
	Email            string
	Birthday         string
	Gender           string
	PhotoURL         string
	State            UserState
	Fans             map[string]model.Fan
	RegistrationTime time.Time
	ArtistIntro      ArtistIntro
	ArtistDetails    ArtistDetails
	OpenCommissions  []domainOpenCommissionModel.OpenCommission
	Artworks         []domainArtworkModel.Artwork
}
