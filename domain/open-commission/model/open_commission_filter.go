package model

type OpenCommissionFilter struct {
	ArtistID              *string
	Count                 *int
	Offset                *int
	Key                   *string
	PriceFrom             *float64
	PriceTo               *float64
	DayNeedGreaterOrEqual *int
	DayNeedLessOrEqual    *int
}
