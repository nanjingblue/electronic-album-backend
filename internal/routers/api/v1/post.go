package v1

import (
	"electronic-gallery/internal/service"
	"electronic-gallery/pkg/convert"
	"github.com/gin-gonic/gin"
)

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
