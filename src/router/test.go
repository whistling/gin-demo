package router

import (
	v1 "gin/src/api/v1"
	"github.com/gin-gonic/gin"
)

func InitTestRouter(Router *gin.RouterGroup)(R gin.IRoutes) {

	TestRouter := Router.Group("test")
	{
		TestRouter.GET("md5", v1.Md5)
	}
	return TestRouter
}
