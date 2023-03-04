package dao

import (
	"Raising/model"
	"context"

	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDbClient(ctx)}
}

// 用于复用db
func NewOrderDaoByDB(db *gorm.DB) *OrderDao {
	return &OrderDao{db}
}
func (dao *OrderDao) GetOrderByUid(uid uint) (order []*model.Order, err error) {
	//获取用户订单
<<<<<<< HEAD
	err = dao.DB.Model(&model.Order{}).Where("uid=?", uid).Find(&order).Error
=======
	err = dao.DB.Model(&model.Order{}).Where("uid=? AND deleted_at IS NULL", uid).Find(&order).Error
>>>>>>> fd910d7 (golang)
	return
}
func (dao *OrderDao) GetOrderById(id uint) (order *model.Order, err error) {
	//获取用户订单
	err = dao.DB.Model(&model.Order{}).Where("id=?", id).Find(&order).Error
	return
}
func (dao *OrderDao) Create_Order(order *model.Order) error {
	//获取用户订单
	return dao.DB.Model(&model.Order{}).Create(&order).Error
}
func (dao *OrderDao) DeleteOrder(oid uint) error {
	return dao.DB.Where("id=?", oid).Delete(&model.Order{}).Error
}
