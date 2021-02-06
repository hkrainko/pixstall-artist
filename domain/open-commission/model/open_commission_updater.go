package model

type OpenCommissionUpdater struct {
	ID                             string
	ArtistID                       string
	Title                          *string
	Desc                           *string
	DepositRule                    *string
	Price                          *Price
	DayNeed                        *DayNeed
	TimesAllowedDraftToChange      *int
	TimesAllowedCompletionToChange *int
	SampleImagePaths               *[]string
	State                          *OpenCommissionSate
}
