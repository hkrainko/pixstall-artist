package model

type OpenCommission struct {
	ID        string
	ArtistID  string
	Title     string
	Desc      string
	PriceFrom string
	PriceTo   string
	DayNeed   DayNeed
	Size      Size
	State     OpenCommissionSate
}

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

type DayNeed struct {
	From int
	To   int
}

type Size struct {
	Width  float64
	Height float64
}

type OpenCommissionSate string

const (
	OpenCommissionStateActive  OpenCommissionSate = "A"
	OpenCommissionStateRemoved OpenCommissionSate = "R"
)
