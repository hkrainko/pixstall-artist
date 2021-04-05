package model

import "time"

type CommissionDetails struct {
	CommissionRequestCount int        `json:"commissionRequestCount" bson:"commissionRequestCount"`
	CommissionAcceptCount  int        `json:"commissionAcceptCount" bson:"commissionAcceptCount"`
	CommissionSuccessCount int        `json:"commissionSuccessCount" bson:"commissionSuccessCount"`
	AvgRatings             *float64   `json:"avgRatings,omitempty" bson:"avgRatings,omitempty"`
	LastRequestTime        *time.Time `json:"lastRequestTime,omitempty" bson:"lastRequestTime,omitempty"`
}
