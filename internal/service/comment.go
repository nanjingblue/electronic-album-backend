package service

import (
	"electronic-gallery/internal/dao"
	"electronic-gallery/internal/model"
	"electronic-gallery/internal/serializer"
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

	// 由于前端中，评论列表属于照片详情页，所以把增加点击数放在这里
	post.AddView()

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
	PostID  uint   `form:"post_id" json:"post_id" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}

func (c *CommentCreateService) CreateComment(svc *Service) serializer.Response {
	// 首先判断 post_id 是不是存在
	//postID := convert.StrTo(c.PostID).MustUInt()
	postID := c.PostID
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

	// 评论数加一
	post.AddComment()
	_ = dao.UserPostDAO.Comment(user.ID, post.ID)

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildComment(&comment, &user),
		Msg:  "添加评论成功",
	}
}
