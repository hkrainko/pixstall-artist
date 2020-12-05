package artist

import (
	"context"
	domainArtistModel "pixstall-artist/app/domain/artist/model"
)

type Repo interface {
	SaveArtist(ctx context.Context, dArtist *domainArtistModel.Artist) error
	GetArtist(ctx context.Context, artistID string) (*domainArtistModel.Artist, error)
	UpdateArtist(ctx context.Context, updater *domainArtistModel.ArtistUpdater) error
}
