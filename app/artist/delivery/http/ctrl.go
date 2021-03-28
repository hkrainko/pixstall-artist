package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"net/http"
	add_open_commission_for_artist "pixstall-artist/app/artist/delivery/model/add-open-commission-for-artist"
	"pixstall-artist/app/artist/delivery/model/get-artist"
	domainArtist "pixstall-artist/domain/artist"
	domain "pixstall-artist/domain/artist/model"
	model2 "pixstall-artist/domain/file/model"
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

	updater := domain.ArtistUpdater{
		ArtistID: artistID,
	}

	yearOfDrawing, exist := c.GetPostForm("artistIntro.yearOfDrawing")
	if exist {
		if value, err := strconv.Atoi(yearOfDrawing); err == nil {
			if updater.ArtistIntro == nil {
				artistIntro := domain.ArtistIntroUpdater{}
				updater.ArtistIntro = &artistIntro
			}
			updater.ArtistIntro.YearOfDrawing = &value
		}
	}
	artTypes, exist := c.GetPostFormArray("artistIntro.artTypes")
	if exist {
		if updater.ArtistIntro == nil {
			artistIntro := domain.ArtistIntroUpdater{}
			updater.ArtistIntro = &artistIntro
		}
		updater.ArtistIntro.ArtTypes = &artTypes
	}
	imageFiles, err := getMultipartFormImages(c, "artistBoard.bannerImage")
	if err == nil {
		imgFiles := *imageFiles
		updater.ArtistBoard.BannerFile = &imgFiles[0]
	}
	desc, exist := c.GetPostForm("artistBoard.desc")
	if exist {
		if updater.ArtistBoard == nil {
			artistBoard := domain.ArtistBoardUpdater{}
			updater.ArtistBoard = &artistBoard
		}
		updater.ArtistBoard.Desc = &desc
	}
	artistId, err := a.artistUseCase.UpdateArtist(c, updater)
	if err != nil {
		c.PureJSON(http.StatusBadRequest, nil)
		return
	}
	c.PureJSON(http.StatusOK, artistId)
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
	creator := model.OpenCommissionCreator{}
	if title, exist := c.GetPostForm("title"); exist {
		creator.Title = title
	} else {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	if desc, exist := c.GetPostForm("desc"); exist {
		creator.Desc = desc
	} else {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	if dayNeedFrom, exist := c.GetPostForm("dayNeed.from"); exist {
		if dayNeedFromInt, err := strconv.Atoi(dayNeedFrom); err == nil {
			creator.DayNeed.From = dayNeedFromInt
		} else {
			c.JSON(http.StatusBadRequest, nil)
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	if dayNeedTo, exist := c.GetPostForm("dayNeed.to"); exist {
		if dayNeedToInt, err := strconv.Atoi(dayNeedTo); err == nil {
			creator.DayNeed.To = dayNeedToInt
		} else {
			c.JSON(http.StatusBadRequest, nil)
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	if isR18, exist := c.GetPostForm("isR18"); exist {
		creator.IsR18 = isR18 == "true"
	} else {
		c.JSON(http.StatusBadRequest, nil)
	}
	if allowBePrivate, exist := c.GetPostForm("allowBePrivate"); exist {
		creator.AllowBePrivate = allowBePrivate == "true"
	} else {
		c.JSON(http.StatusBadRequest, nil)
	}
	if allowAnonymous, exist := c.GetPostForm("allowAnonymous"); exist {
		creator.AllowAnonymous = allowAnonymous == "true"
	} else {
		c.JSON(http.StatusBadRequest, nil)
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
	imageFiles, err := getMultipartFormImages(c, "sampleImages")
	if err == nil {
		creator.SampleImages = *imageFiles
	}

	id, err := a.artistUseCase.AddOpenCommission(c, artistID, creator)
	if err != nil {
		add_open_commission_for_artist.NewErrorResponse(err)
		return
	}
	c.JSON(http.StatusOK, add_open_commission_for_artist.NewResponse(*id))
}

func getMultipartFormImages(ctx *gin.Context, key string) (*[]model2.ImageFile, error) {
	form, err := ctx.MultipartForm()
	if err != nil {
		return nil, err
	}
	fileHeaders := form.File[key]
	imageFiles := make([]model2.ImageFile, 0)
	for _, header := range fileHeaders {
		f, err := header.Open()
		if err != nil {
			continue
		}
		contentType, err := getFileContentType(f)
		if err != nil {
			_ = f.Close()
			continue
		}
		_, err = f.Seek(0, 0)
		if err != nil {
			_ = f.Close()
			continue
		}
		img, _, err := image.Decode(f)
		if err != nil {
			_ = f.Close()
			continue
		}
		_, err = f.Seek(0, 0)
		if err != nil {
			_ = f.Close()
			continue
		}
		imgF := model2.ImageFile{
			File: model2.File{
				File:        f,
				Name:        header.Filename,
				ContentType: contentType,
				Volume:      header.Size,
			},
			Size: model2.Size{
				Width:  float64(img.Bounds().Dx()),
				Height: float64(img.Bounds().Dy()),
				Unit:   "px",
			},
		}
		imageFiles = append(imageFiles, imgF)
		_ = f.Close()
	}
	if len(imageFiles) <= 0 {
		return nil, errors.New("")
	}
	return &imageFiles, nil
}

func getFileContentType(out multipart.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}