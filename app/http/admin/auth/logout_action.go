package auth

import (
	gin "github.com/gin-gonic/gin"
	admin "github.com/go-home-admin/go-admin/generate/proto/admin"
	http "github.com/go-home-admin/home/app/http"
)

// Logout   退出登陆
func (receiver *Controller) Logout(req *admin.LoginRequest, ctx http.Context) (*admin.LoginResponse, error) {
	// TODO 这里写业务
	return &admin.LoginResponse{}, nil
}

// GinHandleLogout gin原始路由处理
// http.Post(/auth/logout)
func (receiver *Controller) GinHandleLogout(ctx *gin.Context) {
	req := &admin.LoginRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}
	
	resp, err := receiver.Logout(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
