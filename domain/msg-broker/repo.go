package msg_broker

import (
	"context"
)

type Repo interface {
	SendValidatedCommissionMsg(ctx context.Context, err error) error
}

