package dao

import (
	"electronic-gallery/global"
	"electronic-gallery/internal/model"
	"sort"
)

type postDao struct{}

var Post *postDao

func init() {
	Post = &postDao{}
}

func (p postDao) GetAllPost() ([]model.Post, error) {
	var posts []model.Post
	err := global.DBEngine.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

// GetPosts 获取好友的所有 post 并按照时间排序 最近的在前面
func (p postDao) GetPosts(uid uint) ([]model.Post, error) {
	friends, err := Friend.GetALLFollowingFriendsByUserID(uid)
	if err != nil {
		return nil, err
	}
	var posts []model.Post
	for _, f := range friends {
		friendPosts, _ := p.GetAllPostByUserID(f.ID)
		if len(friendPosts) != 0 {
			posts = append(posts, friendPosts...)
		}
	}

	myPosts, _ := p.GetAllPostByUserID(uid)
	if len(myPosts) > 0 {
		posts = append(posts, myPosts...)
	}

	sort.Sort(model.PostSlice(posts))
	return posts, nil
}

func (p postDao) GetAllPostByUserID(uid uint) ([]model.Post, error) {
	var posts []model.Post
	err := global.DBEngine.Where("user_id = ?", uid).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p postDao) GetPostsLikedByUser(uid uint) ([]model.Post, error) {
	var posts []model.Post
	err := global.DBEngine.Distinct("id").Table("posts").Select("*").Joins("inner join user_posts on posts.id = user_posts.post_id and user_posts.user_id = ? and user_posts.liked = ?", uid, true).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p postDao) GetPostsCollectedByUser(uid uint) ([]model.Post, error) {
	var posts []model.Post
	err := global.DBEngine.Distinct("id").Table("posts").Select("*").Joins("inner join user_posts on posts.id = user_posts.post_id and user_posts.user_id = ? and user_posts.collected = ?", uid, true).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p postDao) GetPostByID(uid uint) (model.Post, error) {
	var post model.Post
	return post, global.DBEngine.First(&post, uid).Error
}

func (p postDao) CreatePost(post *model.Post) error {
	return global.DBEngine.Create(&post).Error
}

func (p postDao) Update(post *model.Post) error {
	return global.DBEngine.Save(&post).Error
}

func (p postDao) Delete(post *model.Post) error {
	return global.DBEngine.Delete(&post).Error
}
