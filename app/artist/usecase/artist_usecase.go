package usecase

import (
	"context"
	"log"
	"pixstall-artist/domain/artist"
	domainArtistModel "pixstall-artist/domain/artist/model"
	domainFile "pixstall-artist/domain/file"
	model2 "pixstall-artist/domain/file/model"
	msgBroker "pixstall-artist/domain/msg-broker"
	openCommission "pixstall-artist/domain/open-commission"
	domainRegModel "pixstall-artist/domain/reg/model"
	"pixstall-artist/domain/user/model"
	"time"
)

type artistUseCase struct {
	artistRepo    artist.Repo
	openCommRepo  openCommission.Repo
	imageRepo     domainFile.Repo
	msgBrokerRepo msgBroker.Repo
}

func NewArtistUseCase(artistRepo artist.Repo, openCommRepo openCommission.Repo, imageRepo domainFile.Repo, msgBrokerRepo msgBroker.Repo) artist.UseCase {
	return &artistUseCase{
		artistRepo:    artistRepo,
		openCommRepo:  openCommRepo,
		imageRepo:     imageRepo,
		msgBrokerRepo: msgBrokerRepo,
	}
}

func (a artistUseCase) RegisterNewArtist(ctx context.Context, regInfo *domainRegModel.RegInfo) error {

	dArtist := domainArtistModel.Artist{
		ArtistID: regInfo.UserID,
		User: model.User{
			UserID:          regInfo.UserID,
			UserName:        regInfo.DisplayName,
			ProfilePath:     regInfo.ProfilePath,
			Email:           regInfo.Email,
			Birthday:        regInfo.Birthday,
			Gender:          regInfo.Gender,
			State:           model.UserStateActive,
			RegTime:         regInfo.RegTime,
			LastUpdatedTime: time.Now(),
		},
		Fans: domainArtistModel.Fans{
			Meta:  nil,
			Total: 0,
		},
		ArtistIntro: regInfo.RegArtistIntro,
		CommissionDetails: domainArtistModel.CommissionDetails{
			CommissionRequestCount: 0,
			CommissionAcceptCount:  0,
			CommissionSuccessCount: 0,
			AvgRatings:             nil,
			LastRequestTime:        nil,
		},
		ArtistBoard:     domainArtistModel.ArtistBoard{},
		OpenCommissions: nil,
	}

	err := a.artistRepo.SaveArtist(ctx, &dArtist)
	if err != nil {
		return err
	}
	err = a.msgBrokerRepo.SendArtistCreatedEventMsg(ctx, dArtist)
	if err != nil {
		log.Println(err)
		// Ignore error
	}
	return nil
}

func (a artistUseCase) GetArtist(ctx context.Context, artistID string) (*domainArtistModel.Artist, error) {
	dArtist, err := a.artistRepo.GetArtist(ctx, artistID)
	if err != nil {
		return nil, err
	}
	return dArtist, nil
}

func (a artistUseCase) GetArtists(ctx context.Context, filter domainArtistModel.ArtistFilter, sorter domainArtistModel.ArtistSorter) (*[]domainArtistModel.Artist, error) {
	dArtist, err := a.artistRepo.GetArtists(ctx, filter, sorter)
	if err != nil {
		return nil, err
	}
	return dArtist, nil
}

func (a artistUseCase) GetArtistDetails(ctx context.Context, artistID string, requesterID *string) (*domainArtistModel.Artist, error) {
	if requesterID == nil || *requesterID != artistID {
		return nil, domainArtistModel.ArtistErrorUnAuth
	}
	dArtist, err := a.artistRepo.GetArtistDetails(ctx, artistID)
	if err != nil {
		return nil, err
	}
	return dArtist, nil
}

func (a artistUseCase) UpdateArtist(ctx context.Context, updater domainArtistModel.ArtistUpdater) (*string, error) {
	if updater.BannerFile != nil {
		path, err := a.imageRepo.SaveFile(ctx, updater.BannerFile.File, model2.FileTypeRoof, updater.ArtistID, []string{"*"})
		if err != nil {
			return nil, err
		}
		updater.BannerPath = path
	}
	err := a.artistRepo.UpdateArtist(ctx, &updater)
	if err != nil {
		return nil, err
	}
	err = a.msgBrokerRepo.SendArtistUpdatedEventMsg(ctx, updater)
	if err != nil {
		log.Println(err)
		// Ignore error
	}
	return &updater.ArtistID, nil
}

// Open Commission

//func (a artistUseCase) AddOpenCommission(ctx context.Context, requesterID string, openCommCreator domainOpenCommissionModel.OpenCommissionCreator) (*string, error) {
//	if len(openCommCreator.SampleImages) <= 0 {
//		return nil, error2.BadRequestError
//	}
//	var paths []string
//	for _, sampleImage := range openCommCreator.SampleImages {
//		path, err := a.imageRepo.SaveFile(ctx, sampleImage.File, model2.FileTypeOpenCommission, requesterID, []string{"*"})
//		if err != nil {
//			return nil, err
//		}
//		paths = append(paths, *path)
//	}
//	if len(paths) <= 0 {
//		return nil, error2.UnknownError
//	}
//	openCommCreator.SampleImagePaths = paths
//
//	addedOpenComm, err := a.openCommRepo.AddOpenCommission(ctx, requesterID, openCommCreator)
//	if err != nil {
//		return nil, err
//	}
//	err = a.msgBrokerRepo.SendOpenCommCreatedMsg(ctx, *addedOpenComm)
//	if err != nil {
//		log.Println(err)
//		// Ignore error
//	}
//	return &addedOpenComm.ID, err
//}
