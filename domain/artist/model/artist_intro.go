package model

type ArtistIntro struct {
	YearOfDrawing int      `json:"yearOfDrawing" bson:"yearOfDrawing"`
	ArtTypes      []string `json:"artTypes" bson:"artTypes"`
}

type ArtistIntroUpdater struct {
	YearOfDrawing *int      `json:"yearOfDrawing" bson:"yearOfDrawing"`
	ArtTypes      *[]string `json:"artTypes" bson:"artTypes"`
}
