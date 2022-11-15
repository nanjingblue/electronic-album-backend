package v1

import (
	"electronic-album/internal/serializer"
	"electronic-album/internal/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"data": "pong"})
}

func UserRegister(ctx *gin.Context) {
	param := service.UserRegisterRequest{}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&param); err == nil {
		res := svc.Register(&param)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "注册失败",
			"error": err.Error(),
		})
	}
}

func UserLogin(ctx *gin.Context) {
	param := service.UserLoginRequest{}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&param); err == nil {
		res := svc.Login(&param)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "登录失败",
			"error": err.Error(),
		})
	}
}

func UserInfo(ctx *gin.Context) {

}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
