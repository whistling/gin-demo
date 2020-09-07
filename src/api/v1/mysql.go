package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type User struct {
	Id    int
	Name  string
	Age   int
	Sex   byte
	Phone string
}

var (
	db *gorm.DB
)

func init() {
	db, err = gorm.Open("mysql", "homestead:secret@tcp(192.168.10.10:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	defer db.Close()
}

func Insert(c *gin.Context) {
	user := User{Name: "Bob", Age: 22}
	res := db.Table("user").Create(user)
}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}

func Find(c *gin.Context) {

}

func CreateTable(c *gin.Context) {
	res := db.Exec("CREATE TABLE `user1` ( `id` bigint(20) NOT NULL AUTO_INCREMENT,  `name` varchar(30) NOT NULL DEFAULT '',  `age` int(3) NOT NULL DEFAULT '0',  `sex` tinyint(3) NOT NULL DEFAULT '0',  `phone` varchar(40) NOT NULL DEFAULT '',  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,  PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4")

	fmt.Println(res)
}
