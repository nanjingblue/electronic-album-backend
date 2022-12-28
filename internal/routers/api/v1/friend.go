package v1

import (
	"electronic-gallery/internal/service"
	"github.com/gin-gonic/gin"
)

func FriendFollow(ctx *gin.Context) {
	serv := service.FriendService{}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&serv); err == nil {
		res := serv.Follow(&svc)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "关注失败",
			"error": err.Error(),
		})
	}
}

func FriendBan(ctx *gin.Context) {
	serv := service.FriendService{}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&serv); err == nil {
		res := serv.Ban(&svc)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "拉黑失败",
			"error": err.Error(),
		})
	}
}
