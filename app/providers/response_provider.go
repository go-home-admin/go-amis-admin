package providers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/app/http"
	http2 "net/http"
	"strings"
)

// Response @Bean
// 重载框架的输出
type Response struct{}

func (r *Response) Init() {
	http.NewContext = func(ctx *gin.Context) http.Context {
		return &Ctx{
			Context: ctx,
		}
	}
}

type Ctx struct {
	*gin.Context
	UserID   uint64
	UserInfo interface{}
}

func (receiver Ctx) Success(data interface{}) {
	receiver.JSON(http2.StatusOK, map[string]interface{}{
		"data":   data,
		"status": 0,
		"msg":    "",
	})
}

func (receiver Ctx) Fail(err error) {
	receiver.JSON(http2.StatusOK, map[string]interface{}{
		"status": 502,
		"msg":    err.Error(),
	})
}

func (receiver Ctx) Gin() *gin.Context {
	return receiver.Context
}

func (receiver Ctx) User() interface{} {
	u, ok := receiver.Context.Get(http.UserKey)
	if !ok {
		return nil
	}
	return u
}

func (receiver Ctx) Id() uint64 {
	u, ok := receiver.Context.Get(http.UserIdKey)
	if !ok {
		return 0
	}
	return u.(uint64)
}

func (receiver Ctx) Token() string {
	tokenString := receiver.Context.GetHeader("Authorization")

	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = tokenString[7:]
	}
	return tokenString
}
