package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	domainArtistModel "pixstall-artist/app/domain/artist/model"
)

type Artist struct {
	ObjectID primitive.ObjectID `bson:"_id,omitempty"`
	domainArtistModel.Artist
	//ArtistID         string                                     `bson:"artistId,omitempty"`
	//UserID           string                                     `bson:"userId,omitempty"`
	//UserName         string                                     `bson:"userName,omitempty"`
	//Email            string                                     `bson:"email,omitempty"`
	//Birthday         string                                     `bson:"birthday,omitempty"`
	//Gender           string                                     `bson:"gender,omitempty"`
	//PhotoURL         string                                     `bson:"photoURL,omitempty"`
	//State            domainArtistModel.UserState                `bson:"state,omitempty"`
	//FansIDs          map[string]userDomainModel.User            `bson:"fansIDs,omitempty"`
	//LikeIDs          map[string]userDomainModel.User            `bson:"likeIDs,omitempty"`
	//RegistrationTime time.Time                                  `bson:"registrationTime,omitempty"`
	//ArtistIntro      domainArtistModel.ArtistIntro              `bson:"artistIntro,omitempty"`
	//ArtistDetails    domainArtistModel.ArtistDetails            `bson:"artistDetails,omitempty"`
	//House            domainHouseModel.House                     `bson:"house,omitempty"`
	//OpenCommissions  []domainOpenCommissionModel.OpenCommission `bson:"openCommissions,omitempty"`
	//Artworks         []domainArtworkModel.Artwork               `bson:"artworks,omitempty"`
}

func NewFromDomainArtist(d *domainArtistModel.Artist) Artist {
	return Artist{
		primitive.ObjectID{}, *d,
	}
}

func (a *Artist) ToDomainArtist() *domainArtistModel.Artist {
	return &(a.Artist)
}
