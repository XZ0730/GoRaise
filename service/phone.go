package service

import (
	"Raising/dao"
	"Raising/pkg/e"
	"Raising/util"
	"Raising/vo"
	"context"
)

type PhoneService struct {
	Phone string `json:"phone" form:"phone"`
}

func (service *PhoneService) SendPhoneNum(ctx context.Context) vo.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	if service.Phone == "" {
		code = e.Error
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	_, ok := userDao.ExistOrNotPhoneNum(service.Phone)
	if ok {
		code = e.ErrorPhoneExist
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	so := util.NewSms()
	err := so.SendVerificationCode(service.Phone)
	if err != nil {
		code = e.Error
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
