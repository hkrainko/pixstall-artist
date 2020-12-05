package artist

import (
	"context"
	domainArtistModel "pixstall-artist/app/domain/artist/model"
)

type UseCase interface {
	RegisterNewArtist(ctx context.Context, dArtist *domainArtistModel.Artist) error
	GetArtist(ctx context.Context, artistID string) (*domainArtistModel.Artist, error)
	UpdateArtist(ctx context.Context, artistID string, updater domainArtistModel.ArtistUpdater) error
}
