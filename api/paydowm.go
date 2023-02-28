package api

import (
	"Raising/conf"
	"Raising/service"
	"Raising/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PayDown
//
//	@Description	支付
//	@Summary		支付
//	@Accept			application/json
//	@Produce		application/json
//	@Param			order_id		formData	int		true	"订单id"
//	@Param			money			formData	int		true	"支付金额"
//	@Param			Authorization	header		string	true	"token"
//	@Success		200				{object}	vo.Response
//	@Failure		500				{object}	vo.Response
//	@Router			/user/pay [POST]
func PayDown(c *gin.Context) {
	var service service.PayService
	claims, _ := util.ParseToken(c.GetHeader(conf.Head))
	if err := c.ShouldBind(&service); err == nil {
		rsp := service.InvestToProject(c, claims.ID)
		c.JSON(http.StatusOK, rsp)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("[payfor]", err)
	}
}
