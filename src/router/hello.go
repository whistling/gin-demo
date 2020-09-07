package router

import (
	v1 "gin/src/api/v1"
	"github.com/gin-gonic/gin"
)

func InitHelloRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	DefaultRouter := Router.Group("")
	{
		DefaultRouter.GET("/", v1.Hello)
	}
	return DefaultRouter
}
