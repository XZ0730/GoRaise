package api

import (
	"Raising/pkg/e"
	"Raising/service"
	"Raising/util"
	"Raising/vo"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Audit_Project
//
//	@Description	审核项目
//	@Summary		审核项目
//	@Accept			application/json
//	@Produce		application/json
//	@Param			pid				query		string	true	"项目pid标识"
//	@Param			isPass			query		string	true	"是否通过"
//	@Param			Authorization	header		string	true	"token"
//	@Success		200				{object}	vo.Response
//	@Failure		500				{object}	vo.Response
//	@Router			/admin/audit [POST]
func Audit_Project(c *gin.Context) {
	var audit_service service.AuditService
	fmt.Println("1111111111111111111111111")
	pid, ok := c.GetQuery("pid")
	ispass, ok1 := c.GetQuery("isPass")

	if !ok || !ok1 {
		util.LogrusObj.Info("invalid query params")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":    e.GetMsg(e.InvalidParams),
			"status": e.InvalidParams,
		})
	}
	rsp := audit_service.Audit_Project(c, pid, ispass)
	c.JSON(http.StatusOK, rsp)

}

// Get_Audit
//
//	@Description	获取审核项目列表
//	@Summary		获取审核项目列表
//	@Accept			application/json
//	@Produce		application/json
//	@Param			name			query		string	false	"项目名称/根据名称搜索"
//	@Param			page			query		string	true	"是否通过/yes或者no"
//	@Param			Authorization	header		string	true	"token"
//	@Success		200				{object}	vo.Response{data=vo.DataList{item=vo.Audit}}
//	@Failure		500				{object}	vo.Response
//	@Router			/admin/audit [GET]
func Get_Audit(c *gin.Context) {
	var audit_service service.AuditService
	var rsp vo.Response
	var page int
	p, ok := c.GetQuery("page")
	if !ok {
		page = 1
	}
	page, _ = strconv.Atoi(p)
	name, _ := c.GetQuery("name")
	rsp = audit_service.Get_Audit(c, page, name)

	c.JSON(http.StatusOK, rsp)
}

// Get_AuditByID
//
//	@Description	获取具体某个审核项目的信息/点击某个项目获取pid然后得到具体项目信息
//	@Summary		获取具体某个审核项目的信息
//	@Accept			application/json
//	@Produce		application/json
//	@Param			pid				path		string	false	"项目pid"
//	@Param			Authorization	header		string	true	"token"
//	@Success		200				{object}	vo.Response{data=vo.Audit}
//	@Failure		500				{object}	vo.Response
//	@Router			/admin/audit/:pid [GET]
func Get_AuditByID(c *gin.Context) { //
	var audit_service service.AuditService
	pid := c.Param("pid")
	rsp := audit_service.Get_AuditByID(c, pid)
	fmt.Println("pid:", pid)
	c.JSON(http.StatusOK, rsp)

}

// Delete_Project
//
//	@Description	管理员功能:删除项目
//	@Summary		管理员功能:删除项目
//	@Accept			application/json
//	@Produce		application/json
//	@Param			pid				query		string	false	"项目pid"
//	@Param			Authorization	header		string	true	"token"
//	@Success		200				{object}	vo.Response
//	@Failure		500				{object}	vo.Response
//	@Router			/admin/audit [DELETE]
func Delete_Project(c *gin.Context) {
	var audit_service service.AuditService
	var rsp vo.Response
	pid, ok := c.GetQuery("pid")
	if ok {
		rsp = audit_service.Delete_Project(c, pid)

	}
	c.JSON(http.StatusOK, rsp)
}
