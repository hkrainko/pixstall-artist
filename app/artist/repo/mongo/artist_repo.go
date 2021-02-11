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
	domainFanModel "pixstall-artist/domain/fan/model"
	"time"
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
	if updater == nil {
		return model.ArtistErrorUnknown
	}
	collection := m.db.Collection(ArtistCollection)

	filter := bson.M{"artistId": updater.ArtistID}
	update := bson.M{}

	if updater.UserName != nil {
		update["userName"] = updater.UserName
	}
	if updater.Email != nil {
		update["email"] = updater.Email
	}
	if updater.Birthday != nil {
		update["birthday"] = updater.Birthday
	}
	if updater.Gender != nil {
		update["gender"] = updater.Gender
	}
	if updater.ProfilePath != nil {
		update["profilePath"] = updater.ProfilePath
	}
	if updater.State != nil {
		update["state"] = updater.State
	}
	if updater.RegTime != nil {
		update["regTime"] = updater.RegTime
	}
	if updater.ArtistIntro != nil {
		if updater.ArtistIntro.YearOfDrawing != nil {
			update["artistIntro.yearOfDrawing"] = updater.ArtistIntro.YearOfDrawing
		}
		if updater.ArtistIntro.ArtTypes != nil {
			update["artistIntro.artTypes"] = updater.ArtistIntro.ArtTypes
		}
	}
	if updater.ArtistBoard != nil {
		if updater.ArtistBoard.BannerPath != nil {
			update["artistBoard.bannerPath"] = updater.ArtistBoard.BannerPath
		}
		if updater.ArtistBoard.Desc != nil {
			update["artistBoard.desc"] = updater.ArtistBoard.Desc
		}
	}
	update["lastUpdatedTime"] = time.Now()

	result, err := collection.UpdateOne(ctx, filter, bson.M{"$set": update})

	if err != nil {
		return err
	}
	fmt.Printf("UpdateArtist success: %v", result.UpsertedID)
	return nil
}

func (m mongoArtistRepo) AddArtwork(ctx context.Context, artwork *domainArtworkModel.Artwork) error {
	panic("implement me")
}

func (m mongoArtistRepo) AddFan(ctx context.Context, artistID string, fan domainFanModel.Fan) error {
	collection := m.db.Collection(ArtistCollection)

	filter := bson.M{"artistId": artistID}

	change := bson.M{"$push": bson.M{"fans": bson.M{fan.UserID: fan}}}

	result, err := collection.UpdateOne(ctx, filter, change)
	if err != nil {
		return err
	}
	fmt.Printf("AddFan success: %v", result.UpsertedID)
	return nil
}

func (m mongoArtistRepo) RemoveFan(ctx context.Context, artistID string, fanId string) error {
	panic("implement me")
}
