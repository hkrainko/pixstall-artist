package file

import (
	"context"
	"pixstall-artist/domain/file/model"
)

type Repo interface {
	SaveFile(ctx context.Context, file model.File, fileType model.FileType, ownerID string, acl []string) (*string, error)
}
