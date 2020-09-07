package router

import (
	v1 "gin/src/api/v1"
	"github.com/gin-gonic/gin"
)

func InitClickHouseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	ClickHouseRouter := Router.Group("click")
	{
		ClickHouseRouter.GET("createTable", v1.ClickCreateTable)
		ClickHouseRouter.GET("dropTable", v1.ClickDropTable)
		ClickHouseRouter.GET("insert", v1.ClickInsert)
		ClickHouseRouter.GET("query", v1.ClickQuery)
	}
	return ClickHouseRouter
}
