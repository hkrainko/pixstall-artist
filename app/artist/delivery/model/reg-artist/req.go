package reg_artist

import "pixstall-artist/domain/reg/model"

type Request struct {
	*model.RegInfo `json:",inline"`
}
