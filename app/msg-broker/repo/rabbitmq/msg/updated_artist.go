package msg

import "pixstall-artist/domain/artist/model"

type UpdatedArtist struct {
	ArtistID string

	UserName    *string
	ProfilePath *string
	Email       *string
	Birthday    *string
	Gender      *string
	State       *string
	RegTime     *string

	ArtistIntro    *model.ArtistIntro
	ArtistBoard    *model.ArtistBoard
	PaymentMethods *[]string
}
