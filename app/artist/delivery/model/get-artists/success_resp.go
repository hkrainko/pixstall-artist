package get_artist

import (
	domain "pixstall-artist/domain/artist/model"
)

type Response struct {
	Artists []domain.Artist `json:"artists"`
	Count   int             `json:"count"`
	Offset  int             `json:"offset"`
}

func NewResponse(dArtists []domain.Artist, count int, offset int) *Response {
	return &Response{
		Artists: dArtists,
		Count: count,
		Offset: offset,
	}
}
