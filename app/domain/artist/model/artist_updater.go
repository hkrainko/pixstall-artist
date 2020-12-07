package model

import (
	domainArtworkModel "pixstall-artist/app/domain/artwork/model"
	domainOpenCommissionModel "pixstall-artist/app/domain/open-commission/model"
	userDomainModel "pixstall-artist/app/domain/user/model"
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
	FansIDs          *map[string]userDomainModel.User
	LikeIDs          *map[string]userDomainModel.User
	RegistrationTime *time.Time
	ArtistIntro      *ArtistIntroUpdater
	ArtistDetails    *ArtistDetailsUpdater
	OpenCommissions  *[]domainOpenCommissionModel.OpenCommissionUpdater
	Artworks         *[]domainArtworkModel.Artwork
}
