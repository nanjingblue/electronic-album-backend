package routers

import (
	v1 "electronic-album/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/ping", v1.Ping)
		apiv1.POST("/register", v1.Register)
	}
	return r
}
