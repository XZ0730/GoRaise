package api

import (
	"Raising/vo"
	"encoding/json"
)

func ErrorResponse(err error) vo.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return vo.Response{
			Status: 400,
			Msg:    "JSON格式不匹配",
			Error:  err.Error(),
		}
	}
	return vo.Response{
		Status: 400,
		Msg:    "参数错误",
		Error:  err.Error(),
	}
}
