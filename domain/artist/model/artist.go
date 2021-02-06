package model

import (
	domainArtworkModel "pixstall-artist/domain/artwork/model"
	"pixstall-artist/domain/fan/model"
	domainOpenCommissionModel "pixstall-artist/domain/open-commission/model"
	"time"
)

type Artist struct {
	ArtistID          string                                     `json:"artistId"`
	UserID            string                                     `json:"userId"`
	UserName          string                                     `json:"userName"`
	Email             string                                     `json:"email,omitempty"`
	Birthday          string                                     `json:"birthday,omitempty"`
	Gender            string                                     `json:"gender,omitempty"`
	ProfilePath       string                                     `json:"profilePath"`
	State             UserState                                  `json:"state"`
	Fans              map[string]model.Fan                       `json:"fans,omitempty"`
	RegTime           time.Time                                  `json:"regTime"`
	LastUpdatedTime   time.Time                                  `json:"lastUpdatedTime,omitempty"`
	ArtistIntro       ArtistIntro                                `json:"artistIntro"`
	ArtistBoard       ArtistBoard                                `json:"artistBoard"`
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
	State       *UserState
	//Fans            *map[string]domainFanModel.Fan
	RegTime     *time.Time
	ArtistIntro *ArtistIntroUpdater
	ArtistBoard *ArtistBoardUpdater
	//CommissionDetails   *CommissionDetailsUpdater
	//OpenCommissions *[]domainOpenCommissionModel.OpenCommissionUpdater
	//Artworks        *[]domainArtworkModel.ArtworkUpdater
}
