package vo

import (
	"time"
)

type Order struct {
	Id        uint       `json:"orderid"`
	Create_At *time.Time `json:"createat"`
	Uid       uint       `json:"uid"`
	Pid       string     `json:"pid"`
	Money     float64    `json:"money"`
}
