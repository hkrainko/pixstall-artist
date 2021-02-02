package model

import (
	domainArtworkModel "pixstall-artist/domain/artwork/model"
	"pixstall-artist/domain/fan/model"
	domainOpenCommissionModel "pixstall-artist/domain/open-commission/model"
	"time"
)

type Artist struct {
	ArtistID          string
	UserID            string
	UserName          string
	Email             string
	Birthday          string
	Gender            string
	ProfilePath       string
	State             UserState
	Fans              map[string]model.Fan
	RegTime           time.Time
	LastUpdatedTime   time.Time
	ArtistIntro       ArtistIntro
	ArtistBoard       ArtistBoard
	CommissionDetails CommissionDetails
	OpenCommissions   []domainOpenCommissionModel.OpenCommission
	Artworks          []domainArtworkModel.Artwork
}

type ArtistUpdater struct {
	ArtistID    string
	UserName    *string
	Email       *string
	Birthday    *string
	Gender      *string
	ProfilePath *string
	State       *UserState
	//Fans            *map[string]domainFanModel.Fan
	RegTime     *time.Time
	ArtistIntro *ArtistIntroUpdater
	ArtistBoard *ArtistBoardUpdater
	//CommissionDetails   *CommissionDetailsUpdater
	//OpenCommissions *[]domainOpenCommissionModel.OpenCommissionUpdater
	//Artworks        *[]domainArtworkModel.ArtworkUpdater
}
