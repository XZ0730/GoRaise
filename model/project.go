package model

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	PName      string `json:"p_name" form:"p_name"`
	Info       string `json:"info" form:"info"`
	Pid        string `gorm:"not null"`
	Uid        uint   `gorm:"not null"`
	UserScore  float64
	Accumulate float64 `gorm:"default:0"`
}
