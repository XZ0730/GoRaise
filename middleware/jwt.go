package middleware

import (
	"Raising/pkg/e"
	"Raising/util"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		code = 200
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthToken

			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorTokenTimeout
			}
		}
		if code != e.Success {
			util.LogrusObj.Info(e.GetMsg(code))
			ctx.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
