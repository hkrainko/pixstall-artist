package usecase

import (
	"context"
	"pixstall-artist/domain/artist"
	domainArtistModel "pixstall-artist/domain/artist/model"
	domainArtworkModel "pixstall-artist/domain/artwork/model"
	openCommission "pixstall-artist/domain/open-commission"
	domainOpenCommissionModel "pixstall-artist/domain/open-commission/model"
	domainRegModel "pixstall-artist/domain/reg/model"
	"time"
)

type artistUseCase struct {
	artistRepo   artist.Repo
	openCommRepo openCommission.Repo
}

func NewArtistUseCase(artistRepo artist.Repo, openCommRepo openCommission.Repo) artist.UseCase {
	return &artistUseCase{
		artistRepo:   artistRepo,
		openCommRepo: openCommRepo,
	}
}

func (a artistUseCase) RegisterNewArtist(ctx context.Context, regInfo *domainRegModel.RegInfo) error {

	dArtist := domainArtistModel.Artist{
		ArtistID:         regInfo.UserID,
		UserID:           regInfo.UserID,
		UserName:         regInfo.DisplayName,
		Email:            regInfo.Email,
		Birthday:         regInfo.Birthday,
		Gender:           regInfo.Gender,
		ProfilePath:      regInfo.ProfilePath,
		State:            domainArtistModel.UserStateActive,
		Fans:             nil,
		RegTime: time.Time{},
		ArtistIntro:      regInfo.RegArtistIntro,
		ArtistDetails: domainArtistModel.ArtistDetails{
			CommissionRequestCount: 0,
			CommissionAcceptCount:  0,
			CommissionSuccessCount: 0,
			AvgRatings:             nil,
			LastRequestTime:        nil,
		},
		ArtistBoard: domainArtistModel.ArtistBoard{},
		OpenCommissions: nil,
		Artworks:        nil,
	}

	err := a.artistRepo.SaveArtist(ctx, &dArtist)
	return err
}

func (a artistUseCase) GetArtist(ctx context.Context, artistID string, requesterID *string) (*domainArtistModel.Artist, error) {
	if artistID == *requesterID {
		dArtist, err := a.artistRepo.GetArtistDetails(ctx, artistID)
		if err != nil {
			return nil, err
		}
		return dArtist, nil
	} else {
		dArtist, err := a.artistRepo.GetArtist(ctx, artistID)
		if err != nil {
			return nil, err
		}
		return dArtist, nil
	}
}

func (a artistUseCase) GetOpenCommissionsForArtist(ctx context.Context, artistID string, requesterID *string, count int, offset int) ([]domainOpenCommissionModel.OpenCommission, error) {
	panic("implement me")
}

func (a artistUseCase) AddOpenCommission(ctx context.Context, requesterID string, openCommission *domainOpenCommissionModel.OpenCommission) error {
	return a.openCommRepo.AddOpenCommission(ctx, requesterID, openCommission)
}

func (a artistUseCase) UpdateBasicInfo(ctx context.Context, artistID string, updater *domainArtistModel.ArtistUpdater) error {
	return a.artistRepo.UpdateArtist(ctx, updater)
}

func (a artistUseCase) UpdateIntro(ctx context.Context, artistID string, updater *domainArtistModel.ArtistIntroUpdater) error {
	artistUpdater := &domainArtistModel.ArtistUpdater{
		ArtistID:    artistID,
		ArtistIntro: updater,
	}
	return a.artistRepo.UpdateArtist(ctx, artistUpdater)
}

func (a artistUseCase) UpdateDetails(ctx context.Context, artistID string, updater *domainArtistModel.ArtistDetailsUpdater) error {
	artistUpdater := &domainArtistModel.ArtistUpdater{
		ArtistID:      artistID,
		ArtistDetails: updater,
	}
	return a.artistRepo.UpdateArtist(ctx, artistUpdater)
}

func (a artistUseCase) UpdateArtwork(ctx context.Context, artistID string, updater *domainArtworkModel.ArtworkUpdater) error {
	artistUpdater := &domainArtistModel.ArtistUpdater{
		ArtistID: artistID,
		Artworks: &[]domainArtworkModel.ArtworkUpdater{*updater},
	}
	return a.artistRepo.UpdateArtist(ctx, artistUpdater)
}

func (a artistUseCase) AddArtwork(ctx context.Context, artwork *domainArtworkModel.Artwork) error {
	return a.artistRepo.AddArtwork(ctx, artwork)
}

func (a artistUseCase) DeleteArtwork(ctx context.Context, artistID string, artworkID string) error {
	state := domainArtworkModel.ArtworkStateRemoved
	artworkUpdater := domainArtworkModel.ArtworkUpdater{
		ID:           artworkID,
		ArtistID:     artistID,
		Rating:       nil,
		RequestTime:  nil,
		CompleteTime: nil,
		State:        &state,
	}
	artistUpdater := &domainArtistModel.ArtistUpdater{
		ArtistID: artistID,
		Artworks: &[]domainArtworkModel.ArtworkUpdater{artworkUpdater},
	}
	return a.artistRepo.UpdateArtist(ctx, artistUpdater)
}
