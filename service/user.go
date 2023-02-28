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
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
	// redis "github.com/go-redis/redis/v8"
	"gopkg.in/mail.v2"
)

type UserService struct {
	UserName    string `json:"username" form:"username"`
	Passwrod    string `json:"password" form:"password"`
	PhoneNumber string `json:"phone" form:"phone"`
	NickName    string `json:"nickname" form:"nickname"`
	Invest      int64  `json:"invest" form:"invest"`
	OrderId     uint   `json:"orderid" form:"orderid"`
}
type SendEmailService struct {
	Email         string `json:"email" form:"email"`
	Passwrod      string `json:"password" form:"password"`
	OperationType uint   `json:"operationtype" form:"operationtype"`
	//1.绑定邮箱 2.解绑 3.改密码
}
type ValidEmailService struct {
}

// 注册
func (service *UserService) Register(ctx context.Context) vo.Response {
	var user *model.User
	var err error
	code := e.Success
	userdao := dao.NewUserDao(ctx)
	_, ok := userdao.ExistOrNotUsername(service.UserName)
	if ok {
		code = e.ErrorExist
		util.LogrusObj.Info("username have existed")
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user = &model.User{
		UserName:    service.UserName,
		PhoneNumber: service.PhoneNumber,
		NickName:    service.NickName,
	}
	err = user.Setpwd(service.Passwrod)
	if err != nil {
		code = e.ErrorBcrypt
		util.LogrusObj.Info("register error bcrypt")
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	err = userdao.CreateUser(user)
	if err != nil {
		code = e.ErrorCreateUser
		util.LogrusObj.Info("register error create")
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	err2 := cache.RedisClient.ZAdd(util.Key, redis.Z{Score: 0.0, Member: util.GetKey(user.ID)}).Err()
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
		Msg:    e.GetMsg(code),
	}
}

// 登录
func (service *UserService) Login(ctx context.Context) vo.Response {
	code := e.Success
	var user *model.User
	var ok bool
	//判断用户名
	userdao := dao.NewUserDao(ctx)
	user, ok = userdao.ExistOrNotUsername(service.UserName)
	if !ok {
		code = e.ErrorUserNoExist
		util.LogrusObj.Info("[login]", e.GetMsg(code))
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//判断密码
	ok = user.CheckPwd(service.Passwrod)
	if !ok {
		code = e.ErrorPasswrod
		util.LogrusObj.Info("[login]", e.GetMsg(code))
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	//生成token 返回token
	authority := 0
	if user.IsAdmin {
		authority = 1
	}

	token, err := util.GenerateToken(user.ID, service.UserName, authority)
	if err != nil {
		code = e.ErrorAuthToken
		util.LogrusObj.Info("[login]", e.GetMsg(code))
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return vo.Response{
		Status: code,
		Data: vo.TokenData{
			User:  vo.Builduser(user),
			Token: token,
		},
		Msg: e.GetMsg(code),
	}
}

// 更新用户信息
func (service *UserService) Update(ctx context.Context, uid uint) vo.Response {
	code := e.Success
	var user *model.User
	var err error
	//找到用户
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uid)
	if err != nil {
		code = e.ErrorGetUser
		util.LogrusObj.Info(e.GetMsg(code))
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	update(service, user)
	//更新用户
	err = userDao.UpdataUserById(uid, user)
	if err != nil {
		code = e.ErrorUpdateUser
		util.LogrusObj.Info(e.GetMsg(code))
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return vo.Response{
		Status: code,
		Data:   vo.Builduser(user),
		Msg:    e.GetMsg(code),
	}
}

func update(service *UserService, user *model.User) {
	if service.NickName != "" {
		user.NickName = service.NickName
	}
	// if service.PhoneNumber != "" {
	// 	user.PhoneNumber = service.PhoneNumber
	// }
}

// 获取用户信息
func (service *UserService) Getuser(ctx context.Context, uid uint) vo.Response {
	code := e.Success
	var err error
	userdao := dao.NewUserDao(ctx)
	user, err := userdao.GetUserById(uid)
	if err != nil {
		code = e.ErrorGetUser
		util.LogrusObj.Info("[get user]", err)
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return vo.Response{
		Status: code,
		Data:   vo.Builduser(user),
		Msg:    e.GetMsg(code),
	}
}

// 上传头像
func (service *UserService) UploadAvatar(ctx context.Context, uid uint, file multipart.File, fileheader *multipart.FileHeader, fileSize int64) vo.Response {
	code := e.Success
	var user *model.User
	var err error
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uid)
	if err != nil {
		code = e.Error
		util.LogrusObj.Info("[getuser]", err)
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	status, Imgurl := util.UploadToQiNiu(file, fileheader, fileSize, strconv.Itoa(int(uid)))
	user.Avatar = Imgurl
	err = userDao.UpdataUserById(uid, user)
	if err != nil {
		code = e.Error
		util.LogrusObj.Info("[update]", err)
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return vo.Response{
		Status: status,
		Data:   vo.Builduser(user),
		Msg:    e.GetMsg(code),
	}
}

// 发送邮件
func (service *SendEmailService) SendEmail(ctx context.Context, uid uint) vo.Response {
	code := e.Success

	var user *model.User
	var address string
	var mailStr string
	var targetEmail string
	var build strings.Builder

	token, err := util.GenerateEmailToken(uid, service.OperationType, service.Email, service.Passwrod)
	if err != nil {
		code = e.ErrorAuthToken
		util.LogrusObj.Info("[SE TokenAuth]", e.GetMsg(code))

		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  e.GetMsg(code),
		}
	}
	build.WriteString("\n")
	build.WriteString(conf.ValidEmail) //发送方
	build.WriteString(token)
	address = build.String()
	userdao := dao.NewUserDao(ctx)
	user, _ = userdao.GetUserById(uid)

	switch service.OperationType {
	case 1: //绑定
		mailStr = conf.ConTent1
		targetEmail = service.Email
	case 2: //解绑
		mailStr = conf.ConTent2
		targetEmail = user.Email
	case 3: //修改密码
		mailStr = conf.ConTent3
		targetEmail = user.Email
		if targetEmail == "" {
			code = e.ErrorNoExistEmail
			util.LogrusObj.Info("[targetEmail]", e.GetMsg(code))
			return vo.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  e.GetMsg(code),
			}
		}
	}

	mailText := strings.Replace(mailStr, "Email", address, -1)
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)

	m.SetHeader("To", service.Email)
	m.SetHeader("Subject", targetEmail)
	m.SetBody("text/html", mailText)
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err = d.DialAndSend(m); err != nil {
		code = e.ErrorSendEmail
		util.LogrusObj.Info("[SEmail]", e.GetMsg(code))
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return vo.Response{
		Status: code,
		Data: vo.TokenData{
			User:  user,
			Token: token,
		},
		Msg: e.GetMsg(code),
	}
}

// 邮箱验证
func (eService *ValidEmailService) Valid(ctx context.Context, token string) vo.Response {
	var (
		userID        uint
		email         string
		password      string
		operationType uint
	)
	code := e.Success
	if token == "" {
		code = e.InvalidParams
	} else {
		claims, err := util.ParseEmailToken(token)
		if err != nil {
			code = e.ErrorAuthToken
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorTokenTimeout
		} else {
			userID = claims.UserID
			email = claims.Email
			password = claims.Passwrod
			operationType = claims.OperationType
		}
	}
	if code != e.Success {
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userID)
	if err != nil {
		code = e.Error
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	switch operationType {
	case 1:
		//绑定邮箱
		user.Email = email
	case 2:
		//解绑邮箱
		user.Email = ""
	case 3:
		err = user.Setpwd(password)
		if err != nil {
			code = e.Error
			return vo.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	}
	err = userDao.UpdataUserById(userID, user)
	if err != nil {
		code = e.Error
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	fmt.Println("user:", user)
	return vo.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   vo.Builduser(user),
	}
}

// 获取用户自己的审核列表
func (service *UserService) GetAuditByUid(ctx context.Context, uid uint) vo.Response {
	code := e.Success
	auditDao := dao.NewAuditDao(ctx)
	audit, err2 := auditDao.GetAuditByUid(uid)
	if err2 != nil {
		code = e.ErrorGetAudit
		return vo.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err2.Error(),
		}
	}
	return vo.Response{
		Status: code,
		Data:   vo.BuildListResponse(audit, uint(len(audit))),
	}
}
