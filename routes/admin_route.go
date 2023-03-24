// gen for home toolset
package routes

import (
	home_gin_1 "github.com/gin-gonic/gin"
	admin_menu "github.com/go-home-admin/go-admin/app/http/admin/admin_menu"
	admin_role "github.com/go-home-admin/go-admin/app/http/admin/admin_role"
	admin_user "github.com/go-home-admin/go-admin/app/http/admin/admin_user"
	sys_queue "github.com/go-home-admin/go-admin/app/http/admin/sys_queue"
	home_api_1 "github.com/go-home-admin/home/bootstrap/http/api"
)

// @Bean
type AdminRoutes struct {
	admin_menu *admin_menu.Controller `inject:""`
	admin_role *admin_role.Controller `inject:""`
	admin_user *admin_user.Controller `inject:""`
	sys_queue  *sys_queue.Controller  `inject:""`
}

func (c *AdminRoutes) GetGroup() string {
	return "admin"
}
func (c *AdminRoutes) GetRoutes() map[*home_api_1.Config]func(c *home_gin_1.Context) {
	return map[*home_api_1.Config]func(c *home_gin_1.Context){
		home_api_1.Get("menus"):             c.admin_menu.GinHandleCurd,
		home_api_1.Post("menus"):            c.admin_menu.GinHandleCurd,
		home_api_1.Any("menus/:action"):     c.admin_menu.GinHandleCurd,
		home_api_1.Get("roles"):             c.admin_role.GinHandleCurd,
		home_api_1.Post("roles"):            c.admin_role.GinHandleCurd,
		home_api_1.Any("roles/:action"):     c.admin_role.GinHandleCurd,
		home_api_1.Get("users"):             c.admin_user.GinHandleCurd,
		home_api_1.Post("users"):            c.admin_user.GinHandleCurd,
		home_api_1.Any("users/:action"):     c.admin_user.GinHandleCurd,
		home_api_1.Get("/auth/info"):        c.admin_user.GinHandleGetInfo,
		home_api_1.Get("sys/queue"):         c.sys_queue.GinHandleCurd,
		home_api_1.Post("sys/queue"):        c.sys_queue.GinHandleCurd,
		home_api_1.Any("sys/queue/:action"): c.sys_queue.GinHandleCurd,
	}
}
