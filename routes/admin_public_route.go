// gen for home toolset
package routes

import (
	home_gin_1 "github.com/gin-gonic/gin"
	admin_auth "github.com/go-home-admin/go-admin/app/http/admin/admin_auth"
	home_api_1 "github.com/go-home-admin/home/bootstrap/http/api"
)

// @Bean
type AdminPublicRoutes struct {
	admin_auth *admin_auth.Controller `inject:""`
}

func (c *AdminPublicRoutes) GetGroup() string {
	return "admin-public"
}
func (c *AdminPublicRoutes) GetRoutes() map[*home_api_1.Config]func(c *home_gin_1.Context) {
	return map[*home_api_1.Config]func(c *home_gin_1.Context){
		home_api_1.Post("/auth/login"):  c.admin_auth.GinHandleLogin,
		home_api_1.Post("/auth/logout"): c.admin_auth.GinHandleLogout,
		home_api_1.Get("/auth/menus"):   c.admin_auth.GinHandleMyMenu,
	}
}
