// gen for home toolset
package web

import (
	providers "github.com/go-home-admin/home/bootstrap/providers"
	servers "github.com/go-home-admin/home/bootstrap/servers"
)

var _ProviderSingle *Provider

func GetAllProvider() []interface{} {
	return []interface{}{
		NewProvider(),
	}
}

func NewProvider() *Provider {
	if _ProviderSingle == nil {
		_ProviderSingle = &Provider{}
		_ProviderSingle.http = servers.NewHttp()
		_ProviderSingle.route = providers.NewRouteProvider()
		providers.AfterProvider(_ProviderSingle, "")
	}
	return _ProviderSingle
}
