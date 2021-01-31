package http

import (
	"github.com/gin-gonic/gin"
	domainOpenComm "pixstall-artist/domain/open-commission"
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

}