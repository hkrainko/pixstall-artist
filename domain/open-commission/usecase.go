package open_commission

import (
	"context"
	domainOpenCommModel "pixstall-artist/domain/open-commission/model"
)

type UseCase interface {
	AddOpenCommission(ctx context.Context, artistID string, openComm domainOpenCommModel.OpenCommission) (domainOpenCommModel.OpenCommission, error)
	GetOpenCommission(ctx context.Context, id string, requesterID *string) (domainOpenCommModel.OpenCommission, error)
	GetOpenCommissions(ctx context.Context, filter domainOpenCommModel.OpenCommissionFilter) ([]domainOpenCommModel.OpenCommission, error)
	UpdateOpenCommission(ctx context.Context, requesterID string, updater domainOpenCommModel.OpenCommissionUpdater) error
	DeleteOpenCommission(ctx context.Context, requesterID string, openCommissionID string) error
}
