package service

import (
	"electronic-album/internal/dao"
	"electronic-album/internal/model"
	"electronic-album/internal/serializer"
)

type PostService struct{}

type PostGetListService struct {
	PostService
}

func (p PostGetListService) GetList(svc *Service) serializer.Response {
	u, _ := svc.ctx.Get("user")
	user := u.(model.User)

	posts, err := dao.Post.GetPosts(user.ID)
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "服务器内部错误",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildPosts(posts),
		Msg:  "获取所有follow post 成功",
	}
}

type PostGetMyListService struct{}

func (p *PostGetMyListService) GetList(svc *Service) serializer.Response {
	u, _ := svc.ctx.Get("user")
	user := u.(model.User)

	posts, err := dao.Post.GetAllPostByUserID(user.ID)
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "服务器内部错误",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildPosts(posts),
		Msg:  "获取所有my post 成功",
	}
}

type PostCreateService struct {
	PostService
	Content string `form:"content" json:"content" binding:"required"`
	Path    string `form:"path" json:"path"`
}

func (p *PostCreateService) CreatePost(svc *Service) serializer.Response {
	u, _ := svc.ctx.Get("user")
	user := u.(model.User)

	post := model.Post{
		Content: p.Content,
		UserID:  user.ID,
		Image:   p.Path,
	}

	err := dao.Post.CreatePost(&post)
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "服务器内部错误",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildPost(&post, &user),
		Msg:  "创建post成功",
	}
}
