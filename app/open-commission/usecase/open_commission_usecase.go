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
