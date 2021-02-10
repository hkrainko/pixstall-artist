package model

import (
	"time"
)

type User struct {
	UserID          string    `json:"userId" bson:"userId"`
	UserName        string    `json:"userName" bson:"userName"`
	ProfilePath     string    `json:"profilePath" bson:"profilePath"`
	Email           string    `json:"email,omitempty" bson:"email"`
	Birthday        string    `json:"birthday,omitempty" bson:"birthday"`
	Gender          string    `json:"gender,omitempty" bson:"gender"`
	State           UserState `json:"state" bson:"state"`
	RegTime         time.Time `json:"regTime" bson:"regTime"`
	LastUpdatedTime time.Time `json:"lastUpdatedTime,omitempty" bson:"lastUpdatedTime"`
}
