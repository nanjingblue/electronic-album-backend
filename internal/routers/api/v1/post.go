package v1

import (
	"electronic-album/internal/service"
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
