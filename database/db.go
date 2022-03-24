package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Model struct {

}
func (c *Model) Getconn() *gorm.DB {
	dbc, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err!= nil{
		panic(err)
	}
	dbc.SingularTable(true)
	duration, err := time.ParseDuration("1m")
	if err != nil {
		fmt.Println(err)
	}
	dbc.DB().SetConnMaxIdleTime(duration)
	return dbc
}
