package model

import (
	"pixstall-artist/domain/open-commission/model"
)

type CommissionCreator struct {
	OpenCommissionID string `json:"openCommissionId"`
	ArtistID         string `json:"artistID"`
	//ArtistName           string      `json:"artistName"`
	//ArtistProfilePath    *string     `json:"artistProfilePath"`
	RequesterID string `json:"requesterID"`
	//RequesterName        string      `json:"requesterName"`
	//RequesterProfilePath *string     `json:"requesterProfilePath"`
	Price         model.Price `json:"price"`
	DayNeed       int         `json:"dayNeed"`
	Size          *model.Size `json:"size"`
	Resolution    *float64    `json:"resolution"`
	ExportFormat  *string     `json:"exportFormat"`
	Desc          string      `json:"desc"`
	PaymentMethod string      `json:"paymentMethod"`
	IsR18         bool        `json:"isR18"`
	BePrivate     bool        `json:"bePrivate"`
	Anonymous     bool        `json:"anonymous"`
	//RefImages            []image.Image
	RefImagePaths []string `json:"refImagePaths"`
}
