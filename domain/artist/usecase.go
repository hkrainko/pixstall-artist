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
	GetArtist(ctx context.Context, artistID string) (*domainArtistModel.Artist, error)
	UpdateBasicInfo(ctx context.Context, artistID string, updater *domainArtistModel.ArtistUpdater) error
	UpdateIntro(ctx context.Context, artistID string, updater *domainArtistModel.ArtistIntroUpdater) error
	UpdateDetails(ctx context.Context, artistID string, updater *domainArtistModel.ArtistDetailsUpdater) error
	UpdateOpenCommission(ctx context.Context, artistID string, updater *domainOpenCommissionModel.OpenCommissionUpdater) error
	UpdateArtwork(ctx context.Context, artistID string, updater *domainArtworkModel.ArtworkUpdater) error
	AddOpenCommission(ctx context.Context, openCommission *domainOpenCommissionModel.OpenCommission) error
	AddArtwork(ctx context.Context, artwork *domainArtworkModel.Artwork) error
	DeleteOpenCommission(ctx context.Context, artistID string, openCommissionID string) error
	DeleteArtwork(ctx context.Context, artistID string, artworkID string) error
}
