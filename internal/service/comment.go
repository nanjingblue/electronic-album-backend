package service

import (
	"electronic-album/internal/dao"
	"electronic-album/internal/model"
	"electronic-album/internal/serializer"
	"electronic-album/pkg/convert"
)

type CommentService struct{}

type CommentGetListService struct {
	CommentService
	PostID uint `form:"post_id" json:"post_id" binding:"required"`
}

func (c *CommentGetListService) GetList(svc *Service) serializer.Response {
	// 首先判断 post_id 是不是存在
	post, err := dao.Post.GetPostByID(c.PostID)
	if err != nil {
		return serializer.Response{
			Code:  400,
			Msg:   "获取评论列表失败",
			Error: err.Error(),
		}
	}
	comments, err := dao.Comment.GetAllCommentByPostID(post.ID)
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "获取评论列表失败",
			Error: err.Error(),
		}
	}

	// 返回的信息需要包含 写评论人的信息 这部分放在 serializer
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildComments(comments),
		Msg:  "获取评论列表成功",
	}
}

type CommentCreateService struct {
	CommentService
	PostID  string `form:"post_id" json:"post_id" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}

func (c *CommentCreateService) CreateComment(svc *Service) serializer.Response {
	// 首先判断 post_id 是不是存在
	postID := convert.StrTo(c.PostID).MustUInt()
	post, err := dao.Post.GetPostByID(postID)
	if err != nil {
		return serializer.Response{
			Code:  400,
			Msg:   "获取评论列表失败",
			Error: err.Error(),
		}
	}

	u, _ := svc.ctx.Get("user")
	user := u.(model.User)
	comment := model.Comment{
		PostID:  post.ID,
		UserID:  u.(model.User).ID,
		Content: c.Content,
	}

	err = dao.Comment.CreateComment(&comment)
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "服务器内部错误",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildComment(&comment, &user),
		Msg:  "添加评论成功",
	}
}
