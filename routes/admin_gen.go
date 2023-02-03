// gen for home toolset
package routes

import (
	gin "github.com/gin-gonic/gin"
	admin_user "github.com/go-home-admin/go-admin/app/http/admin/admin_user"
	amis "github.com/go-home-admin/go-admin/app/http/admin/amis"
	menu "github.com/go-home-admin/go-admin/app/http/admin/menu"
	api "github.com/go-home-admin/home/bootstrap/http/api"
)

// @Bean
type AdminRoutes struct {
	amis       *amis.Controller       `inject:""`
	menu       *menu.Controller       `inject:""`
	admin_user *admin_user.Controller `inject:""`
}

func (c *AdminRoutes) GetGroup() string {
	return "admin"
}
func (c *AdminRoutes) GetRoutes() map[*api.Config]func(c *gin.Context) {
	return map[*api.Config]func(c *gin.Context){
		api.Get("/amis/tabs"):    c.amis.GinHandleDemo,
		api.Get("/menu"):         c.menu.GinHandleCurd,
		api.Any("/menu/:action"): c.menu.GinHandleCurd,
		api.Get("/auth/info"):    c.admin_user.GinHandleGetInfo,
	}
}
