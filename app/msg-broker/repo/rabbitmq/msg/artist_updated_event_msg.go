package msg

import (
	model2 "pixstall-artist/domain/artist/model"
	"pixstall-artist/domain/user/model"
	"time"
)

type ArtistUpdatedEventMsg struct {
	ArtistID string `json:"artistId"`

	UserName    *string          `json:"userName,omitempty"`
	ProfilePath *string          `json:"profilePath,omitempty"`
	Email       *string          `json:"email,omitempty"`
	Birthday    *string          `json:"birthday,omitempty"`
	Gender      *string          `json:"gender,omitempty"`
	State       *model.UserState `json:"state,omitempty"`

	PaymentMethods *[]string `json:"paymentMethods,omitempty"`

	// ArtistIntro
	YearOfDrawing *int      `json:"yearOfDrawing,omitempty"`
	ArtTypes      *[]string `json:"artTypes,omitempty"`

	// ArtistBoard
	BannerPath *string `json:"bannerPath,omitempty"`
	Desc       *string `json:"desc,omitempty"`

	// CommissionDetails
	CommissionRequestCount *int       `json:"commissionRequestCount,omitempty"`
	CommissionAcceptCount  *int       `json:"commissionAcceptCount,omitempty"`
	CommissionSuccessCount *int       `json:"commissionSuccessCount,omitempty"`
	AvgRatings             *float64   `json:"avgRatings,omitempty"`
	LastRequestTime        *time.Time `json:"lastRequestTime,omitempty"`
	LastUpdatedTime        *time.Time `json:"lastUpdatedTime,omitempty"`
}

func NewArtistUpdatedEventMsg(updater model2.ArtistUpdater) ArtistUpdatedEventMsg {
	return ArtistUpdatedEventMsg{
		ArtistID:               updater.ArtistID,
		UserName:               updater.UserName,
		ProfilePath:            updater.ProfilePath,
		Email:                  updater.Email,
		Birthday:               updater.Birthday,
		Gender:                 updater.Gender,
		State:                  updater.State,
		PaymentMethods:         updater.PaymentMethods,
		YearOfDrawing:          updater.YearOfDrawing,
		ArtTypes:               updater.ArtTypes,
		BannerPath:             updater.BannerPath,
		Desc:                   updater.Desc,
		CommissionRequestCount: updater.CommissionRequestCount,
		CommissionAcceptCount:  updater.CommissionAcceptCount,
		CommissionSuccessCount: updater.CommissionSuccessCount,
		AvgRatings:             updater.AvgRatings,
		LastRequestTime:        updater.LastRequestTime,
		LastUpdatedTime:        updater.LastUpdatedTime,
	}
}
