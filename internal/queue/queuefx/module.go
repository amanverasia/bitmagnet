package queuefx

import (
	"github.com/amanverasia/bitmagnet/internal/queue/manager"
	"github.com/amanverasia/bitmagnet/internal/queue/prometheus"
	"github.com/amanverasia/bitmagnet/internal/queue/server"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"queue",
		fx.Provide(
			server.New,
			manager.New,
			prometheus.New,
		),
	)
}
