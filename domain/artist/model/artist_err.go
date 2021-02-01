package model

import "strconv"

type ArtistError int

func (e ArtistError) Error() string {
	switch e {
	case ArtistErrorNotFound:
		return "ArtistErrorNotFound"
	case ArtistErrorUnAuth:
		return "ArtistErrorUnAuth"
	case ArtistErrorUnknown:
		return "ArtistErrorUnknown"
	default:
		return strconv.Itoa(int(e))
	}
}

const (
	ArtistErrorNotFound ArtistError = 10
	ArtistErrorUnAuth ArtistError = 11
	ArtistErrorUnknown ArtistError = 99
)