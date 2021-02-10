package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	domainArtistModel "pixstall-artist/domain/artist/model"
	domainArtworkModel "pixstall-artist/domain/artwork/model"
	domainOpenCommissionModel "pixstall-artist/domain/open-commission/model"
	model2 "pixstall-artist/domain/user/model"
)

type Artist struct {
	model2.User       `bson:",inline"`
	ObjectID          primitive.ObjectID                         `bson:"_id,omitempty"`
	ArtistID          string                                     `bson:"artistId,omitempty"`
	Fans              domainArtistModel.Fans                     `bson:"fans,omitempty"`
	ArtistIntro       domainArtistModel.ArtistIntro              `bson:"artistIntro,omitempty"`
	CommissionDetails domainArtistModel.CommissionDetails        `bson:"commissionDetails,omitempty"`
	ArtistBoard       domainArtistModel.ArtistBoard              `bson:"artistBoard,omitempty"`
	OpenCommissions   []domainOpenCommissionModel.OpenCommission `bson:"openCommissions,omitempty"`
	Artworks          []domainArtworkModel.Artwork               `bson:"artworks,omitempty"`
}

func NewFromDomainArtist(d *domainArtistModel.Artist) Artist {
	return Artist{
		ArtistID:          d.ArtistID,
		User: model2.User{
			UserID:          d.UserID,
			UserName:        d.UserName,
			ProfilePath:     d.ProfilePath,
			Email:           d.Email,
			Birthday:        d.Birthday,
			Gender:          d.Gender,
			State:           d.State,
			RegTime:         d.RegTime,
			LastUpdatedTime: d.LastUpdatedTime,
		},
		Fans:              d.Fans,
		ArtistIntro:       d.ArtistIntro,
		CommissionDetails: d.CommissionDetails,
		ArtistBoard:       d.ArtistBoard,
		OpenCommissions:   d.OpenCommissions,
		Artworks:          d.Artworks,
	}
}

func (a *Artist) ToDomainArtist() *domainArtistModel.Artist {
	return &domainArtistModel.Artist{
		ArtistID: a.ArtistID,
		User: model2.User{
			UserID:          a.UserID,
			UserName:        a.UserName,
			ProfilePath:     a.ProfilePath,
			Email:           a.Email,
			Birthday:        a.Birthday,
			Gender:          a.Gender,
			State:           a.State,
			RegTime:         a.RegTime,
			LastUpdatedTime: a.LastUpdatedTime,
		},
		Fans:              a.Fans,
		ArtistIntro:       a.ArtistIntro,
		CommissionDetails: a.CommissionDetails,
		ArtistBoard:       a.ArtistBoard,
		OpenCommissions:   a.OpenCommissions,
		Artworks:          a.Artworks,
	}
}
