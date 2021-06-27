package bookmark

import (
	"context"
	"pixstall-artist/domain/artist/model"
)

type UseCase interface {
	AddBookmark(ctx context.Context, userID string, artistID string) error
	GetBookmarksForUser(ctx context.Context, userID string, count int, offset int) (*[]model.Artist, error)
	GetBookmarkIDs(ctx context.Context, userID string) (*[]string, error)
	RemoveBookmark(ctx context.Context, userID string, artistID string) error
}