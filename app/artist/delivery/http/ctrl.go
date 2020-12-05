package http

import (
	"pixstall-artist/app/domain/artist"
)

type RabbitMQArtistController struct {
	artistUseCase artist.UseCase
}

func NewRabbitMQArtistController(useCase artist.UseCase) RabbitMQArtistController {
	return RabbitMQArtistController{
		artistUseCase: useCase,
	}
}

func (a RabbitMQArtistController) RegisterNewArtist() {

}