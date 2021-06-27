package usecase

import (
	"context"
	"pixstall-artist/domain/artist"
	"pixstall-artist/domain/artist/model"
	"pixstall-artist/domain/bookmark"
	msgBroker "pixstall-artist/domain/msg-broker"
	model2 "pixstall-artist/domain/user/model"
)

type bookmarkUseCase struct {
	artistRepo    artist.Repo
	msgBrokerRepo msgBroker.Repo
}

func NewBookmarkUseCase(artistRepo artist.Repo, msgBrokerRepo msgBroker.Repo) bookmark.UseCase {
	return bookmarkUseCase{
		artistRepo:    artistRepo,
		msgBrokerRepo: msgBrokerRepo,
	}
}

func (b bookmarkUseCase) AddBookmark(ctx context.Context, userID string, artistID string) error {
	return b.artistRepo.AddBookmark(ctx, userID, artistID)
}

func (b bookmarkUseCase) GetBookmarkIDs(ctx context.Context, userID string) (*[]string, error) {
	return b.artistRepo.GetBookmarkIDs(ctx, userID)
}

func (b bookmarkUseCase) GetBookmarksForUser(ctx context.Context, userID string, count int, offset int) (*[]model.Artist, error) {
	active := model2.UserStateActive
	filter := model.ArtistFilter{
		Count:          count,
		Offset:         offset,
		BookmarkUserID: &userID,
		State:          &active,
	}
	sorter := model.ArtistSorter{
		RegTime: nil,
	}
	return b.artistRepo.GetArtists(ctx, filter, sorter)
}

func (b bookmarkUseCase) RemoveBookmark(ctx context.Context, userID string, artistID string) error {
	return b.artistRepo.RemoveBookmark(ctx, userID, artistID)
}
