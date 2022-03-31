package indexController

import (
	"github.com/gin-gonic/gin"
	"github.com/wlqaq/myblog/database"
	"github.com/wlqaq/myblog/model"
	"net/http"
)

func Login(r *gin.Context) {
	userame := r.PostForm("username")
	//passwd := r.PostForm("passwd")
	db := database.Getconn()

	var user model.User
	if db.Where("name=?", userame).First(&user).Error == nil {
		r.JSON(http.StatusOK, "用户名不存在")
	}

}
