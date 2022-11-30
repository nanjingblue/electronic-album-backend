package cache

import (
	"fmt"
	"strconv"
)

const (
	// DailyRankKey 每日排行
	DailyRankKey = "rank:daily"
)

// PostLikeKey 视频点击数的key
func PostLikeKey(id uint) string {
	return fmt.Sprintf("post:like:%s", strconv.Itoa(int(id)))
}

// PostViewKey 视频点击数的key
func PostViewKey(id uint) string {
	return fmt.Sprintf("post:view:%s", strconv.Itoa(int(id)))
}

// PostCollectionKey 视频点击数的key
func PostCollectionKey(id uint) string {
	return fmt.Sprintf("post:collection:%s", strconv.Itoa(int(id)))
}

// PostCommentKey 视频点击数的key
func PostCommentKey(id uint) string {
	return fmt.Sprintf("post:comment:%s", strconv.Itoa(int(id)))
}
