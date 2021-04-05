package model

import (
	"pixstall-artist/domain/file/model"
	model2 "pixstall-artist/domain/user/model"
	"time"
)

type ArtistUpdater struct {
	ArtistID    string
	UserName    *string
	ProfilePath *string
	Email       *string
	Birthday    *string
	Gender      *string
	State       *model2.UserState
	RegTime     *string

	PaymentMethods *[]string

	// ArtistIntro
	YearOfDrawing *int
	ArtTypes      *[]string

	// ArtistBoard
	BannerFile *model.ImageFile
	BannerPath *string
	Desc       *string

	// CommissionDetails
	CommissionRequestCount *int
	CommissionAcceptCount  *int
	CommissionSuccessCount *int
	AvgRatings             *float64
	LastRequestTime        *time.Time

	LastUpdatedTime *time.Time
}
