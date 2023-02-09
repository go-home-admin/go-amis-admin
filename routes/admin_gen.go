// gen for home toolset
package routes

import (
	gin "github.com/gin-gonic/gin"
	admin_menu "github.com/go-home-admin/go-admin/app/http/admin/admin_menu"
	admin_user "github.com/go-home-admin/go-admin/app/http/admin/admin_user"
	api "github.com/go-home-admin/home/bootstrap/http/api"
)

// @Bean
type AdminRoutes struct {
	admin_menu *admin_menu.Controller `inject:""`
	admin_user *admin_user.Controller `inject:""`
}

func (c *AdminRoutes) GetGroup() string {
	return "admin"
}
func (c *AdminRoutes) GetRoutes() map[*api.Config]func(c *gin.Context) {
	return map[*api.Config]func(c *gin.Context){
		api.Get("menus"):            c.admin_menu.GinHandleCurd,
		api.Post("menus"):           c.admin_menu.GinHandleCurd,
		api.Get("menus/:action"):    c.admin_menu.GinHandleCurd,
		api.Put("menus/:action"):    c.admin_menu.GinHandleCurd,
		api.Delete("menus/:action"): c.admin_menu.GinHandleCurd,
		api.Get("users"):            c.admin_user.GinHandleCurd,
		api.Post("users"):           c.admin_user.GinHandleCurd,
		api.Get("users/:action"):    c.admin_user.GinHandleCurd,
		api.Put("users/:action"):    c.admin_user.GinHandleCurd,
		api.Delete("users/:action"): c.admin_user.GinHandleCurd,
		api.Get("/auth/info"):       c.admin_user.GinHandleGetInfo,
	}
}
