package httpserverfx

import (
	"github.com/amanverasia/bitmagnet/internal/config/configfx"
	"github.com/amanverasia/bitmagnet/internal/httpserver"
	"github.com/amanverasia/bitmagnet/internal/httpserver/cors"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"http_server",
		configfx.NewConfigModule[httpserver.Config]("http_server", httpserver.NewDefaultConfig()),
		fx.Provide(
			httpserver.New,
			cors.New,
		),
	)
}
