// gen for home toolset
package routes

import (
	admin_user "github.com/go-home-admin/go-admin/app/http/admin/admin_user"
	amis "github.com/go-home-admin/go-admin/app/http/admin/amis"
	auth "github.com/go-home-admin/go-admin/app/http/admin/auth"
	menu "github.com/go-home-admin/go-admin/app/http/admin/menu"
	api_demo "github.com/go-home-admin/go-admin/app/http/api/api_demo"
	public "github.com/go-home-admin/go-admin/app/http/api/public"
	providers "github.com/go-home-admin/home/bootstrap/providers"
)

var _AdminRoutesSingle *AdminRoutes
var _AdminPublicRoutesSingle *AdminPublicRoutes
var _ApiRoutesSingle *ApiRoutes

func GetAllProvider() []interface{} {
	return []interface{}{
		NewAdminRoutes(),
		NewAdminPublicRoutes(),
		NewApiRoutes(),
	}
}

func NewAdminRoutes() *AdminRoutes {
	if _AdminRoutesSingle == nil {
		_AdminRoutesSingle = &AdminRoutes{}
		_AdminRoutesSingle.amis = amis.NewController()
		_AdminRoutesSingle.menu = menu.NewController()
		_AdminRoutesSingle.admin_user = admin_user.NewController()
		providers.AfterProvider(_AdminRoutesSingle, "")
	}
	return _AdminRoutesSingle
}
func NewAdminPublicRoutes() *AdminPublicRoutes {
	if _AdminPublicRoutesSingle == nil {
		_AdminPublicRoutesSingle = &AdminPublicRoutes{}
		_AdminPublicRoutesSingle.auth = auth.NewController()
		providers.AfterProvider(_AdminPublicRoutesSingle, "")
	}
	return _AdminPublicRoutesSingle
}
func NewApiRoutes() *ApiRoutes {
	if _ApiRoutesSingle == nil {
		_ApiRoutesSingle = &ApiRoutes{}
		_ApiRoutesSingle.api_demo = api_demo.NewController()
		_ApiRoutesSingle.public = public.NewController()
		providers.AfterProvider(_ApiRoutesSingle, "")
	}
	return _ApiRoutesSingle
}
