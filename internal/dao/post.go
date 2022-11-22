package dao

import (
	"electronic-album/global"
	"electronic-album/internal/model"
	"sort"
)

type postDao struct{}

var Post *postDao

func init() {
	Post = &postDao{}
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

func (p postDao) GetPostByID(uid uint) (model.Post, error) {
	var post model.Post
	return post, global.DBEngine.First(&post, uid).Error
}

func (p postDao) CreatePost(post *model.Post) error {
	return global.DBEngine.Create(&post).Error
}

func (p postDao) Update(post *model.Post) error {
	return global.DBEngine.Update(&post).Error
}

func (p postDao) Delete(post *model.Post) error {
	return global.DBEngine.Delete(&post).Error
}
