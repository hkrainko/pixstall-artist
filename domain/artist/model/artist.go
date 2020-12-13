package model

import (
	domainArtworkModel "pixstall-artist/domain/artwork/model"
	"pixstall-artist/domain/fan/model"
	domainOpenCommissionModel "pixstall-artist/domain/open-commission/model"
	"time"
)

type Artist struct {
	ArtistID         string
	UserID           string
	UserName         string
	Email            string
	Birthday         string
	Gender           string
	ProfilePath      string
	State            UserState
	Fans             map[string]model.Fan
	RegistrationTime time.Time
	ArtistIntro      ArtistIntro
	ArtistDetails    ArtistDetails
	OpenCommissions  []domainOpenCommissionModel.OpenCommission
	Artworks         []domainArtworkModel.Artwork
}
