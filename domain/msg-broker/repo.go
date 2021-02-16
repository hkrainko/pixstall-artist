package msg_broker

import (
	"context"
	"pixstall-artist/domain/commission/model"
)

type Repo interface {
	SendAddCommissionMsg(ctx context.Context, creator model.CommissionCreator) error
}

