package routers

import (
	"electronic-album/internal/middleware"
	v1 "electronic-album/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Session("hello-world"))
	r.Use(middleware.CurrentUser())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/ping", v1.Ping)
		apiv1.POST("/register", v1.UserRegister)
		apiv1.POST("/login", v1.UserLogin)

		auth := apiv1.Group("")
		auth.Use(middleware.AuthRequired()) // 使用中间件 必须是登录状态才能使用以下接口
		{
			auth.GET("/user/:username", v1.UserInfo)             // 获取用户详情
			auth.POST("/album", v1.AlbumCreateService)           // 创建相册
			auth.GET("/albums/:user_id", v1.AlbumGetListService) // 获取相册列表
		}
	}
	return r
}
