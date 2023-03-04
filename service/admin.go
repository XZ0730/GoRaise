package service

import (
	"Raising/api/cache"
	"Raising/conf"
	"Raising/dao"
	"Raising/model"
	"Raising/pkg/e"
	"Raising/util"
	"Raising/vo"
	"context"
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> fd910d7 (golang)
)

type AuditService struct {
	//审核项目，审核通过则上传，不通过则删除数据库数据，删除七牛云数据
}

var filename1 = make([]string, 5)

// 审核项目 pid用query传进来
func (service *AuditService) Audit_Project(ctx context.Context, pid string, isPass string) vo.Response {
<<<<<<< HEAD
	fmt.Println("112312312312312312312312")
=======
>>>>>>> fd910d7 (golang)
	code = e.Success
	var project *model.Project
	audit := &model.Audit{
		Pid: pid,
	}
	var err error

	auditDao := dao.NewAuditDao(ctx)
	audit, err = auditDao.GetAuditByPid(pid)

	if err != nil || audit.Id == 0 {
		code = e.InvalidParams
		util.LogrusObj.Info("[pid err]", e.GetMsg(code))
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}

	}
	projectDao := dao.NewProjectDao(ctx)
	UserScore, err := cache.RedisClient.ZScore(util.Key, util.GetKey(audit.Uid)).Result()
	if err != nil {
		code = e.ErrorRedis
<<<<<<< HEAD
		
=======

>>>>>>> fd910d7 (golang)
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	project = &model.Project{
		Pid:       pid,
		PName:     audit.PName,
		Info:      audit.Info,
		Uid:       audit.Uid,
		UserScore: UserScore,
	}
	if isPass == "yes" { //审核通过
		err = projectDao.CreateProject(project)
		if err != nil {
			code = e.ErrorCreateProject
			util.LogrusObj.Info("[project cre]", e.GetMsg(code))
			return vo.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}

		}

		code = e.AuditPass
	} else if isPass == "no" { //审核不通过
		wg.Add(1)

		imgDao := dao.NewImgDao(ctx)
		imgs := make([]*model.Project_Img, 5, 20)
		err1 := imgDao.GetImgByPid(imgs, pid)
		go FileDelet(imgs)                 //删除七牛数据
		err2 := imgDao.DeleteImgByPid(pid) //删除数据库数据

		if err1 != nil || err2 != nil {
			code = e.ErrorINDeleteAudit
			return vo.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		code = e.AuditNoPass

		wg.Wait()
	} else {
		code = e.InvalidParams
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	err = auditDao.DeleteAudit(audit)
	if err != nil {
		code = e.ErrorINDeleteAudit
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return vo.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
func FileDelet(imgs []*model.Project_Img) {

<<<<<<< HEAD
	fmt.Println(imgs)
=======
>>>>>>> fd910d7 (golang)
	for _, img := range imgs {
		if img == nil {
			break
		}
		filename1 = append(filename1, img.Img_name)
	}
<<<<<<< HEAD
	fmt.Println("--------------------------")
	err := util.DeleteFiles(filename1)
	if err != nil {
		util.LogrusObj.Info("删除七牛云文件失败")
		fmt.Println(err)
	}
	fmt.Println("=========================")
=======
	err := util.DeleteFiles(filename1)
	if err != nil {
		util.ReLogrusObj(util.Path).Info("[qiniu error]删除云文件失败")
	}
>>>>>>> fd910d7 (golang)
	wg.Done()
}

// 获取审核项目
func (service *AuditService) Get_Audit(ctx context.Context, page int, name string) vo.Response {
	var err error
	code := e.Success
	auditDao := dao.NewAuditDao(ctx)
	auditArray := make([]*model.Audit, conf.Pagesize)
	err = auditDao.GetAllAuditByName(auditArray, page, conf.Pagesize, name)

	if err != nil {
		code = e.ErrorGetAudit
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	total := 0
	for _, v := range auditArray {
		if v != nil {
			total++
		} else {
			break
		}
	}
	auditArray = auditArray[:total]
	return vo.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   vo.BuildListResponse(auditArray, uint(total)),
	}
}

// 获取审核项目具体信息
func (service *AuditService) Get_AuditByID(ctx context.Context, pid string) vo.Response {
	auditDao := dao.NewAuditDao(ctx)
	audit, err2 := auditDao.GetAuditByPid(pid)
	if err2 != nil {
		code = e.ErrorFailtoGET
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err2.Error(),
		}
	}
	project := &model.Project{
		Pid:        pid,
		PName:      audit.PName,
		Info:       audit.Info,
		Uid:        audit.Uid,
		Accumulate: 0,
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

// 删除项目
func (service *AuditService) Delete_Project(ctx context.Context, pid string) vo.Response {
	var err error
	code := e.Success
	//删除项目数据
	projectDao := dao.NewProjectDao(ctx)
	err = projectDao.DeleteProjectByPid(pid)
	imgDao := dao.NewImgDao(ctx)
	wg.Add(1)
	//删除图片数据
	imgs := make([]*model.Project_Img, 5, 20)
	err1 := imgDao.GetImgByPid(imgs, pid)
	go FileDelet(imgs)                 //删除七牛数据
	err2 := imgDao.DeleteImgByPid(pid) //删除数据库数据

	if err1 != nil || err2 != nil {
		code = e.Error
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	wg.Wait()
	return vo.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
