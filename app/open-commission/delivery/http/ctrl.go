package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	http2 "pixstall-artist/app/error/http"
	get_open_commissions "pixstall-artist/app/open-commission/delivery/http/model/get-open-commissions"
	error2 "pixstall-artist/domain/error"
	domainOpenComm "pixstall-artist/domain/open-commission"
	"pixstall-artist/domain/open-commission/model"
	"strconv"
)

type OpenCommissionController struct {
	openCommUseCase domainOpenComm.UseCase
}

func NewOpenCommissionController(useCase domainOpenComm.UseCase) OpenCommissionController {
	return OpenCommissionController{
		openCommUseCase: useCase,
	}
}

func (o OpenCommissionController) GetOpenCommission(c *gin.Context) {

}

func (o OpenCommissionController) GetOpenCommissions(c *gin.Context) {
	//tokenUserID := c.GetString("userId")

	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		c.JSON(http2.NewErrorResponse(error2.BadRequestError))
		return
	}
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		c.JSON(http2.NewErrorResponse(error2.BadRequestError))
		return
	}
	artistID, exist := c.GetQuery("artistId")
	if !exist {
		c.JSON(http2.NewErrorResponse(error2.BadRequestError))
		return
	}

	filter := model.OpenCommissionFilter{
		ArtistID: artistID,
		Count:    count,
		Offset:   offset,
	}

	result, err := o.openCommUseCase.GetOpenCommissions(c, filter)
	if err != nil {
		c.JSON(http2.NewErrorResponse(err))
	}
	c.JSON(http.StatusOK, get_open_commissions.NewResponse(*result, artistID, count))
}

func (o OpenCommissionController) UpdateOpenCommission(c *gin.Context) {
	artistID := c.Query("artistId")
	openCommissionID := c.Query("openCommissionId")
	title := c.Query("title")
	desc := c.Query("desc")
	//price := c.Query("price")

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

	updater := model.OpenCommissionUpdater{
		ID:      openCommissionID,
		Title:   &title,
		Desc:    &desc,
		DayNeed: dayNeed,
	}

	err := o.openCommUseCase.UpdateOpenCommission(c, artistID, updater)
	if err != nil {
		return
	}

	c.PureJSON(200, nil)
}

func (o OpenCommissionController) DeleteOpenCommission(c *gin.Context) {

}
