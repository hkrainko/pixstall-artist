package http

import (
	"github.com/gin-gonic/gin"
	"pixstall-artist/app/artist/delivery/model/get_artist"
	domainArtist "pixstall-artist/app/domain/artist"
	domain "pixstall-artist/app/domain/artist/model"
	"pixstall-artist/app/domain/open-commission/model"
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
	artistID := c.Query("artistId")
	artist, err := a.artistUseCase.GetArtist(c, artistID)
	if err != nil {
		return
	}

	c.PureJSON(200, get_artist.NewResponse(*artist))
}

func (a ArtistController) UpdateIntro(c *gin.Context) {
	artistID := c.Query("artistId")

	updater := &domain.ArtistIntroUpdater{
		YearOfDrawing: nil,
		ArtTypes:      nil,
		SelfIntro:     nil,
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
	selfIntro, exist := c.GetQuery("selfIntro")
	if exist {
		updater.SelfIntro = &selfIntro
	}

	err := a.artistUseCase.UpdateIntro(c, artistID, updater)
	if err != nil {
		return
	}

	c.PureJSON(200, nil)
}

func (a ArtistController) UpdateOpenCommission(c *gin.Context) {
	artistID := c.Query("artistId")
	openCommissionID := c.Query("openCommissionId")
	title := c.Query("title")
	desc := c.Query("desc")
	priceFrom := c.Query("priceFrom")
	priceTo := c.Query("priceTo")

	dayNeedMap := c.QueryMap("dayNeed")
	var dayNeed *model.DayNeed
	if dayNeedMap["from"] != "" && dayNeedMap["to"] != "" {
		dayNeed = &model.DayNeed{}
		if from, err := strconv.Atoi(dayNeedMap["from"]); err == nil {
			dayNeed.From = from
		}
		if to, err := strconv.Atoi(dayNeedMap["to"]); err == nil {
			dayNeed.To = to
		}
	}

	sizeMap := c.QueryMap("size")
	var size *model.Size
	if sizeMap["width"] != "" && sizeMap["height"] != "" {
		size = &model.Size{}
		if width, err := strconv.ParseFloat(sizeMap["width"], 64); err == nil {
			size.Width = width
		}
		if height, err := strconv.ParseFloat(sizeMap["height"], 64); err == nil {
			size.Height = height
		}
	}

	updater := &model.OpenCommissionUpdater{
		ID:        openCommissionID,
		ArtistID:  artistID,
		Title:     &title,
		Desc:      &desc,
		PriceFrom: &priceFrom,
		PriceTo:   &priceTo,
		DayNeed:   dayNeed,
		Size:      size,
	}

	err := a.artistUseCase.UpdateOpenCommission(c, artistID, updater)
	if err != nil {
		return
	}

	c.PureJSON(200, nil)
}

func (a ArtistController) AddOpenCommission(c *gin.Context) {

}

func (a ArtistController) DeleteOpenCommission(c *gin.Context) {

}
