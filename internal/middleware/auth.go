package middleware

import (
	"electronic-gallery/internal/dao"
	"electronic-gallery/pkg/app"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 authorization header（token）
		tokenString := ctx.GetHeader("Authorization")
		log.Println(tokenString)
		// validate token for mate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") { // 固定写法
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort() // 将本次请求抛弃掉
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := app.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		// 验证通过后 获取 claims 中的 userId
		userId := claims.UserId
		user, _ := dao.User.GetUserByID(userId)
		// 用户不存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		// 用户存在 将 user 信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
