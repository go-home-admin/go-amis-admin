// gen for home toolset
package routes

import (
	home_gin_1 "github.com/gin-gonic/gin"
	api_demo "github.com/go-home-admin/go-admin/app/http/api/api_demo"
	public "github.com/go-home-admin/go-admin/app/http/api/public"
	home_api_1 "github.com/go-home-admin/home/bootstrap/http/api"
)

// @Bean
type ApiRoutes struct {
	api_demo *api_demo.Controller `inject:""`
	public   *public.Controller   `inject:""`
}

func (c *ApiRoutes) GetGroup() string {
	return "api"
}
func (c *ApiRoutes) GetRoutes() map[*home_api_1.Config]func(c *home_gin_1.Context) {
	return map[*home_api_1.Config]func(c *home_gin_1.Context){
		home_api_1.Get("/api/demo"): c.api_demo.GinHandleHome,
		home_api_1.Get("/home"):     c.public.GinHandleHome,
	}
}
