package model

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
	State                          *OpenCommissionSate
}
