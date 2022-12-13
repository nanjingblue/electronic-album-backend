package v1

import (
	"electronic-gallery/internal/service"
	"electronic-gallery/pkg/convert"
	"github.com/gin-gonic/gin"
)

// @Summary 获取评论列表
// @Produce json
// @Param post_id body string true "贴子id"
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
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

// @Summary 评论发表
// @Produce json
// @Param post_id body string true "贴子id"
// @Param content body string true "内容"
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
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
