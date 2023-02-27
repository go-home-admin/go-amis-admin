// gen for home toolset
package auth

import (
	providers "github.com/go-home-admin/home/bootstrap/providers"
)

var _JwtSingle *Jwt

func GetAllProvider() []interface{} {
	return []interface{}{
		NewJwt(),
	}
}

func NewJwt() *Jwt {
	if _JwtSingle == nil {
		_JwtSingle = &Jwt{}
		providers.AfterProvider(_JwtSingle, "")
	}
	return _JwtSingle
}
