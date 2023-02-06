package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/bootstrap/services/app"
	"net/http"
)

// Cors 跨域
func Cors() gin.HandlerFunc {
	origin := ""
	origins := app.GetBean("config").(app.Bean).GetBean("cors.origin")
	if sli, ok := origins.([]interface{}); ok {
		for i, s := range sli {
			if i == 0 {
				origin = origin + s.(string)
			} else {
				origin = origin + ", " + s.(string)
			}
		}
	}

	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, *")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
