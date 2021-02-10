package mongo

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	mongoModel "pixstall-artist/app/artist/repo/mongo/model"
	"pixstall-artist/domain/artist"
	"pixstall-artist/domain/artist/model"
	model2 "pixstall-artist/domain/artwork/model"
	model3 "pixstall-artist/domain/user/model"
	"testing"
	"time"
)

var db *mongo.Database
var dbClient *mongo.Client
var repo artist.Repo
var ctx context.Context

const (
	TestDBName = "pixstall-artist-test"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	fmt.Println("Before all tests")
	ctx = context.TODO()
	var err error
	dbClient, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	db = dbClient.Database(TestDBName)
	repo = NewMongoArtistRepo(db)
}

func teardown() {
	dropAll()
	fmt.Println("After all tests")
	err := dbClient.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
}

func TestMongoArtistRepo_SaveArtist(t *testing.T) {
	cleanAll()
	newTime := time.Now()
	avgRatings := 10
	dArtist := model.Artist{
		ArtistID:    "new_ArtistID",
		User: model3.User{
			UserID:          "new_UserID",
			UserName:        "new_UserName",
			ProfilePath:     "/temp/pic",
			Email:           "temp@mail.com",
			Birthday:        "20000101",
			Gender:          "M",
			State:           model3.UserStateActive,
			RegTime:         newTime,
			LastUpdatedTime: newTime,
		},
		Fans:        model.Fans{
			Meta:  nil,
			Total: 0,
		},
		ArtistIntro: model.ArtistIntro{
			YearOfDrawing: 10,
			ArtTypes:      []string{"Art", "Comic"},
		},
		CommissionDetails: model.CommissionDetails{
			CommissionRequestCount: 10,
			CommissionAcceptCount:  20,
			CommissionSuccessCount: 30,
			AvgRatings:             &avgRatings,
			LastRequestTime:        nil,
		},
		OpenCommissions: nil,
		Artworks: []model2.Artwork{
			{
				ID:           "new_ArtworkID1",
				ArtistID:     "new_ArtistID",
				ClientID:     "new_ArtistID1",
				Rating:       5,
				RequestTime:  newTime,
				CompleteTime: newTime,
				State:        model2.ArtworkStateActive,
			},
			{
				ID:           "new_ArtworkID2",
				ArtistID:     "new_ArtistID",
				ClientID:     "new_ArtistID2",
				Rating:       0,
				RequestTime:  newTime,
				CompleteTime: newTime,
				State:        model2.ArtworkStateRemoved,
			},
		},
	}
	err := repo.SaveArtist(ctx, &dArtist)
	assert.NoError(t, err)

	mongoArtist := mongoModel.Artist{}
	err = db.Collection(ArtistCollection).FindOne(ctx, bson.M{"artistId": "new_ArtistID"}).Decode(&mongoArtist)

	assert.Equal(t, "new_ArtistID", mongoArtist.ArtistID)
	assert.Equal(t, "new_UserID", mongoArtist.UserID)
	assert.Equal(t, "new_UserName", mongoArtist.UserName)
	assert.Equal(t, "temp@mail.com", mongoArtist.Email)
	assert.Equal(t, "20000101", mongoArtist.Birthday)
	assert.Equal(t, "M", mongoArtist.Gender)
	assert.Equal(t, "/temp/pic", mongoArtist.ProfilePath)
	assert.Equal(t, model3.UserStateActive, mongoArtist.State)
	assert.Nil(t, mongoArtist.Fans)
	assert.Equal(t, newTime.Truncate(time.Millisecond).UnixNano(), mongoArtist.RegTime.UnixNano())

}

//Private
func cleanAll() {
	_, err := db.Collection(ArtistCollection).DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println(err)
	}
}

func dropAll() {
	err := db.Collection(ArtistCollection).Drop(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
}

func insertDummyArtist(ctx context.Context, userId string, state model3.UserState) primitive.ObjectID {
	c := db.Collection(ArtistCollection)

	user := mongoModel.Artist{
		ObjectID:         primitive.ObjectID{},
		User: model3.User{
			UserID:          userId,
			UserName:        "temp_UserName",
			ProfilePath:     "",
			Email:           "temp_Email",
			Birthday:        "20200101",
			Gender:          "M",
			State:           state,
			RegTime:         time.Now(),
			LastUpdatedTime: time.Now(),
		},
		Fans:             model.Fans{
			Meta:  nil,
			Total: 0,
		},
		ArtistIntro: model.ArtistIntro{
			YearOfDrawing: 10,
			ArtTypes:      []string{"Comic"},
		},
		CommissionDetails: model.CommissionDetails{},
		OpenCommissions:   nil,
		Artworks:          nil,
	}
	result, err := c.InsertOne(ctx, &user)
	if err != nil {
		panic(err)
	}
	return result.InsertedID.(primitive.ObjectID)
}
