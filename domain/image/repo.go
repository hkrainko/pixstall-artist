package image

import (
	"context"
	"pixstall-artist/domain/image/model"
)

type Repo interface {
	SaveImage(ctx context.Context, pathImage model.PathImage) error
	SaveImages(ctx context.Context, pathImages []model.PathImage) ([]string, error)
}