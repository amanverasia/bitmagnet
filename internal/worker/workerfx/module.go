package workerfx

import (
	"github.com/amanverasia/bitmagnet/internal/worker"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"worker",
		fx.Provide(worker.NewRegistry),
	)
}
