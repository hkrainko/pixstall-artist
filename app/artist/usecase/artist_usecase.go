package usecase

import (
	"context"
	"pixstall-artist/app/domain/artist"
	domainArtistModel "pixstall-artist/app/domain/artist/model"
)

type artistUseCase struct {
	artistRepo artist.Repo
}

func NewArtistUseCase(artistRepo artist.Repo) artist.UseCase {
	return &artistUseCase{
		artistRepo: artistRepo,
	}
}

func (a artistUseCase) RegisterNewArtist(ctx context.Context, dArtist *domainArtistModel.Artist) error {
	err := a.artistRepo.SaveArtist(ctx, dArtist)
	return err
}

func (a artistUseCase) GetArtist(ctx context.Context, artistID string) (*domainArtistModel.Artist, error) {
	dArtist, err := a.artistRepo.GetArtist(ctx, artistID)
	if err != nil {
		return nil, err
	}
	return dArtist, nil
}

func (a artistUseCase) UpdateArtist(ctx context.Context, artistID string, updater domainArtistModel.ArtistUpdater) error {
	panic("implement me")
}

