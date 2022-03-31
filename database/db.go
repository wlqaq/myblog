package database

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var db *gorm.DB
var rc redis.Conn

func redisc() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	rc = c
	fmt.Println("redis conn success")

}
func init() {

	redisc()
	dbc, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	dbc.SingularTable(true)
	duration, err := time.ParseDuration("1m")
	if err != nil {
		fmt.Println(err)
	}
	sqlDB := dbc.DB()
	sqlDB.SetMaxOpenConns(1000)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetConnMaxIdleTime(duration)
	db = dbc
}
func Getconn() *gorm.DB {
	return db
}
func GetRc() redis.Conn {
	return rc
}
