package model

import "time"

type OpenCommission struct {
	ID                             string             `json:"id"`
	ArtistID                       string             `json:"artistId"`
	Title                          string             `json:"title"`
	Desc                           string             `json:"desc"`
	DepositRule                    *string            `json:"depositRule"`
	Price                          Price              `json:"price"`
	DayNeed                        DayNeed            `json:"dayNeed"`
	TimesAllowedDraftToChange      *int               `json:"timesAllowedDraftToChange"`
	TimesAllowedCompletionToChange *int               `json:"timesAllowedCompletionToChange"`
	SampleImagePaths               []string           `json:"sampleImagePaths"`
	State                          OpenCommissionSate `json:"state"`
	CreateTime                     time.Time          `json:"createTime"`
	LastUpdatedTime                time.Time          `json:"lastUpdatedTime"`
}

type Price struct {
	Amount   float64  `json:"amount" bson:"amount"`
	Currency Currency `json:"currency" bson:"currency"`
}

type DayNeed struct {
	From int `json:"from" bson:"from"`
	To   int `json:"to" bson:"to"`
}

type Size struct {
	Width  float64 `json:"width" bson:"width"`
	Height float64 `json:"height" bson:"height"`
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
