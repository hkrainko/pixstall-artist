package model

import "image"

type ArtistUpdater struct {
	ArtistID       string
	ArtistIntro    *ArtistIntroUpdater
	ArtistBoard    *ArtistBoardUpdater
	PaymentMethods *[]string
}

type ArtistIntroUpdater struct {
	YearOfDrawing *int      `json:"yearOfDrawing" bson:"yearOfDrawing"`
	ArtTypes      *[]string `json:"artTypes" bson:"artTypes"`
}

type ArtistBoardUpdater struct {
	BannerFile *image.Image
	BannerPath *string
	Desc       *string
}
