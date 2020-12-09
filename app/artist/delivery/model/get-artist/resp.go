package get_artist

import (
	domain "pixstall-artist/app/domain/artist/model"
)

type Response struct {
	domain.Artist
}

func NewResponse(dArtist domain.Artist) *Response {
	return &Response{
		dArtist,
	}
}
