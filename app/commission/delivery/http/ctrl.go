package http

import (
	"github.com/gin-gonic/gin"
	"pixstall-artist/domain/commission"
)

type CommissionController struct {
	commUseCase commission.UseCase
}

func NewCommissionController(commUseCase commission.UseCase) CommissionController {
	return CommissionController{
		commUseCase: commUseCase,
	}
}

func (cc CommissionController) AddCommission(c *gin.Context) {


}