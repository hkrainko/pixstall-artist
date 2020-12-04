package artist

import "context"

type Repo interface {
	SaveArtist(ctx context.Context) error
}