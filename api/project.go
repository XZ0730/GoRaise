package api

import (
	"Raising/conf"
	"Raising/service"
	"Raising/util"
	"Raising/vo"
	"fmt"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Create_Project
//
//	@Description	创建项目
//	@Summary		创建项目
//	@Accept			application/json
//	@Produce		application/json
//	@Param			file			formData	file	true	"图片(3张)"
//	@Param			p_name			formData	string	true	"项目名称"
//	@Param			info			formData	string	true	"项目描述"
//	@Param			Authorization	header		string	true	"token"
//	@Success		200				{object}	vo.Response
//	@Failure		500				{object}	vo.Response
//	@Router			/project [POST]
func Create_Project(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	var project_cre service.ProjectService
	claims, _ := util.ParseToken(c.GetHeader(conf.Head))
	if err := c.ShouldBind(&project_cre); err == nil {
		fmt.Println("claims:", claims)
		rsp := project_cre.Create_project(c.Request.Context(), claims.ID, files)
		c.JSON(http.StatusOK, rsp)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
<<<<<<< HEAD
		util.LogrusObj.Info("[uploadProject]", err)
=======
		util.ReLogrusObj(util.Path).Warn("[uploadProject]", err)
>>>>>>> fd910d7 (golang)
	}
}

// GetProject_Pid
//
//	@Description	根据pid获取具体项目信息
//	@Summary		根据pid获取具体项目信息
//	@Accept			application/json
//	@Produce		application/json
//	@Param			pid				path		string	true	"项目pid"
//	@Param			Authorization	header		string	true	"token"
//	@Success		200				{object}	vo.Response{data=vo.Project{img_url=vo.DataList{item=vo.Project_Img}}}
//	@Failure		500				{object}	vo.Response
//	@Router			/project/:pid [GET]
func GetProject_Pid(c *gin.Context) {
	var service service.ProjectService

	pid := c.Param("pid")
	rsp := service.GetProjectByPId(c, pid)
	c.JSON(http.StatusOK, rsp)

}

// GetProject
//
//	@Description	获取首页项目推送列表或某个用户的项目列表
//	@Summary		获取首页项目推送列表或某个用户的项目列表
//	@Accept			application/json
//	@Produce		application/json
//	@Param			uid				query		int		false	"用户uid"
//	@Param			Authorization	header		string	true	"token"
//	@Success		200				{object}	vo.Response{data=vo.DataList{item=vo.Project}}
//	@Failure		500				{object}	vo.Response
//	@Router			/project [GET]
func GetProject(c *gin.Context) {
	var service service.ProjectService
	var rsp vo.Response
	uid, ok := c.GetQuery("uid")
	if !ok {
		rsp = service.GetProjectList(c)
	} else {
		id, err := strconv.Atoi(uid)
		if err != nil {
<<<<<<< HEAD
=======
			util.ReLogrusObj(util.Path).Warn("[uid error]", err)
>>>>>>> fd910d7 (golang)
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
		rsp = service.GetProjectByUid(c, uint(id))
	}

	c.JSON(http.StatusOK, rsp)
}

// SearchProject
//
//	@Description	根据项目名称搜索
//	@Summary		根据项目名称搜索
//	@Accept			application/json
//	@Produce		application/json
//	@Param			page			query		int		false	"页码"
//	@Param			name			query		string	true	"项目名称"
//	@Param			Authorization	header		string	true	"token"
//	@Success		200				{object}	vo.Response{data=vo.DataList{item=vo.Project}}
//	@Failure		500				{object}	vo.Response
//	@Router			/project/search [GET]
func SearchProject(c *gin.Context) {
	var service service.ProjectService
	var page int
	p, ok := c.GetQuery("page")
	if !ok {
		page = 1
	}
	page, _ = strconv.Atoi(p)
	name, _ := c.GetQuery("name")
	rsp := service.GetProjectByName(c, page, name)
	c.JSON(http.StatusOK, rsp)
}
