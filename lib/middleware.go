package lib

import (
	"github.com/gin-gonic/gin"
)

func CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Header.Get("token") == "" {
			ctx.AbortWithStatusJSON(403, gin.H{"message": "token required"})
		} else {
			ctx.Set("user_name", ctx.Request.Header.Get("token"))
			ctx.Next()
		}
	}
}

func RBAC() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, _ := ctx.Get("user_name")
		if ok, err := E.Enforce(user, ctx.Request.RequestURI, ctx.Request.Method); err != nil || !ok {
			ctx.AbortWithStatusJSON(403, gin.H{"message": "forbidden"})
		} else {
			ctx.Next()
		}
	}
}

func Middlewares() (fs []gin.HandlerFunc) {
	fs = append(fs, CheckLogin(), RBAC())
	return
}
