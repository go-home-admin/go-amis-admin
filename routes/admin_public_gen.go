// gen for home toolset
package routes

import (
	gin "github.com/gin-gonic/gin"
	auth "github.com/go-home-admin/go-admin/app/http/admin/auth"
	api "github.com/go-home-admin/home/bootstrap/http/api"
)

// @Bean
type AdminPublicRoutes struct {
	auth *auth.Controller `inject:""`
}

func (c *AdminPublicRoutes) GetGroup() string {
	return "admin-public"
}
func (c *AdminPublicRoutes) GetRoutes() map[*api.Config]func(c *gin.Context) {
	return map[*api.Config]func(c *gin.Context){
		api.Post("/auth/login"):  c.auth.GinHandleLogin,
		api.Post("/auth/logout"): c.auth.GinHandleLogout,
	}
}
