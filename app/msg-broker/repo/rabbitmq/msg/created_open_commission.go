package msg

import (
	"pixstall-artist/domain/open-commission/model"
	"time"
)

type CreatedOpenCommission struct {
	ID                             string                    `json:"id"`
	ArtistID                       string                    `json:"artistId"`
	Title                          string                    `json:"title"`
	Desc                           string                    `json:"desc"`
	DepositRule                    *string                   `json:"depositRule"`
	Price                          model.Price               `json:"price"`
	DayNeed                        model.DayNeed             `json:"dayNeed"`
	TimesAllowedDraftToChange      *int                      `json:"timesAllowedDraftToChange"`
	TimesAllowedCompletionToChange *int                      `json:"timesAllowedCompletionToChange"`
	SampleImagePaths               []string                  `json:"sampleImagePaths"`
	IsR18                          bool                      `json:"isR18"`
	AllowBePrivate                 bool                      `json:"allowBePrivate"`
	AllowAnonymous                 bool                      `json:"allowAnonymous"`
	State                          model.OpenCommissionState `json:"state"`
	CreateTime                     time.Time                 `json:"createTime"`
	LastUpdatedTime                time.Time                 `json:"lastUpdatedTime"`
}

func NewCreatedOpenCommission(openComm model.OpenCommission) CreatedOpenCommission {
	return CreatedOpenCommission{
		ID:                             openComm.ID,
		ArtistID:                       openComm.ArtistID,
		Title:                          openComm.Title,
		Desc:                           openComm.Desc,
		DepositRule:                    openComm.DepositRule,
		Price:                          openComm.Price,
		DayNeed:                        openComm.DayNeed,
		TimesAllowedDraftToChange:      openComm.TimesAllowedDraftToChange,
		TimesAllowedCompletionToChange: openComm.TimesAllowedCompletionToChange,
		SampleImagePaths:               openComm.SampleImagePaths,
		IsR18:                          openComm.IsR18,
		AllowBePrivate:                 openComm.AllowBePrivate,
		AllowAnonymous:                 openComm.AllowAnonymous,
		State:                          openComm.State,
		CreateTime:                     openComm.CreateTime,
		LastUpdatedTime:                openComm.LastUpdatedTime,
	}
}
