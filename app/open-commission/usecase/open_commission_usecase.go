package usecase

import (
	"context"
	openCommission "pixstall-artist/domain/open-commission"
	domainOpenCommModel "pixstall-artist/domain/open-commission/model"
)

type openCommissionUseCase struct {
	openCommRepo openCommission.Repo
}

func NewOpenCommissionUseCase(openCommRepo openCommission.Repo) openCommission.UseCase {
	return &openCommissionUseCase{
		openCommRepo: openCommRepo,
	}
}

func (o openCommissionUseCase) GetOpenCommission(ctx context.Context, id string, requesterID *string) (domainOpenCommModel.OpenCommission, error) {
	panic("implement me")
}

func (o openCommissionUseCase) GetOpenCommissions(ctx context.Context, filter domainOpenCommModel.OpenCommissionFilter) ([]domainOpenCommModel.OpenCommission, error) {
	panic("implement me")
}

func (o openCommissionUseCase) UpdateOpenCommission(ctx context.Context, requesterID string, updater *domainOpenCommModel.OpenCommissionUpdater) error {
	//artistUpdater := &domainArtistModel.ArtistUpdater{
	//	ArtistID:        artistID,
	//	OpenCommissions: &[]domainOpenCommissionModel.OpenCommissionUpdater{*updater},
	//}
	return o.openCommRepo.UpdateOpenCommission(ctx, *updater)
}

func (o openCommissionUseCase) DeleteOpenCommission(ctx context.Context, requesterID string, openCommissionID string) error {
	//newState := domainOpenCommissionModel.OpenCommissionStateRemoved
	//openCommissionUpdater := domainOpenCommissionModel.OpenCommissionUpdater{
	//	ID:       openCommissionID,
	//	ArtistID: artistID,
	//	State:    &newState,
	//}
	//artistUpdater := &domainArtistModel.ArtistUpdater{
	//	ArtistID:        artistID,
	//	OpenCommissions: &[]domainOpenCommissionModel.OpenCommissionUpdater{openCommissionUpdater},
	//}
	//return o.openCommRepo.UpdateOpenCommission(ctx, artistUpdater)
	return nil
}