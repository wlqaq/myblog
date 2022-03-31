package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header
		if header.Get("Token") == "" {
			fmt.Println("没有登录")
		}
		fmt.Println(header.Get("Token"))
	}
}
