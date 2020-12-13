package get_artist

import (
	domain "pixstall-artist/domain/artist/model"
)

type Response struct {
	domain.Artist
}

func NewResponse(dArtist domain.Artist) *Response {
	return &Response{
		dArtist,
	}
}
