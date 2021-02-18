package model

import (
	"image"
	"pixstall-artist/domain/open-commission/model"
)

type CommissionCreator struct {
	OpenCommissionID string      `json:"openCommissionId"`
	ArtistID         string      `json:"artistID"`
	RequesterID      string      `json:"requesterID"`
	Price            model.Price `json:"price"`
	DayNeed          int         `json:"dayNeed"`
	Size             *model.Size `json:"size"`
	Resolution       *float64    `json:"resolution"`
	ExportFormat     *string     `json:"exportFormat"`
	Desc             string      `json:"desc"`
	PaymentMethod    string      `json:"paymentMethod"`
	IsR18            bool        `json:"isR18"`
	BePrivate        bool        `json:"bePrivate"`
	Anonymous        bool        `json:"anonymous"`
	RefImages        []image.Image
	RefImagePaths    []string `json:"refImagePaths"`
}
