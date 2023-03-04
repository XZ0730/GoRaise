package api

import (
	"Raising/service"
	"Raising/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendPhoneNum
//
//	@Description	发送短信验证码
//	@Summary		发送短信验证码
//	@Accept			application/json
//	@Produce		application/json
//	@Param			phone	formData	string	true	"手机号"
//	@Success		200				{object}	vo.Response
//	@Failure		500				{object}	vo.Response
//	@Router			/phone [POST]
func SendPhoneNum(c *gin.Context) {
	var service service.PhoneService
	if err := c.ShouldBind(&service); err == nil {
		rsp := service.SendPhoneNum(c)
<<<<<<< HEAD
		util.LogrusObj.Info("zxzxzx")
		c.JSON(http.StatusOK, rsp)
	} else {
=======
		c.JSON(http.StatusOK, rsp)
	} else {
		util.ReLogrusObj(util.Path).Warn("[Phone error]", err)
>>>>>>> fd910d7 (golang)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
	fmt.Println("1123")
}
