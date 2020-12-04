package artist

import "context"

type UseCase interface {
	RegisterNewArtist(ctx context.Context) error
}