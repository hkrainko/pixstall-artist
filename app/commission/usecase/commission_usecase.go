package usecase

import (
	"context"
	"pixstall-artist/domain/commission"
	"pixstall-artist/domain/commission/model"
	domainImage "pixstall-artist/domain/image"
	msgBroker "pixstall-artist/domain/msg-broker"
	openComm "pixstall-artist/domain/open-commission"
	model2 "pixstall-artist/domain/open-commission/model"
)

type commissionUseCase struct {
	msgBrokerRepo msgBroker.Repo
	openCommRepo openComm.Repo
	imageRepo domainImage.Repo
}

func NewCommissionUseCase(msgBrokerRepo msgBroker.Repo, openCommRepo openComm.Repo, imageRepo domainImage.Repo) commission.UseCase {
	return &commissionUseCase{
		msgBrokerRepo: msgBrokerRepo,
		openCommRepo: openCommRepo,
		imageRepo: imageRepo,
	}
}

func (c commissionUseCase) ValidateCommission(ctx context.Context, creator model.CommissionCreator) (error) {

	// Checking
	tOpenComm, err := c.openCommRepo.GetOpenCommission(ctx, creator.OpenCommissionID)
	if err != nil {
		return err
	}
	if tOpenComm.State != model2.OpenCommissionStateActive {
		return model.CommissionErrorStateNotAllowed
	}
	if getHKPrice(creator.Price).Amount < getHKPrice(tOpenComm.Price).Amount {
		return model.CommissionErrorPriceInvalid
	}
	if creator.DayNeed < tOpenComm.DayNeed.From {
		return model.CommissionErrorDayNeedInvalid
	}
	if creator.BePrivate && !tOpenComm.AllowBePrivate {
		return model.CommissionErrorNotAllowBePrivate
	}
	if creator.Anonymous && !tOpenComm.AllowAnonymous {
		return model.CommissionErrorNotAllowAnonymous
	}

	err = c.msgBrokerRepo.SendValidatedCommissionMsg(ctx, nil)
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