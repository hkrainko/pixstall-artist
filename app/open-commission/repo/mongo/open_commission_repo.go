package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	openCommission "pixstall-artist/domain/open-commission"
	domainOpenCommissionModel "pixstall-artist/domain/open-commission/model"
)

type mongoOpenCommissionRepo struct {
	db         *mongo.Database
	collection *mongo.Collection
}

const (
	ArtistCollection = "Artists"
)

func NewMongoOpenCommissionRepo(db *mongo.Database) openCommission.Repo {
	return &mongoOpenCommissionRepo{
		db: db,
		collection: db.Collection(ArtistCollection),
	}
}

func (m mongoOpenCommissionRepo) AddOpenCommission(ctx context.Context, artistID string, openComm *domainOpenCommissionModel.OpenCommission) error {
	panic("implement me")
}

func (m mongoOpenCommissionRepo) GetOpenCommission(ctx context.Context, openCommID string) (*domainOpenCommissionModel.OpenCommission, error) {
	panic("implement me")
}

func (m mongoOpenCommissionRepo) GetOpenCommissions(ctx context.Context, filter domainOpenCommissionModel.OpenCommissionFilter) ([]domainOpenCommissionModel.OpenCommission, error) {
	panic("implement me")
}

func (m mongoOpenCommissionRepo) UpdateOpenCommission(ctx context.Context, openCommUpdater domainOpenCommissionModel.OpenCommissionUpdater) error {
	panic("implement me")
}