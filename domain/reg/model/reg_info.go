package model

import (
	"pixstall-artist/domain/artist/model"
	"time"
)

type RegInfo struct {
	AuthID         string            `json:"authId"`
	UserID         string            `json:"userId"`
	DisplayName    string            `json:"name"`
	Email          string            `json:"email"`
	Birthday       string            `json:"birthday"`
	Gender         string            `json:"gender"`
	ProfilePath    string            `json:"profilePath"`
	RegAsArtist    bool              `json:"regAsArtist"`
	RegArtistIntro model.ArtistIntro `json:"regArtistIntro"`
	RegTime        time.Time         `json:"regTime"`
}
