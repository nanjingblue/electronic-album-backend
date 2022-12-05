package serializer

import (
	"electronic-gallery/internal/dao"
	"electronic-gallery/internal/model"
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
	View           uint64 `json:"view"`
	Like           uint64 `json:"likes"`
	Comment        uint64 `json:"comments"`
	Collection     uint64 `json:"collections"`
	LikedByMe      bool   `json:"liked_by_me"`
	CollectedByMe  bool   `json:"collected_by_me"`
}

func BuildPost(p *model.Post) Post {
	return Post{
		ID:             p.ID,
		PostUserID:     p.User.ID,
		PostUsername:   p.User.Username,
		PostNickname:   p.User.Nickname,
		PostUserAvatar: p.User.AvatarURl(),
		Content:        p.Content,
		Image:          p.GetURl(),
		PostTime:       p.CreatedAt.Format("2006-01-02 15:04:05"),
		View:           p.View(),
		Like:           p.Like(),
		Collection:     p.Collection(),
		Comment:        p.Comment(),
	}
}

func BuildPosts(it []model.Post) []Post {
	var posts []Post
	for _, item := range it {
		posts = append(posts, BuildPost(&item))
	}
	return posts
}

func BuildPostsWithUser(it []model.Post, user model.User) []Post {
	var posts []Post
	for _, item := range it {
		posts = append(posts, Post{
			ID:             item.ID,
			PostUserID:     item.User.ID,
			PostUsername:   item.User.Username,
			PostNickname:   item.User.Nickname,
			PostUserAvatar: item.User.AvatarURl(),
			Content:        item.Content,
			Image:          item.GetURl(),
			PostTime:       item.CreatedAt.Format("2006-01-02 15:04:05"),
			View:           item.View(),
			Like:           item.Like(),
			Collection:     item.Collection(),
			Comment:        item.Comment(),
			LikedByMe:      dao.UserPostDAO.IsLikedByUser(user.ID, item.ID),
			CollectedByMe:  dao.UserPostDAO.IsCollectedByUser(user.ID, item.ID),
		})
	}
	return posts
}
