package api

import (
	"Raising/conf"
	"Raising/service"
	"Raising/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create_Order
//
//	@Description	创建订单
//	@Summary		创建订单
//	@Accept			application/json
//	@Produce		application/json
//	@Param			pid				formData	string	true	"项目pid标识"
//	@Param			money			formData	string	true	"投资金额"
//	@Param			Authorization	header		string	true	"token"
//	@Success		200				{object}	vo.Response{data=vo.Order}
//	@Failure		500				{object}	vo.Response
//	@Router			/order [POST]
func Create_Order(c *gin.Context) {
	var orderService service.OrderService
	claims, _ := util.ParseToken(c.GetHeader(conf.Head))

	if err := c.ShouldBind(&orderService); err == nil {
		rsp := orderService.Create_Order(c, claims.ID)
		c.JSON(http.StatusOK, rsp)
	} else {

		c.JSON(http.StatusBadRequest, ErrorResponse(err))
<<<<<<< HEAD
		util.LogrusObj.Info("[order create]", err)
=======
		util.ReLogrusObj(util.Path).Info("[order create]", err)
>>>>>>> fd910d7 (golang)
	}
}

// GetOderList
//
//	@Description	获取当前用户订单列表
//	@Summary		获取当前用户订单列表
//	@Accept			application/json
//	@Produce		application/json
//	@Param			Authorization	header		string	true	"token"
//	@Success		200				{object}	vo.Response{data=vo.DataList{item=vo.Order}}
//	@Failure		500				{object}	vo.Response
//	@Router			/order [GET]
func GetOderList(c *gin.Context) {
	var orderService service.OrderService
	claims, _ := util.ParseToken(c.GetHeader(conf.Head))
	if err := c.ShouldBind(&orderService); err == nil {
		rsp := orderService.GetOderList(c, claims.ID)
		c.JSON(http.StatusOK, rsp)
	} else {

		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// DeleteOrder
//
//	@Description	取消订单
//	@Summary		取消订单
//	@Accept			application/json
//	@Produce		application/json
//	@Param			order_id		formData	int		true	"订单id"
//	@Param			Authorization	header		string	true	"token"
//	@Success		200				{object}	vo.Response
//	@Failure		500				{object}	vo.Response
//	@Router			/order [DELETE]
func DeleteOrder(c *gin.Context) {
	var orderService service.OrderService
	claims, _ := util.ParseToken(c.GetHeader(conf.Head))
	if err := c.ShouldBind(&orderService); err == nil {
		rsp := orderService.DeleteOrder(c, claims.ID)
		c.JSON(http.StatusOK, rsp)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
<<<<<<< HEAD
		util.LogrusObj.Info("[order delete]", err)
=======
		util.ReLogrusObj(util.Path).Warn("[order delete]", err)
>>>>>>> fd910d7 (golang)
	}
}
