package msg_broker

import (
	"context"
	"pixstall-artist/domain/commission/model"
)

type Repo interface {
	SendCommOpenCommValidationMsg(ctx context.Context, validation model.CommissionOpenCommissionValidation) error
}

