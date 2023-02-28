package dao

import (
	"Raising/model"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type ProjectDao struct {
	*gorm.DB
}

func NewProjectDao(ctx context.Context) *ProjectDao {
	return &ProjectDao{NewDbClient(ctx)}
}

// 用于复用db
func NewProjectDaoByDB(db *gorm.DB) *ProjectDao {
	return &ProjectDao{db}
}

// 判断项目是否存在
func (dao *ProjectDao) ExistOrNotProjectname(Pname string) (project *model.Project, ok bool) {
	db := dao.DB.Model(&model.User{}).Where("p_name=?", Pname).Find(&project)
	fmt.Println(project)
	if db.RowsAffected == 0 {
		return project, false
	}
	return project, true
}

// 创建项目
func (dao *ProjectDao) CreateProject(project *model.Project) error {
	return dao.DB.Model(&model.Project{}).Create(&project).Error
}

// 根据id获取项目
func (dao *ProjectDao) GetProjectById(id uint) (project *model.Project, err error) {
	err = dao.DB.Model(&model.Project{}).Where("id=?", id).Find(&project).Error
	return
}

// 根据pid获取项目
func (dao *ProjectDao) GetProjectByPId(pid string) (project *model.Project, err error) {
	err = dao.DB.Model(&model.Project{}).Where("pid=?", pid).Find(&project).Error
	return
} //GetProjectByUid
func (dao *ProjectDao) GetProjectByUid(uid uint) (project []*model.Project, err error) {
	err = dao.DB.Model(&model.Project{}).Where("uid=?", uid).Find(&project).Error
	return
}

// 用户搜索
func (dao *ProjectDao) GetProjectByName(page int, pageSize int, name string) (project []*model.Project, err error) {
	name = "%" + name + "%"
	err = dao.DB.Model(&model.Project{}).Where("p_name LIKE ?", name).Order("user_score desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&project).Error
	return
}

// 首页推送
func (dao *ProjectDao) GetProject() (project []*model.Project, err error) {
	err = dao.DB.Model(&model.Project{}).Order("user_score desc").Find(&project).Error
	return project, err
}

// 更新项目用户积分
func (dao *ProjectDao) UpdateScoreByUid(uid uint, score float64) (err error) {
	return dao.DB.Model(&model.Project{}).Where("uid=?", uid).UpdateColumn("user_score", score).Error
}

// 更新项目信息
func (dao *ProjectDao) UpdataProjectById(pid uint, project *model.Project) error {
	return dao.DB.Model(&model.Project{}).Where("id=?", pid).Updates(&project).Error
}

// 删除项目
func (dao *ProjectDao) DeleteProjectByPid(pid string) error {
	return dao.DB.Where("pid=?", pid).Delete(&model.Project{}).Error
}
