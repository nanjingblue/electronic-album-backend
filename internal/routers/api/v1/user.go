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

// @Summary 用户注册
// @Produce json
// @Param username body string true "用户名"
// @Param nickname body string true "昵称"
// @Param password body string true "密码"
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
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

// @Summary 用户登录
// @Produce json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
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

// @Summary 用户获取信息
// @Produce json
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
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

// @Summary 用户注销
// @Produce json
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
func UserLogout(ctx *gin.Context) {
	ctx.Set("user", nil)
	ctx.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}

// @Summary 用户更新信息
// @Produce json
// @Param nickname body string false "昵称"
// @Param password body string false "密码"
// @Param age body string false "年龄"
// @Param avatar body string false "头像"
// @Param description body string false "描述"
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
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
