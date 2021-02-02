package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	domainArtistModel "pixstall-artist/domain/artist/model"
	domainArtworkModel "pixstall-artist/domain/artwork/model"
	"pixstall-artist/domain/fan/model"
	domainOpenCommissionModel "pixstall-artist/domain/open-commission/model"
	"time"
)

type Artist struct {
	ObjectID        primitive.ObjectID                         `bson:"_id,omitempty"`
	ArtistID        string                                     `bson:"artistId,omitempty"`
	UserID          string                                     `bson:"userId,omitempty"`
	UserName        string                                     `bson:"userName,omitempty"`
	Email           string                                     `bson:"email,omitempty"`
	Birthday        string                                     `bson:"birthday,omitempty"`
	Gender          string                                     `bson:"gender,omitempty"`
	ProfilePath     string                                     `bson:"profilePath,omitempty"`
	State           domainArtistModel.UserState                `bson:"state,omitempty"`
	Fans            map[string]model.Fan                       `bson:"fans,omitempty"`
	RegTime         time.Time                                  `bson:"regTime,omitempty"`
	LastUpdatedTime time.Time                                  `bson:"lastUpdatedTime,omitempty"`
	ArtistIntro     domainArtistModel.ArtistIntro              `bson:"artistIntro,omitempty"`
	ArtistDetails   domainArtistModel.ArtistDetails            `bson:"artistDetails,omitempty"`
	ArtistBoard     domainArtistModel.ArtistBoard              `bson:"artistBoard,omitempty"`
	OpenCommissions []domainOpenCommissionModel.OpenCommission `bson:"openCommissions,omitempty"`
	Artworks        []domainArtworkModel.Artwork               `bson:"artworks,omitempty"`
}

func NewFromDomainArtist(d *domainArtistModel.Artist) Artist {
	return Artist{
		ArtistID:        d.ArtistID,
		UserID:          d.UserID,
		UserName:        d.UserName,
		Email:           d.Email,
		Birthday:        d.Birthday,
		Gender:          d.Gender,
		ProfilePath:     d.ProfilePath,
		State:           d.State,
		Fans:            d.Fans,
		RegTime:         d.RegTime,
		LastUpdatedTime: d.LastUpdatedTime,
		ArtistIntro:     d.ArtistIntro,
		ArtistDetails:   d.ArtistDetails,
		ArtistBoard:     d.ArtistBoard,
		OpenCommissions: d.OpenCommissions,
		Artworks:        d.Artworks,
	}
}

func (a *Artist) ToDomainArtist() *domainArtistModel.Artist {
	return &domainArtistModel.Artist{
		ArtistID:        a.ArtistID,
		UserID:          a.UserID,
		UserName:        a.UserName,
		Email:           a.Email,
		Birthday:        a.Birthday,
		Gender:          a.Gender,
		ProfilePath:     a.ProfilePath,
		State:           a.State,
		Fans:            a.Fans,
		RegTime:         a.RegTime,
		LastUpdatedTime: a.LastUpdatedTime,
		ArtistIntro:     a.ArtistIntro,
		ArtistDetails:   a.ArtistDetails,
		ArtistBoard:     a.ArtistBoard,
		OpenCommissions: a.OpenCommissions,
		Artworks:        a.Artworks,
	}
}
