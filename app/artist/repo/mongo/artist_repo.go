package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	mongoModel "pixstall-artist/app/artist/repo/mongo/model"
	"pixstall-artist/app/domain/artist"
	"pixstall-artist/app/domain/artist/model"
	domainArtworkModel "pixstall-artist/app/domain/artwork/model"
	domainOpenCommissionModel "pixstall-artist/app/domain/open-commission/model"
)

type mongoArtistRepo struct {
	db         *mongo.Database
	collection *mongo.Collection
}

const (
	DBName           = "pixstall-artist"
	ArtistCollection = "Artists"
)

func NewMongoArtistRepo(db *mongo.Database) artist.Repo {
	return &mongoArtistRepo{
		db:         db,
		collection: db.Collection(ArtistCollection),
	}
}

func (m mongoArtistRepo) SaveArtist(ctx context.Context, dArtist *model.Artist) error {
	result, err := m.collection.InsertOne(ctx, mongoModel.NewFromDomainArtist(dArtist))
	if err != nil {
		fmt.Printf("SaveArtist error %v\n", err)
		return err
	}
	fmt.Printf("SaveArtist %v", result.InsertedID)
	return nil
}

func (m mongoArtistRepo) GetArtist(ctx context.Context, artistID string) (*model.Artist, error) {
	filter := bson.M{"artistId": artistID}
	mongoArtist := mongoModel.Artist{}
	err := m.collection.FindOne(ctx, filter).Decode(&mongoArtist)
	if err != nil {
		return nil, err
	}
	return mongoArtist.ToDomainArtist(), nil
}

func (m mongoArtistRepo) UpdateArtist(ctx context.Context, updater *model.ArtistUpdater) error {
	panic("implement me")
}

func (m mongoArtistRepo) AddOpenCommission(ctx context.Context, openCommission *domainOpenCommissionModel.OpenCommission) error {
	panic("implement me")
}

func (m mongoArtistRepo) AddArtwork(ctx context.Context, artwork *domainArtworkModel.Artwork) error {
	panic("implement me")
}