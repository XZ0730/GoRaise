package model

type Audit struct {
	Id     uint
	IsPass bool
	PName  string `json:"p_name" form:"p_name"`
	Info   string `json:"info" form:"info"`
	Pid    string `gorm:"not null"`
	Uid    uint   `gorm:"not null"`
}
