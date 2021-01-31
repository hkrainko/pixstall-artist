package open_commission

import (
	"context"
	domainOpenCommModel "pixstall-artist/domain/open-commission/model"
)

type UseCase interface {
	GetOpenCommission(ctx context.Context, id string, requesterID *string) (domainOpenCommModel.OpenCommission, error)
	GetOpenCommissions(ctx context.Context, filter domainOpenCommModel.OpenCommissionFilter) ([]domainOpenCommModel.OpenCommission, error)
}
