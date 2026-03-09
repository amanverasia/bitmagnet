package importerfx

import (
	"github.com/amanverasia/bitmagnet/internal/importer"
	"github.com/amanverasia/bitmagnet/internal/importer/httpserver"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"importer",
		fx.Provide(
			httpserver.New,
			importer.New,
		),
	)
}
