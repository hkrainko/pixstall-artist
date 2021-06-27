package get_artist

import (
	get_artist "pixstall-artist/app/artist/delivery/model/get-artist"
	domain "pixstall-artist/domain/artist/model"
)

type Response struct {
	Artists []get_artist.Response `json:"artists"`
	Count   int             `json:"count"`
	Offset  int             `json:"offset"`
}

func NewResponse(dArtists []domain.Artist, count int, offset int) *Response {
	rArtists := make([]get_artist.Response, 0)
	for _, v := range dArtists {
		rArtists = append(rArtists, *get_artist.NewResponse(v))
	}

	return &Response{
		Artists: rArtists,
		Count: count,
		Offset: offset,
	}
}
