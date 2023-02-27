package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/app/services/auth"
	http2 "github.com/go-home-admin/home/app/http"
	"github.com/go-home-admin/home/bootstrap/services/app"
	"net/http"
	"strings"
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

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Token")

		if token == "" {
			token = c.GetHeader("Authorization")
			if strings.HasPrefix(token, "Bearer ") {
				token = token[7:]
			}
			if token == "" {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		}

		userID, err := auth.NewJwt().GetUid(token)
		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
		}
		c.Set(http2.UserIdKey, uint64(userID))
		c.Next()
	}
}
