package admin_user

import (
	"errors"
	gin "github.com/gin-gonic/gin"
	admin2 "github.com/go-home-admin/go-admin/app/entity/admin"
	admin "github.com/go-home-admin/go-admin/generate/proto/admin"
	http "github.com/go-home-admin/home/app/http"
)

// GetInfo   登陆信息
func (c *Controller) GetInfo(req *admin.GetInfoRequest, ctx http.Context) (*admin.GetInfoResponse, error) {
	userID := ctx.Id()
	user, has := admin2.NewOrmAdminUsers().WhereId(uint32(userID)).First()
	if !has {
		return nil, errors.New("错误的用户信息")
	}
	avatar := ""
	if user.Avatar != nil {
		avatar = *user.Avatar
	}
	return &admin.GetInfoResponse{
		Name:         user.Name,
		Avatar:       avatar,
		Roles:        "admin",
		Introduction: "",
	}, nil
}

// GinHandleGetInfo gin原始路由处理
// http.Get(/auth/info)
func (c *Controller) GinHandleGetInfo(ctx *gin.Context) {
	req := &admin.GetInfoRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := c.GetInfo(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
