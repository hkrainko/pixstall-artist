package model

import "time"

type CommissionDetails struct {
	CommissionRequestCount int        `bson:"commissionRequestCount"`
	CommissionAcceptCount  int        `bson:"commissionAcceptCount"`
	CommissionSuccessCount int        `bson:"commissionSuccessCount"`
	AvgRatings             *int       `bson:"avgRatings,omitempty"`
	LastRequestTime        *time.Time `bson:"lastRequestTime,omitempty"`
}

type CommissionDetailsUpdater struct {
	CommissionRequestCount *int
	CommissionAcceptCount  *int
	CommissionSuccessCount *int
	AvgRatings             *int
	LastRequestTime        *time.Time
}
