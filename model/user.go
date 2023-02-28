package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName    string  `gorm:"unique"` //登录用户名
	Passwrod    string  //密码
	Email       string  //邮箱
	PhoneNumber string  //联系电话
	NickName    string  //昵称
	Status      string  //是否激活？
	Avatar      string  `gorm:"default:http://rqmfsxrro.hn-bkt.clouddn.com/youtian.jpg"` //头像
	Score       float64 //积分
	IsAdmin     bool    `gorm:"default:false"`
	Money       float64 `gorm:"default:0"`
}

const (
	PwdCost        = 12       //密码加密难度
	Active  string = "active" //激活用户
)

func (u *User) Setpwd(password string) error {
	bcry, err := bcrypt.GenerateFromPassword([]byte(password), PwdCost)
	u.Passwrod = string(bcry)
	return err
}
func (u *User) CheckPwd(passwrod string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Passwrod), []byte(passwrod))
	return err == nil
}
