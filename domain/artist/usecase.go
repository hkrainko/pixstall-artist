package artist

import (
	"context"
	domainArtistModel "pixstall-artist/domain/artist/model"
	domainOpenCommissionModel "pixstall-artist/domain/open-commission/model"
	"pixstall-artist/domain/reg/model"
)

type UseCase interface {
	RegisterNewArtist(ctx context.Context, regInfo *model.RegInfo) error
	GetArtist(ctx context.Context, artistID string) (*domainArtistModel.Artist, error)
	GetArtistDetails(ctx context.Context, artistID string, requesterID *string) (*domainArtistModel.Artist, error)
	UpdateArtist(ctx context.Context, updater domainArtistModel.ArtistUpdater) (*string, error)

	// Open Commission
	GetOpenCommissionsForArtist(ctx context.Context, artistID string, requesterID *string, count int, offset int) (*domainOpenCommissionModel.GetOpenCommissionResult, error)
	AddOpenCommission(ctx context.Context, requesterID string, openCommCreator domainOpenCommissionModel.OpenCommissionCreator) (*string, error)
}
