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
	r.Use(middleware.CurrentUser())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/ping", v1.Ping)
		apiv1.POST("/register", v1.UserRegister)
		apiv1.POST("/login", v1.UserLogin)

		auth := apiv1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			auth.GET("/user", v1.UserInfo)
		}
	}
	return r
}
