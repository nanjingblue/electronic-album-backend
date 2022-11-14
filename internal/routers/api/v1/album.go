package v1

import (
	"electronic-album/internal/service"
	"github.com/gin-gonic/gin"
)

/*
AlbumCreateService 创建相册
*/
func AlbumCreateService(ctx *gin.Context) {
	param := service.AlbumCreateRequest{}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&param); err == nil {
		res := svc.AlbumCreate(&param)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "创建相册失败",
			"error": err.Error(),
		})
	}
}

/*
AlbumGetListService 获取当前用户的所有相册
*/
func AlbumGetListService(ctx *gin.Context) {

}

/*
AlbumUpdateService 更新相册服务
*/
func AlbumUpdateService(ctx *gin.Context) {

}

/*
AlbumDeleteService 删除相册服务
*/
func AlbumDeleteService(ctx *gin.Context) {

}
