package image

import (
	"context"
	"pixstall-artist/domain/image/model"
)

type Repo interface {
	SaveImage(ctx context.Context, pathImage model.PathImage) (*string, error)
	SaveImages(ctx context.Context, pathImages []model.PathImage) ([]string, error)
}