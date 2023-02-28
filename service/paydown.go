package service

import (
	"Raising/api/cache"
	"Raising/dao"
	"Raising/pkg/e"
	"Raising/util"
	"Raising/vo"
	"context"
	"errors"
	"fmt"
)

type PayService struct {
	OrderId uint    `json:"orderid" form:"orderid"`
	Money   float64 `json:"money" form:"money"`
}

// 支付
func (service *PayService) InvestToProject(ctx context.Context, uid uint) vo.Response {
	code := e.Success
	//启动事务
	userDao := dao.NewUserDao(ctx)
	orderDao := dao.NewOrderDao(ctx)
	tx := userDao.Begin()
	//得到用户信息
	user, err1 := userDao.GetUserById(uid)
	order, err := orderDao.GetOrderById(service.OrderId)
	if order.ID == 0 {
		code = e.ErrorOrder
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("order no exists").Error(),
		}
	}
	if err1 != nil || err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if user.Money-order.Money < 0.0 {
		tx.Rollback()
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("金额不足").Error(),
		}
	}
	//用户扣钱
	user.Money -= order.Money
	fmt.Println("user.Money:", user.Money)
	err2 := userDao.UpdataUserMoney(user.ID, user)
	if err2 != nil {
		tx.Rollback()
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//项目筹资加钱
	projectDao := dao.NewProjectDao(ctx)
	project, err3 := projectDao.GetProjectByPId(order.Pid)
	if err3 != nil {
		tx.Rollback()
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	project.Accumulate += order.Money
	fmt.Println("project.Accumulate:", project.Accumulate)
	err4 := projectDao.UpdataProjectById(project.ID, project)
	if err4 != nil {
		tx.Rollback()
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//用户积分增加
	err5 := cache.RedisClient.ZIncrBy("ScoreRank", order.Money*3, util.GetKey(uid)).Err()
	if err5 != nil {
		tx.Rollback()
		code = e.Error
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	score, err6 := cache.RedisClient.ZScore(util.Key, util.GetKey(uid)).Result()
	err7 := projectDao.UpdateScoreByUid(order.Uid, score)
	if err6 != nil || err7 != nil {
		tx.Rollback()
		code = e.Error
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	fmt.Println("score:", score)
	user.Score = score
	err9 := userDao.UpdataUserById(uid, user)
	if err9 != nil {
		code = e.Error
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	fmt.Println(user)
	//结束事务-取消订单
	err8 := orderDao.DeleteOrder(service.OrderId)
	if err8 != nil {
		code = e.Error
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//提交
	tx.Commit()
	fmt.Println("uid:", uid)
	return vo.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
