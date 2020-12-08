package model

import (
	"pixstall-artist/app/domain/user/model"
	"time"
)

type Fan struct {
	model.User
	FollowTime time.Time
}
