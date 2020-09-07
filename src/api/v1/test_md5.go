package v1

import (
	"gin/src/utils"
	"gin/src/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"time"
)

func Md5(c *gin.Context) {
	str := c.Query("string")
	val := utils.Md5([]byte(str))
	response.OkWithData(gin.H{
		"string": str,
		"md5":    val,
	}, c)
}

func Pong(c *gin.Context) {
	response.OkWithData("pong", c)
}

func Cache(c *gin.Context) {
	redisCon := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	)
	key := c.Query("key")
	//pong, err := redisCon.Ping(c).Result()
	res := redisCon.Set(c, key, time.Now().Format("2006-01-02 15:04:05"), 86400*time.Second).Val()
	redisCon.HSet(c, "days", map[string]interface{}{"1": "Mon", "2": "Tue", "3": "Wed"})

	response.OkWithData(gin.H{
		"result": res,
	}, c)
}

func GetCache(c *gin.Context) {
	con := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
	key := c.Query("key")
	val := con.Get(c, key).Val()
	i := con.HMGet(c, "days", "1", "2", "3").Val()
	response.OkWithData(gin.H{
		"key":  key,
		"val":  val,
		"mget": i,
	}, c)
}
