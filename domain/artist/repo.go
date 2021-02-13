package artist

import (
	"context"
	domainArtistModel "pixstall-artist/domain/artist/model"
	domainArtworkModel "pixstall-artist/domain/artwork/model"
	"pixstall-artist/domain/fan/model"
	model2 "pixstall-artist/domain/user/model"
)

type Repo interface {
	SaveArtist(ctx context.Context, dArtist *domainArtistModel.Artist) error
	GetArtist(ctx context.Context, artistID string) (*domainArtistModel.Artist, error)
	GetArtistDetails(ctx context.Context, artistID string) (*domainArtistModel.Artist, error)
	UpdateArtist(ctx context.Context, updater *domainArtistModel.ArtistUpdater) error
	UpdateArtistUser(ctx context.Context, updater *model2.UserUpdater) error
	// Artwork
	AddArtwork(ctx context.Context, artwork *domainArtworkModel.Artwork) error
	// Fan
	AddFan(ctx context.Context, artistID string, fan model.Fan) error
	RemoveFan(ctx context.Context, artistID string, fanId string) error
}
