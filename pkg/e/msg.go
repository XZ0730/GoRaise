package e

var MsgFlags = map[int]string{
	Success:            "ok",
	Error:              "fail",
	InvalidParams:      "参数错误",
	ErrorFailtoGET:     "错误信息获取失败",
	ErrorExist:         "用户名已存在",
	ErrorUserNoExist:   "用户名不存在",
	ErrorPhoneExist:    "手机号已存在",
	ErrorPasswrod:      "密码错误",
	ErrorGetUser:       "获取用户信息失败",
	ErrorUpdateUser:    "修改用户信息失败",
	ErrorBcrypt:        "加密失败",
	ErrorCreateUser:    "创建用户失败",
	ErrorNotCompare:    "密码错误",
	ErrorAuthToken:     "token颁发错误",
	ErrorUploadFile:    "上传加载错误",
	ErrorSendEmail:     "邮件发送失败",
	ErrorNoExistEmail:  "未绑定邮箱",
	ErrorCreateProject: "创建project失败",
	ErrorCreateAudit:   "创建audit失败",
	ErrorCreateimg:     "创建img失败",
	ErrorGetAudit:      "获取审核项目失败",
	ErrorINDeleteAudit: "审核失败",

	AuditPass:   "审核通过",
	AuditNoPass: "审核不通过",

	ErrorRedis: "redis获取数据错误",
}

//get 获取状态码对应的信息

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[ErrorFailtoGET]
	}
	return msg
}
