package model

import (
	domainOpenCommissionModel "pixstall-artist/domain/open-commission/model"
	model2 "pixstall-artist/domain/user/model"
)

type Artist struct {
	model2.User       `json:",inline"`
	ArtistID          string                                     `json:"artistId"`
	Fans              Fans                                       `json:"fans,omitempty"`
	ArtistIntro       ArtistIntro                                `json:"artistIntro"`
	ArtistBoard       ArtistBoard                                `json:"artistBoard"`
	PaymentMethods    []string                                   `json:"paymentMethods"`
	CommissionDetails CommissionDetails                          `json:"commissionDetails"`
	OpenCommissions   []domainOpenCommissionModel.OpenCommission `json:"openCommissions"`
	Bookmarks         []string                                   `json:"bookmarks"`
	BookmarkCount     int                                        `json:"bookmarkCount,omitempty"`
}

type ArtistIntro struct {
	YearOfDrawing int      `json:"yearOfDrawing" bson:"yearOfDrawing"`
	ArtTypes      []string `json:"artTypes" bson:"artTypes"`
}

type ArtistBoard struct {
	BannerPath string `json:"bannerPath" bson:"bannerPath"`
	Desc       string `json:"desc" bson:"desc"`
}
