package model

type Artist struct {
	ArtistID   string
	UserID     string //need?
	UserName   string
	Email      string
	Birthday   string
	Gender     string
	PhotoURL   string
	ArtistInfo ArtistInfo
	State      UserState
}
