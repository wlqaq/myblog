package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"github.com/jinzhu/gorm"
	indexController "github.com/wlqaq/myblog/controller"
	"github.com/wlqaq/myblog/database"
	"github.com/wlqaq/myblog/middleware"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/wlqaq/myblog/database"
	"net/http"
)

type User struct {
	ID   int
	Name string
}
type Wgp_test1 struct {
	gorm.Model
	Content string
}

func main() {
	rc := database.GetRc()
	for i := 0; i < 100000; i++ {
		rc.Do("Set", uuid.New(), i)
	}
	rc.Close()
	db := database.Getconn()
	defer db.Close()
	//err := db.AutoMigrate().CreateTable(&Wgp_test1{}).Error
	//if err!=nil {
	//	panic(err)
	//}
	r := gin.Default()
	r.POST("/login", indexController.Login)
	r.GET("/", middleware.MiddleWare(), hand)
	r.GET("/index", indexController.Handindex)
	r.Run(":9912")
}

func hand(c *gin.Context) {
	db := database.Getconn()
	var user User
	db.First(&user)
	c.JSON(http.StatusOK, user)
}
