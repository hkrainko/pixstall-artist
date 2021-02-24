package model

type OpenCommissionFilter struct {
	ArtistID              *string
	Count                 *int64
	Offset                *int64
	Key                   *string
	PriceFrom             *float64
	PriceTo               *float64
	DayNeedGreaterOrEqual *int
	DayNeedLessOrEqual    *int
}
