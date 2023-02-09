// gen for home toolset
package routes

import (
	gin "github.com/gin-gonic/gin"
	admin_auth "github.com/go-home-admin/go-admin/app/http/admin/admin_auth"
	api "github.com/go-home-admin/home/bootstrap/http/api"
)

// @Bean
type AdminPublicRoutes struct {
	admin_auth *admin_auth.Controller `inject:""`
}

func (c *AdminPublicRoutes) GetGroup() string {
	return "admin-public"
}
func (c *AdminPublicRoutes) GetRoutes() map[*api.Config]func(c *gin.Context) {
	return map[*api.Config]func(c *gin.Context){
		api.Post("/auth/login"):  c.admin_auth.GinHandleLogin,
		api.Post("/auth/logout"): c.admin_auth.GinHandleLogout,
		api.Get("/auth/menus"):   c.admin_auth.GinHandleMyMenu,
	}
}
