package open_commission

import (
	"context"
	domainOpenCommModel "pixstall-artist/domain/open-commission/model"
)

type UseCase interface {
	AddOpenCommission(ctx context.Context, requesterID string, openCommCreator domainOpenCommModel.OpenCommissionCreator) (*string, error)
	GetOpenCommission(ctx context.Context, id string, requesterID *string) (domainOpenCommModel.OpenCommission, error)
	GetOpenCommissions(ctx context.Context, filter domainOpenCommModel.OpenCommissionFilter) (*domainOpenCommModel.GetOpenCommissionsResult, error)
	UpdateOpenCommission(ctx context.Context, requesterID string, updater domainOpenCommModel.OpenCommissionUpdater) error
	DeleteOpenCommission(ctx context.Context, requesterID string, openCommissionID string) error
}
