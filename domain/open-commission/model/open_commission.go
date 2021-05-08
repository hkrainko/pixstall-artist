package model

import "time"

type OpenCommission struct {
	ID                             string              `json:"id" bson:"id"`
	ArtistID                       string              `json:"artistId" bson:"artistId"`
	Title                          string              `json:"title" bson:"title"`
	Desc                           string              `json:"desc" bson:"desc"`
	DepositRule                    *string             `json:"depositRule" bson:"depositRule"`
	Price                          Price               `json:"price" bson:"price"`
	DayNeed                        DayNeed             `json:"dayNeed" bson:"dayNeed"`
	TimesAllowedDraftToChange      *int                `json:"timesAllowedDraftToChange" bson:"timesAllowedDraftToChange"`
	TimesAllowedCompletionToChange *int                `json:"timesAllowedCompletionToChange" bson:"timesAllowedCompletionToChange"`
	SampleImagePaths               []string            `json:"sampleImagePaths" bson:"sampleImagePaths"`
	IsR18                          bool                `json:"isR18" bson:"isR18"`
	AllowBePrivate                 bool                `json:"allowBePrivate" bson:"allowBePrivate"`
	AllowAnonymous                 bool                `json:"allowAnonymous" bson:"allowAnonymous"`
	State                          OpenCommissionState `json:"state" bson:"state"`
	CreateTime                     time.Time           `json:"createTime" bson:"createTime"`
	LastUpdatedTime                time.Time           `json:"lastUpdatedTime" bson:"lastUpdatedTime"`
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
	Unit   string  `json:"unit" bson:"unit"`
}

type Currency string

const (
	CurrencyHKD Currency = "HKD"
	CurrencyTWD Currency = "TWD"
	CurrencyUSE Currency = "USD"
)

type OpenCommissionState string

const (
	OpenCommissionStateActive  OpenCommissionState = "A"
	OpenCommissionStateHidden  OpenCommissionState = "H"
	OpenCommissionStateRemoved OpenCommissionState = "R"
)
