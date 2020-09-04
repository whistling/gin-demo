package main

import (
	"gin/src/core"
)

func main() {
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	response.OkWithData( gin.H{
	//		"message": "pong",
	//	}, c)
	//})
	core.RunServer()
}
