package v1

import (
	"gin/src/utils/response"
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	response.OkWithData("hello world", c)
}
