package admin_user

import (
	gin "github.com/gin-gonic/gin"
	admin "github.com/go-home-admin/go-admin/generate/proto/admin"
	http "github.com/go-home-admin/home/app/http"
)

// GetInfo   登陆信息
func (receiver *Controller) GetInfo(req *admin.GetInfoRequest, ctx http.Context) (*admin.GetInfoResponse, error) {
	// TODO 这里写业务
	return &admin.GetInfoResponse{
		Name:         "test",
		Avatar:       "https://avatars.githubusercontent.com/u/18717080?s=30&v=4",
		Roles:        "admin",
		Introduction: "",
	}, nil
}

// GinHandleGetInfo gin原始路由处理
// http.Get(/auth/info)
func (receiver *Controller) GinHandleGetInfo(ctx *gin.Context) {
	req := &admin.GetInfoRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.GetInfo(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
