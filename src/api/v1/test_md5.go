package v1

import (
	"gin/src/utils"
	"gin/src/utils/response"
	"github.com/gin-gonic/gin"
)

func Md5(c *gin.Context)  {
	str := c.Query("string")
	val := utils.Md5([]byte(str))
	response.OkWithData(gin.H{
		"string": str,
		"md5":    val,
	}, c)
}
