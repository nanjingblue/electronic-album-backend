package v1

import (
	"electronic-album/internal/service"
	"electronic-album/pkg/convert"
	"github.com/gin-gonic/gin"
)

func CommentList(ctx *gin.Context) {
	serv := service.CommentGetListService{
		PostID: convert.StrTo(ctx.Param("post_id")).MustUInt(),
	}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&serv); err == nil {
		res := serv.GetList(&svc)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "获取评论失败",
			"error": err.Error(),
		})
	}
}

func CommentCreate(ctx *gin.Context) {
	serv := service.CommentCreateService{}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&serv); err == nil {
		res := serv.CreateComment(&svc)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "添加评论失败",
			"error": err.Error(),
		})
	}
}
