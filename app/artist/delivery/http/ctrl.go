package http

import (
	"github.com/gin-gonic/gin"
	"pixstall-artist/app/artist/delivery/model/get-artist"
	domainArtist "pixstall-artist/domain/artist"
	domain "pixstall-artist/domain/artist/model"
	"strconv"
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
	artistID := c.Param("id")
	artist, err := a.artistUseCase.GetArtist(c, artistID, nil)
	if err != nil {
		c.JSON(get_artist.NewErrorResponse(err))
		return
	}

	c.JSON(200, get_artist.NewResponse(*artist))
}

func (a ArtistController) GetArtistDetails(c *gin.Context) {
	artistID := c.Param("id")
	tokenUserID := c.GetString("userId")
	if artistID != tokenUserID {
		c.JSON(get_artist.NewErrorResponse(domain.ArtistErrorUnAuth))
		return
	}
	artist, err := a.artistUseCase.GetArtist(c, artistID, &tokenUserID)
	if err != nil {
		c.JSON(get_artist.NewErrorResponse(err))
		return
	}
	c.JSON(200, get_artist.NewResponse(*artist))
}

func (a ArtistController) UpdateArtist(c *gin.Context) {
	artistID := c.Param("id")
	tokenUserID := c.GetString("userId")
	if artistID != tokenUserID {
		c.JSON(get_artist.NewErrorResponse(domain.ArtistErrorUnAuth))
		return
	}

	updater := &domain.ArtistIntroUpdater{
		YearOfDrawing: nil,
		ArtTypes:      nil,
	}

	yearOfDrawing, exist := c.GetQuery("yearOfDrawing")
	if exist {
		if value, err := strconv.Atoi(yearOfDrawing); err == nil {
			updater.YearOfDrawing = &value
		}
	}
	artTypes, exist := c.GetQueryArray("artTypes")
	if exist {
		updater.ArtTypes = &artTypes
	}

	err := a.artistUseCase.UpdateIntro(c, artistID, updater)
	if err != nil {
		return
	}

	c.PureJSON(200, nil)
}

func (a ArtistController) GetOpenCommissionsForArtist(c *gin.Context) {

}

func (a ArtistController) GetOpenCommissionsDetailsForArtist(c *gin.Context) {

}

func (a ArtistController) AddOpenCommissionForArtist(c *gin.Context) {

}


