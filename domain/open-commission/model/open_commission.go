package model

import "time"

type OpenCommission struct {
	ID                   string
	ArtistID             string
	Title                string
	Desc                 string
	PriceRange           string
	DayNeed              DayNeed
	TimesAllowedToChange *int
	SampleImagePath      *string
	Size                 Size
	State                OpenCommissionSate
	CreateTime           time.Time
	LastUpdatedTime      time.Time
}

type PriceRange struct {
	AmountFrom float64
	AmountTo   float64
	Currency   Currency
}

type DayNeed struct {
	From int
	To   int
}

type Size struct {
	Width  float64
	Height float64
}

type Currency string

const (
	CurrencyHKD Currency = "HKD"
	CurrencyTWD Currency = "TWD"
	CurrencyUSE Currency = "USD"
)

type OpenCommissionSate string

const (
	OpenCommissionStateActive  OpenCommissionSate = "A"
	OpenCommissionStateHidden  OpenCommissionSate = "H"
	OpenCommissionStateRemoved OpenCommissionSate = "R"
)
