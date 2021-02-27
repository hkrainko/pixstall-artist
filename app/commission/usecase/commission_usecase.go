package usecase

import (
	"context"
	"pixstall-artist/domain/commission"
	"pixstall-artist/domain/commission/model"
	msgBroker "pixstall-artist/domain/msg-broker"
	openComm "pixstall-artist/domain/open-commission"
	model2 "pixstall-artist/domain/open-commission/model"
)

type commissionUseCase struct {
	msgBrokerRepo msgBroker.Repo
	openCommRepo openComm.Repo
}

func NewCommissionUseCase(msgBrokerRepo msgBroker.Repo, openCommRepo openComm.Repo) commission.UseCase {
	return &commissionUseCase{
		msgBrokerRepo: msgBrokerRepo,
		openCommRepo: openCommRepo,
	}
}

func (c commissionUseCase) ValidateNewCommission(ctx context.Context, comm model.Commission) error {

	// Checking
	tOpenComm, err := c.openCommRepo.GetOpenCommission(ctx, comm.OpenCommissionID)
	if err == nil && tOpenComm.State != model2.OpenCommissionStateActive {
		err = model.CommissionErrorStateNotAllowed
	}
	if err == nil && getHKPrice(comm.Price).Amount < getHKPrice(tOpenComm.Price).Amount {
		err = model.CommissionErrorPriceInvalid
	}
	if err == nil && comm.DayNeed < tOpenComm.DayNeed.From {
		err = model.CommissionErrorDayNeedInvalid
	}
	if err == nil && comm.BePrivate && !tOpenComm.AllowBePrivate {
		err = model.CommissionErrorNotAllowBePrivate
	}
	if err == nil && comm.Anonymous && !tOpenComm.AllowAnonymous {
		err = model.CommissionErrorNotAllowAnonymous
	}
	err = c.msgBrokerRepo.SendCommOpenCommValidationMsg(ctx, getCommissionOpenCommissionValidation(comm.ID, err))
	if err != nil {
		return err
	}
	return nil
}

func getHKPrice(price model2.Price) model2.Price {
	return model2.Price{
		Amount: price.Amount,
		Currency: model2.CurrencyHKD,
	}
}

func getCommissionOpenCommissionValidation(commID string, err error) model.CommissionOpenCommissionValidation {
	if err != nil {
		reason := err.Error()
		return model.CommissionOpenCommissionValidation{
			ID:            commID,
			IsValid:       false,
			InvalidReason: &reason,
		}
	} else {
		return model.CommissionOpenCommissionValidation{
			ID:            commID,
			IsValid:       true,
		}
	}
}