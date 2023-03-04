package dao

import (
	"Raising/model"
	"context"
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> fd910d7 (golang)

	"gorm.io/gorm"
)

type ImgDao struct {
	*gorm.DB
}

func NewImgDao(ctx context.Context) *ImgDao {
	return &ImgDao{NewDbClient(ctx)}
}

// 用于复用db
func NewImgDaoByDB(db *gorm.DB) *ImgDao {
	return &ImgDao{db}
}
func (dao *ImgDao) CreateImg(img *model.Project_Img) error {
<<<<<<< HEAD
	fmt.Println(img)
	return dao.DB.Model(&model.Project_Img{}).Create(&img).Error
}
func (dao *ImgDao) GetImgByPid(img []*model.Project_Img, pid string) error {
	fmt.Println(img)
=======
	return dao.DB.Model(&model.Project_Img{}).Create(&img).Error
}
func (dao *ImgDao) GetImgByPid(img []*model.Project_Img, pid string) error {
>>>>>>> fd910d7 (golang)
	return dao.DB.Model(&model.Project_Img{}).Where("pid=?", pid).Find(&img).Error
}
func (dao *ImgDao) DeleteImgByPid(pid string) error {

	return dao.DB.Where("pid=?", pid).Delete(&model.Project_Img{}).Error
}
