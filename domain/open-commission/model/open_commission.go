package model

import "time"

type OpenCommission struct {
	ID                             string
	ArtistID                       string
	Title                          string
	Desc                           string
	DepositRule                    *string
	Price                          Price
	DayNeed                        DayNeed
	TimesAllowedDraftToChange      *int
	TimesAllowedCompletionToChange *int
	SampleImagePaths               []string
	State                          OpenCommissionSate
	CreateTime                     time.Time
	LastUpdatedTime                time.Time
}

type Price struct {
	Amount   float64  `bson:"amount"`
	Currency Currency `bson:"currency"`
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
