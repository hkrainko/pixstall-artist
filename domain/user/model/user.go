package model

import (
	"pixstall-artist/domain/artist/model"
	"time"
)

type User struct {
	UserID          string          `json:"userId"`
	UserName        string          `json:"userName"`
	ProfilePath     string          `json:"profilePath"`
	Email           string          `json:"email,omitempty"`
	Birthday        string          `json:"birthday,omitempty"`
	Gender          string          `json:"gender,omitempty"`
	State           model.UserState `json:"state"`
	RegTime         time.Time       `json:"regTime"`
	LastUpdatedTime time.Time       `json:"lastUpdatedTime,omitempty"`
}
