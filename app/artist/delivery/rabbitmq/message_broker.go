package rabbitmq

import (
	"context"
	"pixstall-artist/app/domain/artist"
)

type ArtistMessageBroker struct {
	artistUseCase artist.UseCase
}

func NewRabbitMQArtistMessageBroker(useCase artist.UseCase) ArtistMessageBroker {
	return ArtistMessageBroker{
		artistUseCase: useCase,
	}
}

func (a ArtistMessageBroker) RegisterNewArtist(ctx context.Context) {

}
