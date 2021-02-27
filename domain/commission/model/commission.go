package model

import "pixstall-artist/domain/open-commission/model"

type Commission struct {
	ID               string      `json:"id"`
	OpenCommissionID string      `json:"openCommissionId"`
	ArtistID         string      `json:"artistID"`
	RequesterID      string      `json:"requesterID"`
	Price            model.Price `json:"price"`
	DayNeed          int         `json:"dayNeed"`
	BePrivate        bool        `json:"bePrivate"`
	Anonymous        bool        `json:"anonymous"`
}
