package service

import (
	"Raising/conf"
	"Raising/dao"
	"Raising/model"
	"Raising/pkg/e"
	"Raising/util"
	"Raising/vo"
	"context"
	"mime/multipart"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ProjectService struct {
	Uid   uint   `json:"uid" form:"uid"`
	PName string `json:"pname" form:"pname" `
	Info  string `json:"info" form:"info"`
}

var wg = sync.WaitGroup{}
var code = e.Success
var err error

// 创建项目
func (service *ProjectService) Create_project(ctx context.Context, uid uint, files []*multipart.FileHeader) vo.Response {
	code = e.Success
	var Pimg *model.Project_Img

	var audit *model.Audit
	//首先创建项目

	AuditDao := dao.NewAuditDao(ctx)
	//绑定项目名称 内容
	pid := strings.Join([]string{strconv.Itoa(int(uid)), ":", strconv.Itoa(int(time.Now().Unix()))}, "")
	//for range files 上传文件
	wg.Add(len(files))

	go uploadFiles(files, ctx, pid, Pimg)
	//持久化入数据库

	// project = &model.Project{
	// 	PName: service.PName,
	// 	Info:  service.Info,
	// 	Pid:   pid,
	// 	Uid:   uid,
	// }
	// err = projectDao.CreateProject(project)
	// if err != nil {
	// 	code = e.ErrorCreateProject
	// 	util.LogrusObj.Info("[project cre]", e.GetMsg(code))
	// 	return vo.Response{
	// 		Status: code,
	// 		Msg:    e.GetMsg(code),
	// 		Error:  err.Error(),
	// 	}
	// }
	//然后存入审核
	audit = &model.Audit{
		PName:  service.PName,
		Info:   service.Info,
		Pid:    pid,
		Uid:    uid,
		IsPass: false,
	}
	err = AuditDao.CreateAudit(audit)
	if err != nil {
		code = e.ErrorCreateAudit
		util.LogrusObj.Info("[Audit cre]", e.GetMsg(code))
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//待审核，

	wg.Wait()
	//文件url存入project_url————————先传到数据库，审核不通过则批量删除

	return vo.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// 上传图片
func uploadFiles(files []*multipart.FileHeader, ctx context.Context, pid string, p_img *model.Project_Img) {
	imgDao := dao.NewImgDao(ctx)

	for _, filehead := range files {
		// util.UploadToQiNiu(,file,file)
		file, _ := filehead.Open()
		_, url := util.UploadToQiNiu(file, filehead, filehead.Size, pid)
		if url == "bad token" {
			wg.Done()
			return
		}
		p_img = &model.Project_Img{
			Pid:      pid,
			Img_url:  url,
			Img_name: filehead.Filename,
		}

		err = imgDao.CreateImg(p_img)
		if err != nil {
			code = e.ErrorCreateimg
		}
		wg.Done()
	}
}

// 获取项目具体信息
func (service *ProjectService) GetProjectByPId(ctx context.Context, pid string) vo.Response {
	code = e.Success
	projectDao := dao.NewProjectDao(ctx)
	project, err2 := projectDao.GetProjectByPId(pid)
	if err2 != nil {
		code = e.ErrorFailtoGET
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err2.Error(),
		}
	}
	img := dao.NewImgDao(ctx)
	pImg := make([]*model.Project_Img, 5)
	err3 := img.GetImgByPid(pImg, pid)
	if err3 != nil {
		code = e.Error
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err3.Error(),
		}
	}
	total := 0
	for _, v := range pImg {
		if v != nil {
			total++
		} else {
			break
		}
	}
	pImg = pImg[:total]
	return vo.Response{
		Status: code,
		Data:   vo.BuildProject(project, pImg, uint(total)),
		Msg:    e.GetMsg(code),
	}
}

// 获取项目列表
func (service *ProjectService) GetProjectList(ctx context.Context) vo.Response {
	code = e.Success
	projectDao := dao.NewProjectDao(ctx)
	project, err2 := projectDao.GetProject()
	if err2 != nil {
		code = e.Error
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return vo.Response{
		Status: code,
		Data:   vo.BuildListResponse(project, uint(len(project))),
		Msg:    e.GetMsg(code),
	}
}

// 获取用户项目列表 	可以是自己也可以是其他用户
func (service *ProjectService) GetProjectByUid(ctx context.Context, uid uint) vo.Response {
	code = e.Success
	projectDao := dao.NewProjectDao(ctx)
	project, err2 := projectDao.GetProjectByUid(uid)
	if err2 != nil {
		code = e.Error
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return vo.Response{
		Status: code,
		Data:   vo.BuildListResponse(project, uint(len(project))),
		Msg:    e.GetMsg(code),
	}
}
func (service *ProjectService) GetProjectByName(ctx context.Context, page int, p_name string) vo.Response {
	code := e.Success
	projectDao := dao.NewProjectDao(ctx)
	project, err2 := projectDao.GetProjectByName(page, conf.Pagesize, p_name)
	if err2 != nil {
		code = e.Error
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return vo.Response{
		Status: code,
		Data:   vo.BuildListResponse(project, uint(len(project))),
		Msg:    e.GetMsg(code),
	}
}
