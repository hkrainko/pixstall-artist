package open_commission

import (
	"context"
	domainOpenCommissionModel "pixstall-artist/domain/open-commission/model"
)

type Repo interface {
	AddOpenCommission(ctx context.Context, artistID string, openComm domainOpenCommissionModel.OpenCommissionCreator) (*domainOpenCommissionModel.OpenCommission, error)
	GetOpenCommission(ctx context.Context, openCommID string) (*domainOpenCommissionModel.OpenCommission, error)
	GetOpenCommissions(ctx context.Context, filter domainOpenCommissionModel.OpenCommissionFilter) (*domainOpenCommissionModel.GetOpenCommissionsResult, error)
	UpdateOpenCommission(ctx context.Context, openCommUpdater domainOpenCommissionModel.OpenCommissionUpdater) error
}