package mongo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"pixstall-artist/app/open-commission/repo/mongo/dao"
	error2 "pixstall-artist/domain/error"
	openCommission "pixstall-artist/domain/open-commission"
	domainOpenCommissionModel "pixstall-artist/domain/open-commission/model"
)

type mongoOpenCommissionRepo struct {
	db         *mongo.Database
	collection *mongo.Collection
}

const (
	ArtistCollection = "Artists"
	openCommIDPrefix = "OC-"
)

func NewMongoOpenCommissionRepo(db *mongo.Database) openCommission.Repo {
	return &mongoOpenCommissionRepo{
		db:         db,
		collection: db.Collection(ArtistCollection),
	}
}

func (m mongoOpenCommissionRepo) AddOpenCommission(ctx context.Context, artistID string, openCommCreator domainOpenCommissionModel.OpenCommissionCreator) (*domainOpenCommissionModel.OpenCommission, error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, domainOpenCommissionModel.OpenCommissionErrorUnknown
	}
	newID := openCommIDPrefix + newUUID.String()
	mongoOpenComm := dao.NewFromDomainOpenCommissionCreator(artistID, openCommCreator, newID)

	filter := bson.M{"artistId": artistID}
	change := bson.M{"$push": bson.M{"openCommissions": mongoOpenComm}}

	_, err = m.collection.UpdateOne(ctx, filter, change)
	if err != nil {
		return nil, domainOpenCommissionModel.OpenCommissionErrorUnknown
	}
	fmt.Printf("AddOpenCommission success, id:%v", newID)
	return mongoOpenComm.ToDomainOpenCommission(), nil
}

func (m mongoOpenCommissionRepo) GetOpenCommission(ctx context.Context, openCommID string) (*domainOpenCommissionModel.OpenCommission, error) {
	pipeline := []bson.M{
		{"$match": bson.M{"openCommissions.id": openCommID}},
		{"$project": bson.M{"openCommissions": 1}},
		{"$unwind": "$openCommissions"},
		{"$match": bson.M{"openCommissions.id": openCommID}},
		{"$replaceRoot": bson.M{"newRoot": "$openCommissions"}},
	}
	cursor, err := m.collection.Aggregate(ctx, pipeline)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, domainOpenCommissionModel.OpenCommissionErrorNotFound
		default:
			return nil, domainOpenCommissionModel.OpenCommissionErrorUnknown
		}
	}
	defer cursor.Close(ctx)
	var daoOpenComm dao.OpenCommission
	for cursor.Next(ctx) {
		if err := cursor.Decode(&daoOpenComm); err != nil {
			return nil, err
		} else {
			return daoOpenComm.ToDomainOpenCommission(), nil
		}
	}
	return nil, domainOpenCommissionModel.OpenCommissionErrorNotFound
}

func (m mongoOpenCommissionRepo) GetOpenCommissions(ctx context.Context, filter domainOpenCommissionModel.OpenCommissionFilter) (*domainOpenCommissionModel.GetOpenCommissionResult, error) {

	pipeline := []bson.M{
		{"$match": bson.M{"openCommissions.id": filter.ArtistID}},
		{"$project": bson.M{"openCommissions": 1, "$total": bson.M{"$size": "$openCommissions"}}},
		{"$slice": bson.A{"$openCommissions", filter.Offset, filter.Count}},
	}

	cursor, err := m.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, error2.UnknownError
	}
	defer cursor.Close(ctx)

	var dGetOpenCommResult *domainOpenCommissionModel.GetOpenCommissionResult
	for cursor.Next(ctx) {
		var r dao.GetOpenCommissionResult
		if err := cursor.Decode(&r); err != nil {
			return nil, err
		}
		dGetOpenCommResult = r.ToDomainGetOpenCommissionResult(filter.Offset)
	}
	return dGetOpenCommResult, nil
}

func (m mongoOpenCommissionRepo) UpdateOpenCommission(ctx context.Context, openCommUpdater domainOpenCommissionModel.OpenCommissionUpdater) error {
	filter := bson.M{
		"openCommissions.openCommId": openCommUpdater.ID,
	}
	updater := dao.NewUpdaterFromOpenCommissionUpdater(openCommUpdater)

	_, err := m.collection.UpdateOne(ctx, filter, updater)
	if err != nil {
		return domainOpenCommissionModel.OpenCommissionErrorUnknown
	}
	return nil
}
