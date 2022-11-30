package v1

import (
	"electronic-gallery/internal/model"
	"electronic-gallery/internal/serializer"
	"electronic-gallery/internal/service"
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

func UserMe(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(400, gin.H{"code": "401", "msg": "权限不足"})
	}
	ctx.JSON(200, serializer.Response{
		Code: 200,
		Data: serializer.BuildUser(user.(model.User)),
		Msg:  "查看Userinfo成功",
	})
}

// UserLogout 用户登出
func UserLogout(ctx *gin.Context) {
	ctx.Set("user", nil)
	ctx.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}

func UserUpdate(ctx *gin.Context) {
	serv := service.UserUpdateProfileService{}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&serv); err == nil {
		res := serv.Update(&svc)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "更新失败",
			"error": err.Error(),
		})
	}
}
