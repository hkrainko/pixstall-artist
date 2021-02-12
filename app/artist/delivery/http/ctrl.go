package http

import (
	"github.com/gin-gonic/gin"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	add_open_commission_for_artist "pixstall-artist/app/artist/delivery/model/add-open-commission-for-artist"
	"pixstall-artist/app/artist/delivery/model/get-artist"
	domainArtist "pixstall-artist/domain/artist"
	domain "pixstall-artist/domain/artist/model"
	"pixstall-artist/domain/open-commission/model"
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
	artist, err := a.artistUseCase.GetArtist(c, artistID)
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
	artist, err := a.artistUseCase.GetArtistDetails(c, artistID, &tokenUserID)
	if err != nil {
		c.JSON(get_artist.NewErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, get_artist.NewResponse(*artist))
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

	c.PureJSON(http.StatusOK, nil)
}

func (a ArtistController) GetOpenCommissionsForArtist(c *gin.Context) {

}

func (a ArtistController) GetOpenCommissionsDetailsForArtist(c *gin.Context) {

}

func (a ArtistController) AddOpenCommissionForArtist(c *gin.Context) {
	artistID := c.Param("id")
	tokenUserID := c.GetString("userId")
	if artistID != tokenUserID {
		c.JSON(add_open_commission_for_artist.NewErrorResponse(domain.ArtistErrorUnAuth))
		return
	}
	title, exist := c.GetPostForm("title")
	if !exist {
		c.JSON(http.StatusBadRequest, nil)
	}
	desc, exist := c.GetPostForm("desc")
	if !exist {
		c.JSON(http.StatusBadRequest, nil)
	}
	dayNeedFrom, exist := c.GetPostForm("dayNeed.from")
	if !exist {
		c.JSON(http.StatusBadRequest, nil)
	}
	dayNeedFromInt, err := strconv.Atoi(dayNeedFrom)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	dayNeedTo, exist := c.GetPostForm("dayNeed.to")
	if !exist {
		c.JSON(http.StatusBadRequest, nil)
	}
	dayNeedToInt, err := strconv.Atoi(dayNeedTo)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	creator := model.OpenCommissionCreator{
		Title:       title,
		Desc:        desc,
		DayNeed: model.DayNeed{
			From: dayNeedFromInt,
			To:   dayNeedToInt,
		},
	}
	depositRule, exist := c.GetPostForm("depositRule")
	if exist {
		creator.DepositRule = &depositRule
	}
	if priceAmount, exist := c.GetPostForm("price.amount"); exist {
		if amount, err := strconv.ParseFloat(priceAmount, 64); err == nil {
			if priceCurrency, exist := c.GetPostForm("price.currency"); exist {
				creator.Price = model.Price{
					Amount:   amount,
					Currency: model.Currency(priceCurrency),
				}
			}
		}
	}
	timesAllowedDraftToChange, exist := c.GetPostForm("timesAllowedDraftToChange")
	if exist {
		if i, err := strconv.Atoi(timesAllowedDraftToChange); err == nil {
			creator.TimesAllowedDraftToChange = &i
		}
	}
	timesAllowedCompletionToChange, exist := c.GetPostForm("timesAllowedCompletionToChange")
	if exist {
		if i, err := strconv.Atoi(timesAllowedCompletionToChange); err == nil {
			creator.TimesAllowedCompletionToChange = &i
		}
	}

	form, err := c.MultipartForm()
	if err == nil {
		fileHeaders := form.File["sampleImages"]
		images := make([]image.Image, 0)
		for _, header := range fileHeaders {
			decodedImage := func() image.Image {
				if err != nil {
					return nil
				}
				f, err := header.Open()
				if err != nil {
					return nil
				}
				decodedImage, _, err := image.Decode(f)
				if err != nil {
					return nil
				}
				return decodedImage
			}()
			if decodedImage != nil {
				images = append(images, decodedImage)
			}
		}
		if len(images) > 0 {
			creator.SampleImages = images
		}
	}

	id, err := a.artistUseCase.AddOpenCommission(c, artistID, creator)
	if err != nil {
		add_open_commission_for_artist.NewErrorResponse(err)
		return
	}
	c.JSON(http.StatusOK, add_open_commission_for_artist.NewResponse(*id))
}


