package serializer

import (
	"electronic-album/internal/dao"
	"electronic-album/internal/model"
)

type Post struct {
	ID           uint   `json:"id"`
	PostUserID   uint   `json:"post_user_id"`
	PostUsername string `json:"post_username"`
	PostNickname string `json:"post_nickname"`
	Content      string `json:"content"`
	Image        string `json:"image"`
	PostTime     int64  `json:"post_time"`
}

func BuildPost(p *model.Post, u *model.User) Post {
	return Post{
		ID:           p.ID,
		PostUserID:   u.ID,
		PostUsername: u.Username,
		PostNickname: u.Nickname,
		Content:      p.Content,
		Image:        p.GetURl(),
		PostTime:     p.CreatedAt.Unix(),
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
