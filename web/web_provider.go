package web

import (
	"embed"
	"github.com/go-home-admin/home/app"
	"github.com/go-home-admin/home/bootstrap/providers"
	"github.com/go-home-admin/home/bootstrap/servers"
	"io/fs"
	"net/http"
)

//go:embed dist
var static embed.FS

// Provider @Bean
type Provider struct {
	http  *servers.Http            `inject:""`
	route *providers.RouteProvider `inject:""`
}

func (s *Provider) Boot() {
	fSys, err := fs.Sub(static, "dist")
	if err != nil {
		return
	}
	s.http.StaticFS("/web", http.FS(fSys))

	s.http.StaticFS("/files", http.Dir(app.Config("filesystem.disks.local.root", "/files")))
}
