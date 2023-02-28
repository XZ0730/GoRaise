package model

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	Uid   uint
	Pid   string
	Money float64
}
