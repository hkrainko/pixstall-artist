package model

import (
	"pixstall-artist/domain/file/model"
)

type ArtistUpdater struct {
	ArtistID          string
	ArtistIntro       *ArtistIntroUpdater
	ArtistBoard       *ArtistBoardUpdater
	PaymentMethods    *[]string
	CommissionDetails *CommissionDetailsUpdater
}

type ArtistIntroUpdater struct {
	YearOfDrawing *int      `json:"yearOfDrawing" bson:"yearOfDrawing"`
	ArtTypes      *[]string `json:"artTypes" bson:"artTypes"`
}

type ArtistBoardUpdater struct {
	BannerFile *model.ImageFile
	BannerPath *string
	Desc       *string
}
