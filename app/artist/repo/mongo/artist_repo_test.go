package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	mongoModel "pixstall-artist/app/artist/repo/mongo/model"
	"pixstall-artist/app/domain/artist"
	"pixstall-artist/app/domain/artist/model"
	domainFanModel "pixstall-artist/app/domain/fan/model"
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

func insertDummyArtist(ctx context.Context, userId string, state model.UserState) primitive.ObjectID {
	c := db.Collection(ArtistCollection)

	user := mongoModel.Artist{
		ObjectID: primitive.ObjectID{},
		Artist:   model.Artist{
			ArtistID:         "temp_ArtistID",
			UserID:           "temp_UserID",
			UserName:         "temp_UserName",
			Email:            "temp_Email",
			Birthday:         "20200101",
			Gender:           "M",
			PhotoURL:         "",
			State:            state,
			Fans:             map[string]domainFanModel.Fan{},
			RegistrationTime: time.Now(),
			ArtistIntro:      model.ArtistIntro{
				YearOfDrawing: 10,
				ArtTypes:      []string{"Comic"},
				SelfIntro:     "",
			},
			ArtistDetails:    model.ArtistDetails{},
			OpenCommissions:  nil,
			Artworks:         nil,
		},
	}
	result, err := c.InsertOne(ctx, &user)
	if err != nil {
		panic(err)
	}
	return result.InsertedID.(primitive.ObjectID)
}