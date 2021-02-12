package model

import "time"

type CommissionDetails struct {
	CommissionRequestCount int        `json:"commissionRequestCount" bson:"commissionRequestCount"`
	CommissionAcceptCount  int        `json:"commissionAcceptCount" bson:"commissionAcceptCount"`
	CommissionSuccessCount int        `json:"commissionSuccessCount" bson:"commissionSuccessCount"`
	AvgRatings             *int       `json:"avgRatings,omitempty" bson:"avgRatings,omitempty"`
	LastRequestTime        *time.Time `json:"lastRequestTime,omitempty" bson:"lastRequestTime,omitempty"`
}

type CommissionDetailsUpdater struct {
	CommissionRequestCount *int
	CommissionAcceptCount  *int
	CommissionSuccessCount *int
	AvgRatings             *int
	LastRequestTime        *time.Time
}
