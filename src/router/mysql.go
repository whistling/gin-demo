package router

import (
	v1 "gin/src/api/v1"
	"github.com/gin-gonic/gin"
)

func InitMysqlRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	MysqlRouter := Router.Group("mysql")
	{
		MysqlRouter.GET("insert", v1.Insert)
		MysqlRouter.GET("update", v1.Update)
		MysqlRouter.GET("delete", v1.Delete)
		MysqlRouter.GET("find", v1.Find)
	}
	return MysqlRouter
}
