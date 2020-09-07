package router

import (
	v1 "gin/src/api/v1"
	"github.com/gin-gonic/gin"
)

func InitKafkaRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	KafkaRouter := Router.Group("kafka")
	{
		KafkaRouter.GET("/produce", v1.Produce)
		KafkaRouter.GET("/consume", v1.Consume)
	}
	return KafkaRouter
}
