package dao

import (
	"Raising/model"
	"context"

	"gorm.io/gorm"
)

type AuditDao struct {
	*gorm.DB
}

func NewAuditDao(ctx context.Context) *AuditDao {
	return &AuditDao{NewDbClient(ctx)}
}

// 用于复用db
func NewAuditDaoByDB(db *gorm.DB) *AuditDao {
	return &AuditDao{db}
}
func (dao *AuditDao) CreateAudit(audit *model.Audit) error {
	return dao.DB.Model(&model.Audit{}).Create(&audit).Error
}
func (dao *AuditDao) GetAuditByPid(pid string) (audit *model.Audit, err error) {
	err = dao.DB.Model(&model.Audit{}).Where("pid=?", pid).Find(&audit).Error
	return
}
func (dao *AuditDao) DeleteAudit(audit *model.Audit) error {
	return dao.DB.Delete(&audit).Error
}
func (dao *AuditDao) GetAllAuditByPage(audit []*model.Audit, page int, pageSize int) error {
	return dao.DB.Model(&model.Audit{}).Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&audit).Error
}
func (dao *AuditDao) GetAllAuditByName(audit []*model.Audit, page int, pageSize int, name string) error {
	name = "%" + name + "%"
	return dao.DB.Model(&model.Audit{}).Where("p_name LIKE ?", name).Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&audit).Error
}
func (dao *AuditDao) GetAuditByUid(uid uint) (audit []*model.Audit, err error) {
	err = dao.DB.Model(&model.Audit{}).Where("uid=?", uid).Find(&audit).Error
	return
}
