package model

import (
	domainArtworkModel "pixstall-artist/domain/artwork/model"
	domainFanModel "pixstall-artist/domain/fan/model"
	domainOpenCommissionModel "pixstall-artist/domain/open-commission/model"
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
	RegTime *time.Time
	ArtistIntro      *ArtistIntroUpdater
	ArtistDetails    *ArtistDetailsUpdater
	OpenCommissions  *[]domainOpenCommissionModel.OpenCommissionUpdater
	Artworks         *[]domainArtworkModel.ArtworkUpdater
}
