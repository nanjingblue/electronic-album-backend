package cache

import (
	"fmt"
	"strconv"
)

const (
	// DailyRankKey 每日排行
	DailyRankKey = "rank:daily"
)

func PostLikeKey(id uint) string {
	return fmt.Sprintf("post:like:%s", strconv.Itoa(int(id)))
}

func PostViewKey(id uint) string {
	return fmt.Sprintf("post:view:%s", strconv.Itoa(int(id)))
}

func PostCollectionKey(id uint) string {
	return fmt.Sprintf("post:collection:%s", strconv.Itoa(int(id)))
}

func PostCommentKey(id uint) string {
	return fmt.Sprintf("post:comment:%s", strconv.Itoa(int(id)))
}
