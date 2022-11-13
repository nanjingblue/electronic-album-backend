package middleware

import (
	"electronic-album/internal/model"
	"electronic-album/internal/serializer"
	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, ok := ctx.Get("user_id")
		if ok {
			var user model.User
			err := user.GetUserByID(uid.(uint))
			if err != nil {
				ctx.Set("user", &user)
			}
		}
		ctx.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}

		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}
