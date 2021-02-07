package model

import "pixstall-artist/domain/fan/model"

type Fans struct {
	Meta  map[string]model.Fan `json:"meta"`
	Total int64                `json:"total"`
}
