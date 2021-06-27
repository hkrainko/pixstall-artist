package model

import "pixstall-artist/domain/user/model"

type ArtistFilter struct {
	Count          int
	Offset         int
	BookmarkUserID *string
	State          *model.UserState
}
