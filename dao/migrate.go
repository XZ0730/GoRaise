package dao

import (
	"Raising/model"
	"fmt"
)

func Migrate() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.User{},
		&model.Project{},
		&model.Audit{},
		&model.Project_Img{},
		&model.Order{},
	)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}
