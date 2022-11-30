package model

import (
	"electronic-gallery/global"
	"electronic-gallery/internal/cache"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gorm.io/gorm"
	"strconv"
)

/*
Post 朋友圈发贴子
这里为了偷懒 规定贴子最多只能1张照片
*/
type Post struct {
	gorm.Model
	UserID  uint `gorm:"not null"`
	User    User
	Content string
	Image   string
}

func (p *Post) GetURl() string {
	client, _ := oss.New(global.OSSSetting.END_POINT, global.OSSSetting.ACCESS_KEY_ID, global.OSSSetting.ACCESS_KEY_SECRET)
	bucket, _ := client.Bucket(global.OSSSetting.BUCKET)
	signedGetURL, _ := bucket.SignURL(p.Image, oss.HTTPGet, 600)
	return signedGetURL
}

// View 点击数
func (p *Post) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.PostViewKey(p.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView post点击
func (p *Post) AddView() {
	cache.RedisClient.Incr(cache.PostViewKey(p.ID))
}

// Like post喜欢数
func (p *Post) Like() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.PostLikeKey(p.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddLike post喜欢+1
func (p *Post) AddLike() {
	cache.RedisClient.Incr(cache.PostLikeKey(p.ID))
}

func (p *Post) CancelLike() {
	cache.RedisClient.Decr(cache.PostLikeKey(p.ID))
}

// Collection 收藏数
func (p *Post) Collection() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.PostCollectionKey(p.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddCollection 收藏添加
func (p *Post) AddCollection() {
	cache.RedisClient.Incr(cache.PostCollectionKey(p.ID))
}

// CancelCollection 取消收藏
func (p *Post) CancelCollection() {
	cache.RedisClient.Decr(cache.PostCollectionKey(p.ID))
}

func (p *Post) Comment() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.PostCommentKey(p.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (p *Post) AddComment() {
	cache.RedisClient.Incr(cache.PostCommentKey(p.ID))
}

/*
PostSlice 切片排序
*/
type PostSlice []Post

func (p PostSlice) Len() int {
	return len(p)
}

func (p PostSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PostSlice) Less(i, j int) bool {
	return p[i].CreatedAt.Before(p[j].CreatedAt)
}
