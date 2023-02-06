package amis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/app"
	"github.com/sirupsen/logrus"
	"strconv"
)

// GetUrl action = /list | /edit | /del
func GetUrl(ctx *gin.Context, action string) string {
	domain := app.Config("app.url", "http://127.0.0.1")
	return domain + ctx.Request.URL.RequestURI() + action
}

func GetInt(ctx *gin.Context, k string, def int) int {
	v := ctx.Query(k)
	if v == "" {
		return def
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		logrus.Error("GetInt", err)
		return 0
	}
	return i
}
