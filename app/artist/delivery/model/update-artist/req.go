package update_artist

import (
	model "pixstall-artist/domain/user/model"
)

type Request struct {
	*model.UserUpdater
}