package model

import "github.com/jinzhu/gorm"

type Project_Img struct {
	*gorm.Model
	Img_url  string
	Img_name string
	Pid      string
}
