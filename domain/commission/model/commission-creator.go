package model

import (
	"image"
	"pixstall-artist/domain/open-commission/model"
)

type CommissionCreator struct {
	OpenCommissionID string
	ArtistID         string
	RequesterID      string
	Price            model.Price
	DayNeed          int
	Size             *model.Size
	Resolution       *float64
	ExportFormat     *string
	Desc             string
	PaymentMethod    string
	IsR18            bool
	BePrivate        bool
	Anonymous        bool
	RegImages        []image.Image
	RegImagePaths    []string
}
