package validationfx

import (
	"github.com/amanverasia/bitmagnet/internal/validation"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"validation",
		fx.Provide(validation.New),
	)
}
