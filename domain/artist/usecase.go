package artist

import (
	"context"
	domainArtistModel "pixstall-artist/domain/artist/model"
	"pixstall-artist/domain/reg/model"
)

type UseCase interface {
	RegisterNewArtist(ctx context.Context, regInfo *model.RegInfo) error
	GetArtist(ctx context.Context, artistID string) (*domainArtistModel.Artist, error)
	GetArtists(ctx context.Context, filter domainArtistModel.ArtistFilter, sorter domainArtistModel.ArtistSorter) (*[]domainArtistModel.Artist, error)
	GetArtistDetails(ctx context.Context, artistID string, requesterID *string) (*domainArtistModel.Artist, error)
	UpdateArtist(ctx context.Context, updater domainArtistModel.ArtistUpdater) (*string, error)
}
