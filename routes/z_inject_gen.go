// gen for home toolset
package routes

import (
	admin_auth "github.com/go-home-admin/go-admin/app/http/admin/admin_auth"
	admin_menu "github.com/go-home-admin/go-admin/app/http/admin/admin_menu"
	admin_role "github.com/go-home-admin/go-admin/app/http/admin/admin_role"
	admin_user "github.com/go-home-admin/go-admin/app/http/admin/admin_user"
	sys_queue "github.com/go-home-admin/go-admin/app/http/admin/sys_queue"
	api_demo "github.com/go-home-admin/go-admin/app/http/api/api_demo"
	public "github.com/go-home-admin/go-admin/app/http/api/public"
	providers "github.com/go-home-admin/home/bootstrap/providers"
)

var _AdminPublicRoutesSingle *AdminPublicRoutes
var _AdminRoutesSingle *AdminRoutes
var _ApiRoutesSingle *ApiRoutes

func GetAllProvider() []interface{} {
	return []interface{}{
		NewAdminPublicRoutes(),
		NewAdminRoutes(),
		NewApiRoutes(),
	}
}

func NewAdminPublicRoutes() *AdminPublicRoutes {
	if _AdminPublicRoutesSingle == nil {
		_AdminPublicRoutesSingle = &AdminPublicRoutes{}
		_AdminPublicRoutesSingle.admin_auth = admin_auth.NewController()
		providers.AfterProvider(_AdminPublicRoutesSingle, "")
	}
	return _AdminPublicRoutesSingle
}
func NewAdminRoutes() *AdminRoutes {
	if _AdminRoutesSingle == nil {
		_AdminRoutesSingle = &AdminRoutes{}
		_AdminRoutesSingle.admin_menu = admin_menu.NewController()
		_AdminRoutesSingle.admin_role = admin_role.NewController()
		_AdminRoutesSingle.admin_user = admin_user.NewController()
		_AdminRoutesSingle.sys_queue = sys_queue.NewController()
		providers.AfterProvider(_AdminRoutesSingle, "")
	}
	return _AdminRoutesSingle
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
