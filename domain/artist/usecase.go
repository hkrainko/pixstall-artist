package artist

import (
	"context"
	domainArtistModel "pixstall-artist/domain/artist/model"
	domainArtworkModel "pixstall-artist/domain/artwork/model"
	domainOpenCommissionModel "pixstall-artist/domain/open-commission/model"
	"pixstall-artist/domain/reg/model"
	model2 "pixstall-artist/domain/user/model"
)

type UseCase interface {
	RegisterNewArtist(ctx context.Context, regInfo *model.RegInfo) error
	GetArtist(ctx context.Context, artistID string) (*domainArtistModel.Artist, error)
	GetArtistDetails(ctx context.Context, artistID string, requesterID *string) (*domainArtistModel.Artist, error)
	UpdateArtist(ctx context.Context, updater domainArtistModel.ArtistUpdater) (*string, error)
	UpdateArtistUser(ctx context.Context, updater model2.UserUpdater) (*string, error)
	UpdateDetails(ctx context.Context, artistID string, updater *domainArtistModel.CommissionDetailsUpdater) error

	// Open Commission
	GetOpenCommissionsForArtist(ctx context.Context, artistID string, requesterID *string, count int64, offset int64) ([]domainOpenCommissionModel.OpenCommission, error)
	AddOpenCommission(ctx context.Context, requesterID string, openCommCreator domainOpenCommissionModel.OpenCommissionCreator) (*string, error)

	// Artwork
	UpdateArtwork(ctx context.Context, artistID string, updater *domainArtworkModel.ArtworkUpdater) error
	AddArtwork(ctx context.Context, artwork *domainArtworkModel.Artwork) error
	DeleteArtwork(ctx context.Context, artistID string, artworkID string) error
}
