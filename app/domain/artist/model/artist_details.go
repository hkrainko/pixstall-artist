package model

import "time"

type ArtistDetails struct {
	CommissionRequestCount int
	CommissionAcceptCount  int
	CommissionSuccessCount int
	AvgRatings             int
	LastRequestTime        time.Time
}
