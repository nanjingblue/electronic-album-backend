package serializer

import (
	"electronic-gallery/internal/dao"
	"electronic-gallery/internal/model"
)

type Comment struct {
	ID                  uint   `json:"id"`
	PostID              uint   `json:"post_id"`
	Content             string `json:"content"`
	UserID              uint   `json:"user_id"`
	CommentTime         string `json:"comment_time"`
	CommentUserNickname string `json:"comment_user_nickname"`
	CommentUserAvatar   string `json:"comment_user_avatar"`
}

func BuildComment(it *model.Comment, u *model.User) Comment {
	return Comment{
		ID:                  it.ID,
		PostID:              it.PostID,
		Content:             it.Content,
		UserID:              it.UserID,
		CommentTime:         it.CreatedAt.Format("2006-01-02 15:04:05"),
		CommentUserNickname: u.Nickname,
		CommentUserAvatar:   u.AvatarURl(),
	}
}

func BuildComments(it []model.Comment) []Comment {
	var comments []Comment
	for _, item := range it {
		user, _ := dao.User.GetUserByID(item.UserID)
		comments = append(comments, BuildComment(&item, &user))
	}
	return comments
}
