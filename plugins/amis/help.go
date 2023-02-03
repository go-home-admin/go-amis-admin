package amis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/app"
)

func GetUrl(ctx *gin.Context) string {
	domain := app.Config("app.url", "http://127.0.0.1")
	return domain + ctx.Request.URL.RequestURI()
}
