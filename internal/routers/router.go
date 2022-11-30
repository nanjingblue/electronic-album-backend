package routers

import (
	"electronic-gallery/internal/middleware"
	v1 "electronic-gallery/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
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
			auth.GET("/user/me", v1.UserMe)          // 获取用户详情
			auth.POST("/user/update", v1.UserUpdate) // 更新用户资料
			auth.GET("/user/logout", v1.UserLogout)

			auth.POST("/gallery", v1.GalleryCreateService)  // 创建相册
			auth.GET("/gallerys", v1.GalleryGetListService) // 获取相册列表
			auth.PUT("gallery", v1.GalleryUpdateService)
			auth.DELETE("/gallery", v1.GalleryDeleteService)

			auth.POST("/picture", v1.PictureCreate)
			auth.GET("/pictures", v1.PictureGetPicture) // 获取某个相册的所有照片

			auth.GET("/posts", v1.PostList)                              // 获取所有关注者的 post
			auth.GET("/posts/me", v1.PostMyList)                         // 获取所有自己的 post
			auth.POST("/post", v1.PostCreate)                            // 发表 post
			auth.GET("/post/like", v1.PostLike)                          // 喜欢post
			auth.GET("/post/cancel_like", v1.PostCancelLike)             // 取消喜欢 post
			auth.GET("/post/collection", v1.PostCollection)              // 收藏post
			auth.GET("/post/cancel_collection", v1.PostCancelCollection) // 取消收藏post

			auth.GET("/comments", v1.CommentList)   // 根据post_id获取其comment list
			auth.POST("/comment", v1.CommentCreate) // 添加用户
		}
	}
	return r
}
