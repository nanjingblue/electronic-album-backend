package serializer

import (
	"electronic-album/internal/dao"
	"electronic-album/internal/model"
)

type Post struct {
	ID             uint   `json:"id"`
	PostUserID     uint   `json:"post_user_id"`
	PostUsername   string `json:"post_username"`
	PostNickname   string `json:"post_nickname"`
	PostUserAvatar string `json:"post_user_avatar"`
	Content        string `json:"content"`
	Image          string `json:"image"`
	PostTime       string `json:"post_time"`
}

func BuildPost(p *model.Post, u *model.User) Post {
	return Post{
		ID:             p.ID,
		PostUserID:     u.ID,
		PostUsername:   u.Username,
		PostNickname:   u.Nickname,
		PostUserAvatar: u.Avatar,
		Content:        p.Content,
		Image:          p.GetURl(),
		PostTime:       p.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func BuildPosts(it []model.Post) []Post {
	var posts []Post
	for _, item := range it {
		u, _ := dao.User.GetUserByID(item.UserID)
		posts = append(posts, BuildPost(&item, &u))
	}
	return posts
}
