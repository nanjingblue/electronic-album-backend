package service

import (
	"electronic-album/global"
	"electronic-album/internal/model"
	"electronic-album/internal/serializer"
	"fmt"
)

type UserRegisterRequest struct {
	Username        string `form:"username" json:"username" binding:"required,min=1,max=12"`
	Password        string `form:"password" json:"password" binding:"required,min=6,max=30"`
	PasswordConfirm string `form:"confirm" json:"confirm" binding:"required,min=6,max=30"`
	Sex             string `form:"sex" json:"sex"`
	Age             uint   `form:"age" json:"age"`
}

func (rr *UserRegisterRequest) Valid() *serializer.Response {
	if rr.Password != rr.PasswordConfirm {
		return &serializer.Response{
			Code:  400,
			Msg:   "注册失败",
			Error: fmt.Errorf("两次密码不一致").Error(),
		}
	}
	var u model.User
	err := u.GetUserByUsername(rr.Username)
	if err == nil {
		return &serializer.Response{
			Code:  400,
			Msg:   "注册失败",
			Error: "该用户已存在",
		}
	}
	return nil
}

func (svc *Service) Register(param *UserRegisterRequest) serializer.Response {
	// 表单验证
	if err := param.Valid(); err != nil {
		return *err
	}

	user := model.User{
		Username: param.Username,
		Status:   model.Active,
		Sex:      param.Sex,
		Age:      param.Age,
	}

	// 加密密码
	if err := user.SetPassword(param.Password); err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "注册失败",
			Error: fmt.Errorf("加密密码失败").Error(),
		}
	}
	// 创建用户
	if err := global.DBEngine.Create(&user).Error; err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "注册失败",
			Error: fmt.Errorf("服务器内部错误").Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "注册成功",
		Data: serializer.BuildUser(user),
	}
}

type UserLoginRequest struct {
	Username string `form:"username" json:"username" binding:"required,min=1,max=12"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=30"`
}

func (svc *Service) Login(param *UserLoginRequest) serializer.Response {
	var user model.User

	if err := global.DBEngine.Where("username = ?", param.Username).First(&user).Error; err != nil {
		return serializer.Response{
			Code:  400,
			Msg:   "登录失败",
			Error: fmt.Errorf("未找到用户").Error(),
		}
	}

	if user.CheckPassword(param.Password) == false {
		return serializer.Response{
			Code:  400,
			Msg:   "登录失败",
			Error: fmt.Errorf("密码错误").Error(),
		}
	}

	// 将用户id保存在上下文
	svc.ctx.Set("user_id", user.ID)

	return serializer.Response{
		Code: 200,
		Msg:  "登录成功",
		Data: serializer.BuildUser(user),
	}
}
