package vo

import (
	"Raising/api/cache"
	"Raising/model"
	"Raising/util"
)

type User struct {
	ID       uint    `json:"id"`
	UserName string  `json:"username"`
	NickName string  `json:"nickname"`
	Email    string  `json:"email"`
	Avatar   string  `json:"avatar"`
	CreateAt int64   `json:"createat"`
	IsAdmin  bool    `json:"isadmin"`
	Score    float64 `json:"score"`
}

func Builduser(user *model.User) *User {
	score, _ := cache.RedisClient.ZScore(util.Key, util.GetKey(user.ID)).Result()
	return &User{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Avatar:   user.Avatar,
		CreateAt: user.CreatedAt.Unix(),
		IsAdmin:  user.IsAdmin,
		Score:    score,
	}
}
