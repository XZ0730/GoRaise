package service

import (
	"Raising/dao"
	"Raising/model"
	"Raising/pkg/e"
	"Raising/vo"
	"context"
)

type OrderService struct {
	Id    uint    `json:"orderid" form:"orderid"`
	Uid   uint    `json:"uid" form:"uid"`
	Pid   string  `json:"pid" form:"pid"`
	Money float64 `json:"money" form:"money"`
}

// 创建订单
func (service *OrderService) Create_Order(ctx context.Context, uid uint) vo.Response {
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	Order := &model.Order{
		Uid:   uid,
		Pid:   service.Pid,
		Money: service.Money,
	}
	err := orderDao.Create_Order(Order)
	if err != nil {
		code = e.ErrorOrder
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return vo.Response{
		Status: code,
		Data:   Order,
		Msg:    e.GetMsg(code),
	}
}

// 获取用户订单
func (service *OrderService) GetOderList(ctx context.Context, uid uint) vo.Response {
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	order, err2 := orderDao.GetOrderByUid(uid)
	if err2 != nil {
		code = e.ErrorOrder
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return vo.Response{
		Status: code,
		Data:   vo.BuildListResponse(order, uint(len(order))),
		Msg:    e.GetMsg(code),
	}
}

// 取消订单
func (service *OrderService) DeleteOrder(ctx context.Context, uid uint) vo.Response {
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	err2 := orderDao.DeleteOrder(service.Id)
	if err2 != nil {
		code = e.ErrorOrder
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return vo.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
