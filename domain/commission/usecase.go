package commission

import (
	"context"
	"pixstall-artist/domain/commission/model"
)

type UseCase interface {
	ValidateNewCommission(ctx context.Context, creator model.CommissionCreator) error
}