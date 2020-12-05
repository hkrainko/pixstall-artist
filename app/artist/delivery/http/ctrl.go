package http

import (
	"github.com/gin-gonic/gin"
	"pixstall-artist/app/domain/artist"
)

type ArtistController struct {
	artistUseCase artist.UseCase
}

func NewArtistController(useCase artist.UseCase) ArtistController {
	return ArtistController{
		artistUseCase: useCase,
	}
}

func (a ArtistController) GetArtist(c *gin.Context) {

}