package usecase

import (
	"context"
	"pixstall-artist/app/domain/artist"
	domainArtistModel "pixstall-artist/app/domain/artist/model"
	domainArtworkModel "pixstall-artist/app/domain/artwork/model"
	domainOpenCommissionModel "pixstall-artist/app/domain/open-commission/model"
	domainRegModel "pixstall-artist/app/domain/reg/model"
	"time"
)

type artistUseCase struct {
	artistRepo artist.Repo
}

func NewArtistUseCase(artistRepo artist.Repo) artist.UseCase {
	return &artistUseCase{
		artistRepo: artistRepo,
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
		PhotoURL:         "",
		State:            domainArtistModel.UserStateActive,
		Fans:             nil,
		RegistrationTime: time.Time{},
		ArtistIntro:      regInfo.RegArtistInfo,
		ArtistDetails:    domainArtistModel.ArtistDetails{
			CommissionRequestCount: 0,
			CommissionAcceptCount:  0,
			CommissionSuccessCount: 0,
			AvgRatings:             0,
			LastRequestTime:        nil,
		},
		OpenCommissions:  nil,
		Artworks:         nil,
	}

	err := a.artistRepo.SaveArtist(ctx, &dArtist)
	return err
}

func (a artistUseCase) GetArtist(ctx context.Context, artistID string) (*domainArtistModel.Artist, error) {
	dArtist, err := a.artistRepo.GetArtist(ctx, artistID)
	if err != nil {
		return nil, err
	}
	return dArtist, nil
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

func (a artistUseCase) UpdateOpenCommission(ctx context.Context, artistID string, updater *domainOpenCommissionModel.OpenCommissionUpdater) error {
	artistUpdater := &domainArtistModel.ArtistUpdater{
		ArtistID:        artistID,
		OpenCommissions: &[]domainOpenCommissionModel.OpenCommissionUpdater{*updater},
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

func (a artistUseCase) AddOpenCommission(ctx context.Context, openCommission *domainOpenCommissionModel.OpenCommission) error {
	return a.artistRepo.AddOpenCommission(ctx, openCommission)
}

func (a artistUseCase) AddArtwork(ctx context.Context, artwork *domainArtworkModel.Artwork) error {
	return a.artistRepo.AddArtwork(ctx, artwork)
}

func (a artistUseCase) DeleteOpenCommission(ctx context.Context, artistID string, openCommissionID string) error {
	newState := domainOpenCommissionModel.OpenCommissionStateRemoved
	openCommissionUpdater := domainOpenCommissionModel.OpenCommissionUpdater{
		ID:       openCommissionID,
		ArtistID: artistID,
		State:    &newState,
	}
	artistUpdater := &domainArtistModel.ArtistUpdater{
		ArtistID:        artistID,
		OpenCommissions: &[]domainOpenCommissionModel.OpenCommissionUpdater{openCommissionUpdater},
	}
	return a.artistRepo.UpdateArtist(ctx, artistUpdater)
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
