package model

type ArtistBoard struct {
	BannerPath string `bson:"bannerPath"`
	Desc       string `bson:"desc"`
}

type ArtistBoardUpdater struct {
	BannerPath *string
	Desc       *string
}
