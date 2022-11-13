package v1

import (
	"electronic-album/internal/service"
	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"data": "pong"})
}

func Register(ctx *gin.Context) {
	param := service.RegisterRequest{}
	svc := service.New(ctx.Request.Context())
	if err := ctx.ShouldBind(&param); err == nil {
		res := svc.Register(&param)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(500, gin.H{
			"msg":   "注册失败",
			"error": err.Error(),
		})
	}
}
