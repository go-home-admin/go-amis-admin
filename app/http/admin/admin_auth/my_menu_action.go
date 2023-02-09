package admin_auth

import (
	gin "github.com/gin-gonic/gin"
	admin2 "github.com/go-home-admin/go-admin/app/entity/admin"
	admin "github.com/go-home-admin/go-admin/generate/proto/admin"
	http "github.com/go-home-admin/home/app/http"
)

// MyMenu   我的菜单
func (receiver *Controller) MyMenu(req *admin.MyMenuRequest, ctx http.Context) (*admin.MyMenuResponse, error) {
	menusSli := make([]*admin.MenuInfo, 0)
	menusSliChildren := make([]*admin.MenuInfo, 0)
	menusMap := make(map[uint32]*admin.MenuInfo, 0)

	for _, data := range admin2.NewOrmAdminMenu().Order("sort DESC").Get() {
		meta := &admin.Meta{}
		data.Meta.Trans(meta)
		Tdata := &admin.MenuInfo{
			Id:        data.Id,
			ParentId:  data.ParentId,
			Path:      *data.Path,
			Hidden:    data.Hidden == 1,
			Name:      data.Name,
			Redirect:  *data.Redirect,
			Component: data.Component,
			Meta:      meta,
			Children:  make([]*admin.MenuInfo, 0),
		}
		if data.ParentId == 0 {
			menusSli = append(menusSli, Tdata)
		} else {
			menusSliChildren = append(menusSliChildren, Tdata)
		}
		menusMap[data.Id] = Tdata
	}

	for _, data := range menusSliChildren {
		if ParentData, ok := menusMap[data.ParentId]; ok {
			ParentData.Children = append(ParentData.Children, data)
		}
	}

	return &admin.MyMenuResponse{
		Menus: menusSli,
	}, nil
}

// GinHandleMyMenu gin原始路由处理
// http.Post(/auth/logout)
func (receiver *Controller) GinHandleMyMenu(ctx *gin.Context) {
	req := &admin.MyMenuRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.MyMenu(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
