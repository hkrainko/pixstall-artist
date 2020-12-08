package model

import (
	domainArtworkModel "pixstall-artist/app/domain/artwork/model"
	domainFanModel "pixstall-artist/app/domain/fan/model"
	domainOpenCommissionModel "pixstall-artist/app/domain/open-commission/model"
	"time"
)

type ArtistUpdater struct {
	ArtistID         string
	UserName         *string
	Email            *string
	Birthday         *string
	Gender           *string
	PhotoURL         *string
	State            *UserState
	Fans             *map[string]domainFanModel.Fan
	RegistrationTime *time.Time
	ArtistIntro      *ArtistIntroUpdater
	ArtistDetails    *ArtistDetailsUpdater
	OpenCommissions  *[]domainOpenCommissionModel.OpenCommissionUpdater
	Artworks         *[]domainArtworkModel.ArtworkUpdater
}
