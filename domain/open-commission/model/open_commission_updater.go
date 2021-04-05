package model

import "time"

type OpenCommissionUpdater struct {
	ID                             string
	Title                          *string
	Desc                           *string
	DepositRule                    *string
	Price                          *Price
	DayNeed                        *DayNeed
	TimesAllowedDraftToChange      *int
	TimesAllowedCompletionToChange *int
	SampleImagePaths               *[]string
	IsR18                          *bool
	AllowBePrivate                 *bool
	AllowAnonymous                 *bool
	State                          *OpenCommissionState
	LastUpdatedTime                *time.Time
}
