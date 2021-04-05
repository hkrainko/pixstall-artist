package msg_broker

import (
	"context"
	model4 "pixstall-artist/domain/artist/model"
	"pixstall-artist/domain/commission/model"
	model3 "pixstall-artist/domain/open-commission/model"
)

type Repo interface {
	SendCommOpenCommValidationMsg(ctx context.Context, validation model.CommissionOpenCommissionValidation) error
	SendArtistCreatedMsg(ctx context.Context, artist model4.Artist) error
	SendArtistUpdatedMsg(ctx context.Context, updater model4.ArtistUpdater) error
	SendOpenCommCreatedMsg(ctx context.Context, openComm model3.OpenCommission) error
	SendOpenCommUpdatedMsg(ctx context.Context, openComm model3.OpenCommission) error
}
