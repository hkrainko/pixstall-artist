package msg

import (
	"pixstall-artist/domain/artist/model"
	model2 "pixstall-artist/domain/user/model"
)

type CreatedArtist struct {
	model2.User    `json:",inline"`
	ArtistID       string            `json:"artistId"`
	ArtistIntro    model.ArtistIntro `json:"artistIntro"`
	PaymentMethods []string          `json:"paymentMethods"`
}
