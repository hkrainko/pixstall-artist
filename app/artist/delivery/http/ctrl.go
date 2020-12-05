package http

import (
	"github.com/gin-gonic/gin"
	"pixstall-artist/app/artist/delivery/model/get_artist"
	domainArtist "pixstall-artist/app/domain/artist"
)

type ArtistController struct {
	artistUseCase domainArtist.UseCase
}

func NewArtistController(useCase domainArtist.UseCase) ArtistController {
	return ArtistController{
		artistUseCase: useCase,
	}
}

func (a ArtistController) GetArtist(c *gin.Context) {
	artistID := c.Query("artistId")
	artist, err := a.artistUseCase.GetArtist(c, artistID)
	if err != nil {
		return
	}

	c.PureJSON(200, get_artist.NewResponse(*artist))
}