package dao

import (
	"Raising/model"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDbClient(ctx)}
}

// 用于复用db
func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

func (dao *UserDao) ExistOrNotUsername(userName string) (user *model.User, ok bool) {
	db := dao.DB.Model(&model.User{}).Where("user_name=?", userName).Find(&user)
	fmt.Println(user)
	if db.RowsAffected == 0 {
		//fmt.Println("=====================")
		return user, false
	}
	//fmt.Println("=====================")
	return user, true
}
func (dao *UserDao) ExistOrNotPhoneNum(phone string) (user *model.User, ok bool) {
	db := dao.DB.Model(&model.User{}).Where("phone_number=?", phone).Find(&user)
	fmt.Println(user)
	if db.RowsAffected == 0 {
		//	fmt.Println("=====================")
		return user, false
	}
	//fmt.Println("=====================")
	return user, true
}
func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}

func (dao *UserDao) GetUserById(uid uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", uid).Find(&user).Error
	return
}
func (dao *UserDao) UpdataUserMoney(uid uint, user *model.User) error {
	return dao.DB.Model(&model.User{}).Where("id=?", uid).UpdateColumn("money", user.Money).Error
}
func (dao *UserDao) UpdataUserById(uid uint, user *model.User) error {
	return dao.DB.Model(&model.User{}).Where("id=?", uid).Save(&user).Error
}
