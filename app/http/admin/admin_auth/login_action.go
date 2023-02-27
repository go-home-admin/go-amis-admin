package admin_auth

import (
	"errors"
	gin "github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/app"
	admin2 "github.com/go-home-admin/go-admin/app/entity/admin"
	"github.com/go-home-admin/go-admin/app/services/auth"
	admin "github.com/go-home-admin/go-admin/generate/proto/admin"
	http "github.com/go-home-admin/home/app/http"
)

// Login   登陆
func (receiver *Controller) Login(req *admin.LoginRequest, ctx http.Context) (*admin.LoginResponse, error) {
	user, has := admin2.NewOrmAdminUsers().WhereUsername(req.Username).First()
	if !has {
		return nil, errors.New("用户不存在")
	}
	if user.Password != app.MD5(req.Password) {
		admin2.NewOrmAdminUsers().WhereUsername(req.Username).Update("password", app.MD5(req.Password))
		return nil, errors.New("密码不存在")
	}

	token, err := auth.NewJwt().GenerateToken(int(user.Id))
	if err != nil {
		return nil, err
	}
	return &admin.LoginResponse{
		Token: token,
	}, nil
}

// GinHandleLogin gin原始路由处理
// http.Post(/auth/login)
func (receiver *Controller) GinHandleLogin(ctx *gin.Context) {
	req := &admin.LoginRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.Login(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
