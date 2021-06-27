package artist

import (
	"context"
	domainArtistModel "pixstall-artist/domain/artist/model"
	"pixstall-artist/domain/fan/model"
)

type Repo interface {
	SaveArtist(ctx context.Context, dArtist *domainArtistModel.Artist) error
	GetArtist(ctx context.Context, artistID string) (*domainArtistModel.Artist, error)
	GetArtists(ctx context.Context, filter domainArtistModel.ArtistFilter, sorter domainArtistModel.ArtistSorter) (*[]domainArtistModel.Artist, error)
	GetArtistDetails(ctx context.Context, artistID string) (*domainArtistModel.Artist, error)
	UpdateArtist(ctx context.Context, updater *domainArtistModel.ArtistUpdater) error
	// bookmark
	AddBookmark(ctx context.Context, userID string, artistID string) error
	RemoveBookmark(ctx context.Context, userID string, artistID string) error
	// fan
	AddFan(ctx context.Context, artistID string, fan model.Fan) error
	RemoveFan(ctx context.Context, artistID string, fanId string) error
}
