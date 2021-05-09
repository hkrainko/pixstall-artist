package usecase

import (
	"context"
	"log"
	error2 "pixstall-artist/domain/error"
	domainFile "pixstall-artist/domain/file"
	model2 "pixstall-artist/domain/file/model"
	msgBroker "pixstall-artist/domain/msg-broker"
	openCommission "pixstall-artist/domain/open-commission"
	domainOpenCommModel "pixstall-artist/domain/open-commission/model"
	"time"
)

type openCommissionUseCase struct {
	openCommRepo  openCommission.Repo
	msgBrokerRepo msgBroker.Repo
	imageRepo     domainFile.Repo
}

func NewOpenCommissionUseCase(openCommRepo openCommission.Repo, msgBrokerRepo msgBroker.Repo, imageRepo domainFile.Repo) openCommission.UseCase {
	return &openCommissionUseCase{
		openCommRepo:  openCommRepo,
		msgBrokerRepo: msgBrokerRepo,
		imageRepo: imageRepo,
	}
}

func (o openCommissionUseCase) AddOpenCommission(ctx context.Context, requesterID string, openCommCreator domainOpenCommModel.OpenCommissionCreator) (*string, error) {
	if len(openCommCreator.SampleImages) <= 0 {
		return nil, error2.BadRequestError
	}
	var paths []string
	for _, sampleImage := range openCommCreator.SampleImages {
		path, err := o.imageRepo.SaveFile(ctx, sampleImage.File, model2.FileTypeOpenCommission, requesterID, []string{"*"})
		if err != nil {
			return nil, err
		}
		paths = append(paths, *path)
	}
	if len(paths) <= 0 {
		return nil, error2.UnknownError
	}
	openCommCreator.SampleImagePaths = paths

	addedOpenComm, err := o.openCommRepo.AddOpenCommission(ctx, requesterID, openCommCreator)
	if err != nil {
		return nil, err
	}
	err = o.msgBrokerRepo.SendOpenCommCreatedMsg(ctx, *addedOpenComm)
	if err != nil {
		log.Println(err)
		// Ignore error
	}
	return &addedOpenComm.ID, err
}

func (o openCommissionUseCase) GetOpenCommission(ctx context.Context, id string, requesterID *string) (domainOpenCommModel.OpenCommission, error) {
	panic("implement me")
}

func (o openCommissionUseCase) GetOpenCommissions(ctx context.Context, filter domainOpenCommModel.OpenCommissionFilter) (*domainOpenCommModel.GetOpenCommissionResult, error) {
	return o.openCommRepo.GetOpenCommissions(ctx, filter)
}

func (o openCommissionUseCase) UpdateOpenCommission(ctx context.Context, requesterID string, updater domainOpenCommModel.OpenCommissionUpdater) error {
	dOpenComm, err := o.openCommRepo.GetOpenCommission(ctx, updater.ID)
	if err != nil {
		return nil
	}
	if dOpenComm.ArtistID != requesterID {
		return error2.UnAuthError
	}
	now := time.Now()
	updater.LastUpdatedTime = &now
	err = o.openCommRepo.UpdateOpenCommission(ctx, updater)
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
		ID:    openCommissionID,
		State: &newState,
	}
	return o.openCommRepo.UpdateOpenCommission(ctx, openCommissionUpdater)
}
