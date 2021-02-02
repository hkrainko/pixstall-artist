package model

import (
	"pixstall-artist/domain/user/model"
	"time"
)

type Fan struct {
	model.User
	FollowTime time.Time
}
