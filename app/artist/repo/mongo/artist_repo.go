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
	error2 "pixstall-artist/domain/error"
	domainFanModel "pixstall-artist/domain/fan/model"
	model2 "pixstall-artist/domain/model"
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

func (m mongoArtistRepo) GetArtists(ctx context.Context, filter model.ArtistFilter, sorter model.ArtistSorter) (*[]model.Artist, error) {

	opts := options.Find()
	if sorter.RegTime != nil {
		desc := -1 //desc
		if *sorter.RegTime == model2.SortOrderAscending {
			desc = 1 //asc
		}
		opts.SetSort(bson.D{{"regTime", desc}})
	}
	opts.SetSkip(int64(filter.Offset))
	opts.SetLimit(int64(filter.Count))
	mongoFilter := bson.M{}
	if filter.State != nil {
		mongoFilter["state"] = *filter.State
	}

	cursor, err := m.collection.Find(ctx, mongoFilter, opts)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, model.ArtistErrorNotFound
		default:
			return nil, model.ArtistErrorUnknown
		}
	}
	mongoArtists := make([]mongoModel.Artist, 0)
	if err = cursor.All(ctx, &mongoArtists); err != nil {
		return nil, model.ArtistErrorUnknown
	}
	domainArtists := make([]model.Artist, 0)
	for _, v := range mongoArtists {
		domainArtists = append(domainArtists, *v.ToDomainArtist())
	}
	return &domainArtists, nil
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

	// ArtistIntro
	if updater.YearOfDrawing != nil {
		update["artistIntro.yearOfDrawing"] = updater.YearOfDrawing
	}
	if updater.ArtTypes != nil {
		update["artistIntro.artTypes"] = updater.ArtTypes
	}

	// ArtistBoard
	if updater.BannerPath != nil {
		update["artistBoard.bannerPath"] = updater.BannerPath
	}
	if updater.Desc != nil {
		update["artistBoard.desc"] = updater.Desc
	}

	// CommissionDetails
	if updater.CommissionRequestCount != nil {
		update["commissionDetails.commissionRequestCount"] = updater.CommissionRequestCount
	}
	if updater.CommissionAcceptCount != nil {
		update["commissionDetails.commissionAcceptCount"] = updater.CommissionAcceptCount
	}
	if updater.CommissionSuccessCount != nil {
		update["commissionDetails.commissionSuccessCount"] = updater.CommissionSuccessCount
	}
	if updater.AvgRatings != nil {
		update["commissionDetails.avgRatings"] = updater.AvgRatings
	}
	if updater.LastRequestTime != nil {
		update["commissionDetails.lastRequestTime"] = updater.LastRequestTime
	}

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
	if updater.PaymentMethods != nil {
		update["paymentMethods"] = updater.PaymentMethods
	}
	if updater.LastUpdatedTime != nil {
		update["lastUpdatedTime"] = updater.LastUpdatedTime
	}

	result, err := collection.UpdateOne(ctx, filter, bson.M{"$set": update})

	if err != nil {
		return err
	}
	fmt.Printf("UpdateArtist success: %v", result.UpsertedID)
	return nil
}

// bookmark
func (m mongoArtistRepo) AddBookmark(ctx context.Context, userID string, artistID string) error {
	collection := m.db.Collection(ArtistCollection)

	filter := bson.M{"artistId": artistID}

	change := bson.M{"$set": bson.M{"bookmarks." + userID: time.Now()}}

	result, err := collection.UpdateOne(ctx, filter, change)
	if err != nil {
		return error2.UnknownError
	}
	fmt.Printf("AddBookmark success: %v", result.UpsertedID)
	return nil
}

func (m mongoArtistRepo) RemoveBookmark(ctx context.Context, userID string, artistID string) error {
	collection := m.db.Collection(ArtistCollection)

	filter := bson.M{"artistId": artistID}

	change := bson.M{"$unset": "bookmarks." + userID}

	result, err := collection.UpdateOne(ctx, filter, change)
	if err != nil {
		return error2.UnknownError
	}
	fmt.Printf("RemoveBookmark success: %v", result.UpsertedID)
	return nil
}

// fan
func (m mongoArtistRepo) AddFan(ctx context.Context, artistID string, fan domainFanModel.Fan) error {
	collection := m.db.Collection(ArtistCollection)

	filter := bson.M{"artistId": artistID}

	change := bson.M{"$set": bson.M{"fans": bson.M{fan.UserID: fan}}}

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
