package dao

import (
	"electronic-album/global"
	"electronic-album/internal/model"
)

type commentDAO struct{}

var Comment *commentDAO

func init() {
	Comment = &commentDAO{}
}

// GetAllCommentByPostID 获取某条 post 的所有评论
func (c commentDAO) GetAllCommentByPostID(postID uint) ([]model.Comment, error) {
	var comments []model.Comment
	err := global.DBEngine.Where("post_id = ?", postID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (c commentDAO) CreateComment(comment *model.Comment) error {
	return global.DBEngine.Create(&comment).Error
}

func (c commentDAO) DeleteComment(comment *model.Comment) error {
	return global.DBEngine.Delete(&comment).Error
}

func (c commentDAO) UpdateComment(comment *model.Comment) error {
	return global.DBEngine.Update(&comment).Error
}
