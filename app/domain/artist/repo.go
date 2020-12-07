package artist

import (
	"context"
	domainArtistModel "pixstall-artist/app/domain/artist/model"
	domainArtworkModel "pixstall-artist/app/domain/artwork/model"
	domainOpenCommissionModel "pixstall-artist/app/domain/open-commission/model"
)

type Repo interface {
	SaveArtist(ctx context.Context, dArtist *domainArtistModel.Artist) error
	GetArtist(ctx context.Context, artistID string) (*domainArtistModel.Artist, error)
	UpdateArtist(ctx context.Context, updater *domainArtistModel.ArtistUpdater) error
	AddOpenCommission(ctx context.Context, openCommission *domainOpenCommissionModel.OpenCommission) error
	AddArtwork(ctx context.Context, artwork *domainArtworkModel.Artwork) error
}
