package processorfx

import (
	"github.com/amanverasia/bitmagnet/internal/processor"
	batchqueue "github.com/amanverasia/bitmagnet/internal/processor/batch/queue"
	processorqueue "github.com/amanverasia/bitmagnet/internal/processor/queue"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"processor",
		fx.Provide(
			processor.New,
			processorqueue.New,
			batchqueue.New,
		),
	)
}
