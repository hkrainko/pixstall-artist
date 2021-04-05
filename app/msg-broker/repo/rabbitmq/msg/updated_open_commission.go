package msg

import (
	"pixstall-artist/domain/open-commission/model"
	"time"
)

type UpdatedOpenCommission struct {
	ID                             string                     `json:"id"`
	Title                          *string                    `json:"title,omitempty"`
	Desc                           *string                    `json:"desc,omitempty"`
	DepositRule                    *string                    `json:"depositRule,omitempty"`
	Price                          *model.Price               `json:"price,omitempty"`
	DayNeed                        *model.DayNeed             `json:"dayNeed,omitempty"`
	TimesAllowedDraftToChange      *int                       `json:"timesAllowedDraftToChange,omitempty"`
	TimesAllowedCompletionToChange *int                       `json:"timesAllowedCompletionToChange,omitempty"`
	SampleImagePaths               *[]string                  `json:"sampleImagePaths,omitempty"`
	IsR18                          *bool                      `json:"isR18,omitempty"`
	AllowBePrivate                 *bool                      `json:"allowBePrivate,omitempty"`
	AllowAnonymous                 *bool                      `json:"allowAnonymous,omitempty"`
	State                          *model.OpenCommissionState `json:"state,omitempty"`
	LastUpdatedTime                *time.Time                 `json:"lastUpdatedTime,omitempty"`
}

func NewUpdatedOpenCommission(updater model.OpenCommissionUpdater) UpdatedOpenCommission {
	return UpdatedOpenCommission{
		ID:                             updater.ID,
		Title:                          updater.Title,
		Desc:                           updater.Desc,
		DepositRule:                    updater.DepositRule,
		Price:                          updater.Price,
		DayNeed:                        updater.DayNeed,
		TimesAllowedDraftToChange:      updater.TimesAllowedDraftToChange,
		TimesAllowedCompletionToChange: updater.TimesAllowedCompletionToChange,
		SampleImagePaths:               updater.SampleImagePaths,
		IsR18:                          updater.IsR18,
		AllowBePrivate:                 updater.AllowBePrivate,
		AllowAnonymous:                 updater.AllowAnonymous,
		State:                          updater.State,
		LastUpdatedTime:                updater.LastUpdatedTime,
	}
}
