package commission

import (
	"context"
	"pixstall-artist/domain/commission/model"
)

type UseCase interface {
	ValidateCommission(ctx context.Context, creator model.CommissionCreator) error
}