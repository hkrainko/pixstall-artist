package model

import "time"

type Artwork struct {
	ID           string
	ArtistID     string
	ClientID     string
	Rating       int
	RequestTime  time.Time
	CompleteTime time.Time
	State        ArtworkState
}

type ArtworkState string
const (
	ArtworkStateActive    = "A"
	ArtworkStateHidden    = "H"
	ArtworkStateRemoved   = "R"
	ArtworkStateForbidden = "F"
)
