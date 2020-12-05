package rabbitmq

import (
	"github.com/gin-gonic/gin"
	"pixstall-artist/app/domain/artist"
)

type HTTPArtistController struct {
	artistUseCase artist.UseCase
}

func NewHTTPArtistController(useCase artist.UseCase) HTTPArtistController {
	return HTTPArtistController {
		artistUseCase: useCase,
	}
}

func (a HTTPArtistController) GetArtist(c *gin.Context) {

}