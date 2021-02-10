package mongo

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pixstall-artist/app/open-commission/repo/mongo/dao"
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
	mongoOpenComm := dao.NewFromDomainOpenCommissionCreator(artistID, openCommCreator, openCommIDPrefix + newUUID.String())
	_, err = m.collection.InsertOne(ctx, mongoOpenComm)
	if err != nil {
		return nil, domainOpenCommissionModel.OpenCommissionErrorUnknown
	}
	return mongoOpenComm.ToDomainOpenCommission(), nil
}

func (m mongoOpenCommissionRepo) GetOpenCommission(ctx context.Context, openCommID string) (*domainOpenCommissionModel.OpenCommission, error) {
	filter := bson.M{"openCommId": openCommID}
	daoOpenComm := dao.OpenCommission{}
	err := m.collection.FindOne(ctx, filter, nil).Decode(&daoOpenComm)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, domainOpenCommissionModel.OpenCommissionErrorNotFound
		default:
			return nil, domainOpenCommissionModel.OpenCommissionErrorUnknown
		}
	}
	return daoOpenComm.ToDomainOpenCommission(), nil
}

func (m mongoOpenCommissionRepo) GetOpenCommissions(ctx context.Context, filter domainOpenCommissionModel.OpenCommissionFilter) ([]domainOpenCommissionModel.OpenCommission, error) {
	daoFilter := dao.NewFilterFromDomainOpenCommissionFilter(filter)
	opts := options.FindOptions{}
	if filter.Offset != nil {
		opts.Skip = filter.Offset
	}
	if filter.Count != nil {
		opts.Limit = filter.Count
	}

	cursor, err := m.collection.Find(ctx, bson.M{
		"openCommissions": daoFilter,
	}, &opts)
	if err != nil {
		return nil, domainOpenCommissionModel.OpenCommissionErrorUnknown
	}
	defer cursor.Close(ctx)

	var dOpenComm []domainOpenCommissionModel.OpenCommission
	for cursor.Next(ctx) {
		var r dao.OpenCommission
		if err := cursor.Decode(&r); err != nil {
			return nil, err
		}
		dOpenComm = append(dOpenComm, *r.ToDomainOpenCommission())
	}
	return dOpenComm, nil
}

func (m mongoOpenCommissionRepo) UpdateOpenCommission(ctx context.Context, openCommUpdater domainOpenCommissionModel.OpenCommissionUpdater) error {
	filter := bson.M {
		"artistId": openCommUpdater.ArtistID,
		"openCommissions.openCommId": openCommUpdater.ID,
	}
	updater := dao.NewUpdaterFromOpenCommissionUpdater(openCommUpdater)

	_, err := m.collection.UpdateOne(ctx, filter, updater)
	if err != nil {
		return domainOpenCommissionModel.OpenCommissionErrorUnknown
	}
	return nil
}
