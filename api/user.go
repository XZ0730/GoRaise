package api

import (
	"Raising/conf"
	"Raising/service"
	"Raising/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Register
//
//	@Description	注册
//	@Summary		注册用户名和密码
//	@Accept			application/json
//	@Produce		application/json
//	@Param			user_name	formData	string	true	"用户名"
//	@Param			password	formData	string	true	"密码"
//	@Param			email		formData	string	true	"邮箱"
//	@Param			phone		formData	string	true	"手机号码"
//	@Success		200			{object}	vo.Response
//	@Failure		500			{object}	vo.Response
//	@Router			/register [POST]
func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	if err := c.ShouldBind(&userRegister); err == nil {
		rsp := userRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, rsp)
	} else {
		c.JSON(http.StatusBadRequest, err)
		util.LogrusObj.Info("[register]:", err)
	}
}

// Login
//
//	@Description	登录
//	@Summary		登录并返回token
//	@Accept			application/json
//	@Produce		application/json
//	@Param			user_name	formData	string	true	"用户名"
//	@Param			password	formData	string	true	"密码"
//	@Success		200			{object}	vo.Response{data=vo.TokenData{user=vo.User}}
//	@Failure		500			{object}	vo.Response
//	@Router			/login [POST]
func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err == nil {
		rsp := userLogin.Login(c.Request.Context())
		c.JSON(http.StatusOK, rsp)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("[login]:", err)
	}
}

// UpdateUser
//
//	@Description	修改用户信息
//	@Summary		修改用户信息
//	@Accept			application/json
//	@Produce		application/json
//	@Param			nick_name		formData	string	true	"昵称"
//	@Param			phone			formData	string	false	"手机号"
//	@Param			Authorization	header		string	true	"token"
//	@Success		200				{object}	vo.Response{data=vo.User}
//	@Failure		500				{object}	vo.Response
//	@Router			/user [PUT]
func UserUpdate(c *gin.Context) {
	var userUpdate service.UserService
	claims, _ := util.ParseToken(c.GetHeader(conf.Head))
	if err := c.ShouldBind(&userUpdate); err == nil {
		rsp := userUpdate.Update(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, rsp)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("[update]:", err)
	}
}

// GetUser
//
//	@Description	获取用户信息
//	@Summary		获取用户信息
//	@Accept			application/json
//	@Produce		application/json
//	@Param			id				query		string	false	"用户id"
//	@Param			Authorization	header		string	true	"token"
//	@Success		200				{object}	vo.Response{data=vo.User}
//	@Failure		500				{object}	vo.Response
//	@Router			/user [GET]
func UserGet(c *gin.Context) {
	var userGet service.UserService
	id, ok := c.GetQuery("id")
	if !ok {
		claims, _ := util.ParseToken(c.GetHeader(conf.Head))
		rsp := userGet.Getuser(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, rsp)
	} else {
		id, err := strconv.Atoi(id)
		if err != nil {
			util.LogrusObj.Info("[get]", err)
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
		rsp := userGet.Getuser(c.Request.Context(), uint(id))
		c.JSON(http.StatusOK, rsp)
	}
}

// UploadAvatar
//
//	@Description	上传用户头像
//	@Summary		上传用户头像
//	@Accept			application/json
//	@Produce		application/json
//	@Param			file			formData	file	true	"头像"
//	@Param			Authorization	header		string	true	"token"
//	@Success		200				{object}	vo.Response
//	@Failure		500				{object}	vo.Response
//	@Router			/user/upload_ava [POST]
func UploadAvatar(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	var uploadAvatar service.UserService
	claims, _ := util.ParseToken(c.GetHeader(conf.Head))
	if err := c.ShouldBind(&uploadAvatar); err == nil {
		rsp := uploadAvatar.UploadAvatar(c.Request.Context(), claims.ID, file, fileHeader, fileSize)

		c.JSON(http.StatusOK, rsp)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("[uploadAva]", err)
	}
}

// SendEmail
//
//	@Description	发送验证邮件
//	@Summary		发送验证邮件
//	@Accept			application/json
//	@Produce		application/json
//	@Param			operation_type	formData	int		true	"操作类型/1-绑定邮箱/2-解绑邮箱/3-修改密码"
//	@Param			email			formData	string	true	"目标邮箱"
//	@Param			passwrod		formData	string	true	"修改后的密码"
//	@Param			Authorization	header		string	true	"token"
//	@Success		200				{object}	vo.Response
//	@Failure		500				{object}	vo.Response
//	@Router			/user/email [POST]
func SendEmail(c *gin.Context) {
	var sendeamil service.SendEmailService
	claims, _ := util.ParseToken(c.GetHeader(conf.Head))
	if err := c.ShouldBind(&sendeamil); err == nil {

		rsp := sendeamil.SendEmail(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, rsp)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Info("[send email]", err)
	}
}

// ValidEmail
//
//	@Description	验证邮件
//	@Summary		验证邮件
//	@Accept			application/json
//	@Produce		application/json
//	@Param			token			path		string	true	"验证token"
//	@Param			Authorization	header		string	true	"token"
//	@Success		200				{object}	vo.Response
//	@Failure		500				{object}	vo.Response
//	@Router			/user/email/:token [GET]
func ValidEmail(c *gin.Context) {
	var valideamil service.ValidEmailService
	token := c.Param("token")
	rsp := valideamil.Valid(c.Request.Context(), token)
	c.JSON(http.StatusOK, rsp)

}

func IsAdmin(c *gin.Context) {
	c2, err := util.ParseToken(c.GetHeader(conf.Head))
	if c2.Authority == 0 || err != nil {
		c.Abort()
	} else {
		c.Next()
	}
}

// GetUserAudit
//
//	@Description	获取用户正在审核的项目
//	@Summary		获取用户正在审核的项目
//	@Accept			application/json
//	@Produce		application/json
//	@Param			Authorization	header		string	true	"token"
//	@Success		200				{object}	vo.Response{data=vo.DataList{item=vo.Audit}}
//	@Failure		500				{object}	vo.Response
//	@Router			/user/audit [GET]
func GetUserAudit(c *gin.Context) {
	var service service.UserService
	claims, _ := util.ParseToken(c.GetHeader(conf.Head))
	rsp := service.GetAuditByUid(c, claims.ID)
	c.JSON(http.StatusOK, rsp)
}
