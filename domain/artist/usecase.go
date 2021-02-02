package artist

import (
	"context"
	domainArtistModel "pixstall-artist/domain/artist/model"
	domainArtworkModel "pixstall-artist/domain/artwork/model"
	domainOpenCommissionModel "pixstall-artist/domain/open-commission/model"
	"pixstall-artist/domain/reg/model"
)

type UseCase interface {
	RegisterNewArtist(ctx context.Context, regInfo *model.RegInfo) error
	GetArtist(ctx context.Context, artistID string, requesterID *string) (*domainArtistModel.Artist, error)
	// Open Commission
	GetOpenCommissionsForArtist(ctx context.Context, artistID string, requesterID *string, count int, offset int) ([]domainOpenCommissionModel.OpenCommission, error)
	AddOpenCommission(ctx context.Context, requesterID string, openCommission *domainOpenCommissionModel.OpenCommission) error

	UpdateBasicInfo(ctx context.Context, artistID string, updater *domainArtistModel.ArtistUpdater) error
	UpdateIntro(ctx context.Context, artistID string, updater *domainArtistModel.ArtistIntroUpdater) error
	UpdateDetails(ctx context.Context, artistID string, updater *domainArtistModel.CommissionDetailsUpdater) error

	// Artwork
	UpdateArtwork(ctx context.Context, artistID string, updater *domainArtworkModel.ArtworkUpdater) error
	AddArtwork(ctx context.Context, artwork *domainArtworkModel.Artwork) error
	DeleteArtwork(ctx context.Context, artistID string, artworkID string) error
}
