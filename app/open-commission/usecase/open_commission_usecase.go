package usecase

import (
	"context"
	"log"
	msgBroker "pixstall-artist/domain/msg-broker"
	openCommission "pixstall-artist/domain/open-commission"
	domainOpenCommModel "pixstall-artist/domain/open-commission/model"
	"time"
)

type openCommissionUseCase struct {
	openCommRepo openCommission.Repo
	msgBrokerRepo msgBroker.Repo
}

func NewOpenCommissionUseCase(openCommRepo openCommission.Repo, msgBrokerRepo msgBroker.Repo) openCommission.UseCase {
	return &openCommissionUseCase{
		openCommRepo: openCommRepo,
		msgBrokerRepo: msgBrokerRepo,
	}
}

func (o openCommissionUseCase) GetOpenCommission(ctx context.Context, id string, requesterID *string) (domainOpenCommModel.OpenCommission, error) {
	panic("implement me")
}

func (o openCommissionUseCase) GetOpenCommissions(ctx context.Context, filter domainOpenCommModel.OpenCommissionFilter) ([]domainOpenCommModel.OpenCommission, error) {
	panic("implement me")
}

func (o openCommissionUseCase) UpdateOpenCommission(ctx context.Context, requesterID string, updater domainOpenCommModel.OpenCommissionUpdater) error {
	now := time.Now()
	updater.LastUpdatedTime = &now
	err := o.openCommRepo.UpdateOpenCommission(ctx, updater)
	if err != nil {
		return err
	}
	err = o.msgBrokerRepo.SendOpenCommUpdatedMsg(ctx, updater)
	if err != nil {
		log.Println(err)
		// Ignore error
	}
	return nil
}

func (o openCommissionUseCase) DeleteOpenCommission(ctx context.Context, requesterID string, openCommissionID string) error {
	newState := domainOpenCommModel.OpenCommissionStateRemoved
	openCommissionUpdater := domainOpenCommModel.OpenCommissionUpdater{
		ID:       openCommissionID,
		State:    &newState,
	}
	return o.openCommRepo.UpdateOpenCommission(ctx, openCommissionUpdater)
}