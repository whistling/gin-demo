package initialize

import (
	"gin/src/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()

	ApiGroup := Router.Group("")
	router.InitTestRouter(ApiGroup)
	router.InitMysqlRouter(ApiGroup)
	router.InitClickHouseRouter(ApiGroup)
	router.InitHelloRouter(ApiGroup)
	router.InitKafkaRouter(ApiGroup)
	return Router
}
