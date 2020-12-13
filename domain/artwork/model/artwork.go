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

type ArtworkUpdater struct {
	ID           string
	ArtistID     string
	ClientID     string
	Rating       *int
	RequestTime  *time.Time
	CompleteTime *time.Time
	State        *ArtworkState
}

type ArtworkState string

const (
	ArtworkStateActive    ArtworkState = "A"
	ArtworkStateHidden    ArtworkState = "H"
	ArtworkStateRemoved   ArtworkState = "R"
	ArtworkStateForbidden ArtworkState = "F"
)
