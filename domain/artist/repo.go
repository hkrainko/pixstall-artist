package artist

import (
	"context"
	domainArtistModel "pixstall-artist/domain/artist/model"
	domainArtworkModel "pixstall-artist/domain/artwork/model"
	domainOpenCommissionModel "pixstall-artist/domain/open-commission/model"
)

type Repo interface {
	SaveArtist(ctx context.Context, dArtist *domainArtistModel.Artist) error
	GetArtist(ctx context.Context, artistID string) (*domainArtistModel.Artist, error)
	UpdateArtist(ctx context.Context, updater *domainArtistModel.ArtistUpdater) error
	AddOpenCommission(ctx context.Context, openCommission *domainOpenCommissionModel.OpenCommission) error
	AddArtwork(ctx context.Context, artwork *domainArtworkModel.Artwork) error
}