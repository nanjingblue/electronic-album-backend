package v1

import (
	"electronic-gallery/internal/service"
	"electronic-gallery/pkg/convert"
	"github.com/gin-gonic/gin"
)

// @Summary 贴子列表
// @Produce json
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
func PostList(ctx *gin.Context) {
	serv := service.PostGetListService{}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&serv); err == nil {
		res := serv.GetList(&svc)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "添加图片失败",
			"error": err.Error(),
		})
	}
}

// @Summary 当前用户贴子列表
// @Produce json
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
func PostMyList(ctx *gin.Context) {
	serv := service.PostGetMyListService{}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&serv); err == nil {
		res := serv.GetList(&svc)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "添加图片失败",
			"error": err.Error(),
		})
	}
}

// @Summary 发表贴子
// @Produce json
// @Param content body string true "内容"
// @Param path body string true "图片路径"
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
func PostCreate(ctx *gin.Context) {
	serv := service.PostCreateService{}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&serv); err == nil {
		res := serv.CreatePost(&svc)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "添加图片失败",
			"error": err.Error(),
		})
	}
}

// @Summary 贴子点赞
// @Produce json
// @Param post_id body string true "贴子id号"
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
func PostLike(ctx *gin.Context) {
	serv := service.PostLikeService{
		PostID: convert.StrTo(ctx.Param("post_id")).MustUInt(),
	}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&serv); err == nil {
		res := serv.Like(&svc)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "喜欢Post失败",
			"error": err.Error(),
		})
	}
}

// @Summary 贴子取消点赞
// @Produce json
// @Param post_id body string true "贴子id号"
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
func PostCancelLike(ctx *gin.Context) {
	serv := service.PostLikeService{
		PostID: convert.StrTo(ctx.Param("post_id")).MustUInt(),
	}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&serv); err == nil {
		res := serv.CancelLike(&svc)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "取消喜欢Post失败",
			"error": err.Error(),
		})
	}
}

// @Summary 贴子添加收藏
// @Produce json
// @Param post_id body string true "贴子id"
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
func PostCollection(ctx *gin.Context) {
	serv := service.PostCollectionService{
		PostID: convert.StrTo(ctx.Param("post_id")).MustUInt(),
	}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&serv); err == nil {
		res := serv.Collection(&svc)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "收藏Post失败",
			"error": err.Error(),
		})
	}
}

// @Summary 贴子取消收藏
// @Produce json
// @Param post_id body string true "贴子id"
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
func PostCancelCollection(ctx *gin.Context) {
	serv := service.PostCollectionService{
		PostID: convert.StrTo(ctx.Param("post_id")).MustUInt(),
	}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&serv); err == nil {
		res := serv.CancelCollection(&svc)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "取消收藏Post失败",
			"error": err.Error(),
		})
	}
}

// @Summary 获取所有被当前用户喜欢的贴子
// @Produce json
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
func PostListLikedByMe(ctx *gin.Context) {
	serv := service.PostListLikedByMeService{}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&serv); err == nil {
		res := serv.GetList(&svc)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "获取用户所有喜欢的 post 失败",
			"error": err.Error(),
		})
	}
}

// @Summary 获取所有被当前用户收藏的贴子
// @Produce json
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
func PostListCollectedByMe(ctx *gin.Context) {
	serv := service.PostListCollectedByMeService{}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&serv); err == nil {
		res := serv.GetList(&svc)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "获取用户所有收藏的 post 失败",
			"error": err.Error(),
		})
	}
}
