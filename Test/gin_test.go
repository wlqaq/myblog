package Test

import (
   "github.com/jinzhu/gorm"
   "github.com/wlqaq/myblog/database"
   "testing"
)

type User struct {
   ID int
   Name string
}
type Wgptest struct {
   gorm.Model
   ID int
   Content string
}
func Testgin(t *testing.T)  {
   var dbc database.Model
   db := dbc.Getconn()
   var user User
   db.First(&user)
   var wg Wgptest
   db.HasTable(&wg)
   db.CreateTable(&wg)
   defer db.Close()
}
