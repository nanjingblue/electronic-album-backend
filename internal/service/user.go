package service

import (
	"electronic-gallery/internal/dao"
	"electronic-gallery/internal/model"
	"electronic-gallery/internal/serializer"
	"electronic-gallery/pkg/app"
	"fmt"
)

// UserRegisterRequest 注册表单结构体
type UserRegisterRequest struct {
	Username string `form:"username" json:"username" binding:"required,min=1,max=12"`
	Nickname string `form:"nickname" json:"nickname" binding:"required,min=1,max=12"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=30"`
}

// Valid 表单验证
func (rr *UserRegisterRequest) Valid() *serializer.Response {
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
		Nickname: param.Nickname,
		Status:   model.Active,
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
	if err := dao.User.CreateUser(&user); err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "注册失败",
			Error: fmt.Errorf("服务器内部错误").Error(),
		}
	}

	// 发放token
	token, err := app.ReleaseToken(user)
	if err != nil {
		return serializer.Response{
			Code:  500,
			Data:  nil,
			Msg:   "登录失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code:  200,
		Msg:   "注册成功",
		Data:  serializer.BuildUser(user),
		Token: token,
	}
}

// UserLoginRequest 登录
type UserLoginRequest struct {
	Username string `form:"username" json:"username" binding:"required,min=1,max=12"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=30"`
}

func (svc *Service) Login(param *UserLoginRequest) serializer.Response {
	var user model.User

	user, err := dao.User.GetUserByUsername(param.Username)
	if err != nil {
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

	// 发放token
	token, err := app.ReleaseToken(user)
	if err != nil {
		return serializer.Response{
			Code:  500,
			Data:  nil,
			Msg:   "生成token失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code:  200,
		Msg:   "登录成功",
		Data:  serializer.BuildUser(user),
		Token: token,
	}
}

type UserUpdateProfileService struct {
	Nickname    string `form:"nickname" json:"nickname"`
	Password    string `form:"password" json:"password"`
	Age         uint   `form:"age" json:"age"`
	Avatar      string `form:"avatar" json:"avatar"`
	Description string `form:"description" json:"description"`
}

func (uup *UserUpdateProfileService) Update(svc *Service) serializer.Response {
	// 只有当前用户才能修改
	u, _ := svc.ctx.Get("user")
	user := u.(model.User)

	if uup.Nickname != "" {
		user.Nickname = uup.Nickname
	}
	if uup.Description != "" {
		user.Description = uup.Description
	}
	if uup.Password != "" {
		user.SetPassword(uup.Password)
	}
	if uup.Avatar != "" {
		user.Avatar = uup.Avatar
	}
	if 0 != uup.Age {
		uup.Age = uup.Age
	}
	err := dao.User.UpdateUser(&user)
	if err != nil {
		return serializer.Response{
			Code:  400,
			Msg:   "更新失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{}
}
