package model

type OpenCommissionUpdater struct {
	ID        string
	ArtistID  string
	Title     *string
	Desc      *string
	PriceFrom *string
	PriceTo   *string
	DayNeed   *DayNeed
	Size      *Size
	State     *OpenCommissionSate
}