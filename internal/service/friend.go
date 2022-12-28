package service

import (
	"electronic-gallery/internal/dao"
	"electronic-gallery/internal/model"
	"electronic-gallery/internal/serializer"
)

type FriendService struct {
	Username string `form:"username" json:"username" binding:"required"`
}

func (ffs *FriendService) Follow(svc *Service) serializer.Response {
	u, _ := svc.ctx.Get("user")
	user := u.(model.User)
	err := dao.Friend.FollowByUsername(user, ffs.Username)
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "关注失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: 200,
		Msg:  "关注成功",
	}
}

func (ffs *FriendService) Ban(svc *Service) serializer.Response {
	u, _ := svc.ctx.Get("user")
	user := u.(model.User)
	err := dao.Friend.BanByUsername(user, ffs.Username)
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "拉黑失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: 200,
		Msg:  "拉黑成功",
	}
}
