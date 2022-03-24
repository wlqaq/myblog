package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/wlqaq/myblog/database"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	_ "github.com/wlqaq/myblog/database"
	"net/http"
)

type User struct {
	ID int
	Name string
}
type Wgp_test struct {
	gorm.Model
	ID int
	Content string
}
func main()  {
	var model database.Model
	db:=model.Getconn()
	defer db.Close()
	db.HasTable(&User{})
	db.CreateTable(&Wgp_test{})
	r:=gin.Default()
	r.GET("/",hand)
	r.Run(":9911")
}
func hand(c *gin.Context)  {
	var model database.Model
	db:=model.Getconn()
	defer db.Close()
	var user User
	db.First(&user)
	c.JSON(http.StatusOK,user)
}
