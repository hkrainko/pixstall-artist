package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mongoModel "pixstall-artist/app/artist/repo/mongo/model"
	"pixstall-artist/domain/artist"
	"pixstall-artist/domain/artist/model"
	domainArtworkModel "pixstall-artist/domain/artwork/model"
)

type mongoArtistRepo struct {
	db         *mongo.Database
	collection *mongo.Collection
}

const (
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
	opt := options.FindOneOptions{
		Projection: bson.D{
			{"email", 0},
			{"birthday", 0},
			{"gender", 0},
			{"openCommissions", 0},
			{"artworks", 0},
		},
	}
	err := m.collection.FindOne(ctx, filter, &opt).Decode(&mongoArtist)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, model.ArtistErrorNotFound
		default:
			return nil, model.ArtistErrorUnknown
		}
	}
	return mongoArtist.ToDomainArtist(), nil
}

func (m mongoArtistRepo) GetArtistDetails(ctx context.Context, artistID string) (*model.Artist, error) {
	filter := bson.M{"artistId": artistID}
	mongoArtist := mongoModel.Artist{}
	opt := options.FindOneOptions{
		Projection: bson.D{
			{"openCommissions", 0},
			{"artworks", 0},
		},
	}
	err := m.collection.FindOne(ctx, filter, &opt).Decode(&mongoArtist)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, model.ArtistErrorNotFound
		default:
			return nil, model.ArtistErrorUnknown
		}
	}
	return mongoArtist.ToDomainArtist(), nil
}

func (m mongoArtistRepo) UpdateArtist(ctx context.Context, updater *model.ArtistUpdater) error {
	panic("implement me")
}

func (m mongoArtistRepo) AddArtwork(ctx context.Context, artwork *domainArtworkModel.Artwork) error {
	panic("implement me")
}