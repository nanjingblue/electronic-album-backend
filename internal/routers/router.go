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
	//r.Use(middleware.Session("hello-world"))
	//r.Use(middleware.CurrentUser())
	r.Use(middleware.Cors())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/ping", v1.Ping)
		apiv1.POST("/register", v1.UserRegister)
		apiv1.POST("/login", v1.UserLogin)
		apiv1.POST("upload/token", v1.UploadToken)

		auth := apiv1.Group("")
		auth.Use(middleware.AuthMiddleware()) // 使用中间件 必须是登录状态才能使用以下接口
		{
			auth.GET("/user/me", v1.UserMe) // 获取用户详情
			auth.GET("/user/logout", v1.UserLogout)
			auth.POST("/gallery", v1.GalleryCreateService)  // 创建相册
			auth.GET("/gallerys", v1.GalleryGetListService) // 获取相册列表
		}
	}
	return r
}
