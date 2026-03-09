package appfx

import (
	"github.com/amanverasia/bitmagnet/internal/app/cli"
	"github.com/amanverasia/bitmagnet/internal/app/cli/args"
	"github.com/amanverasia/bitmagnet/internal/app/cli/hooks"
	"github.com/amanverasia/bitmagnet/internal/app/cmd/classifiercmd"
	"github.com/amanverasia/bitmagnet/internal/app/cmd/configcmd"
	"github.com/amanverasia/bitmagnet/internal/app/cmd/healthcheckcmd"
	"github.com/amanverasia/bitmagnet/internal/app/cmd/processcmd"
	"github.com/amanverasia/bitmagnet/internal/app/cmd/reprocesscmd"
	"github.com/amanverasia/bitmagnet/internal/app/cmd/workercmd"
	"github.com/amanverasia/bitmagnet/internal/blocking/blockingfx"
	"github.com/amanverasia/bitmagnet/internal/classifier/classifierfx"
	"github.com/amanverasia/bitmagnet/internal/config/configfx"
	"github.com/amanverasia/bitmagnet/internal/database/databasefx"
	"github.com/amanverasia/bitmagnet/internal/database/migrations"
	"github.com/amanverasia/bitmagnet/internal/dhtcrawler/dhtcrawlerfx"
	"github.com/amanverasia/bitmagnet/internal/gql/gqlfx"
	"github.com/amanverasia/bitmagnet/internal/health/healthfx"
	"github.com/amanverasia/bitmagnet/internal/httpserver/httpserverfx"
	"github.com/amanverasia/bitmagnet/internal/importer/importerfx"
	"github.com/amanverasia/bitmagnet/internal/logging/loggingfx"
	"github.com/amanverasia/bitmagnet/internal/metrics/metricsfx"
	"github.com/amanverasia/bitmagnet/internal/processor/processorfx"
	"github.com/amanverasia/bitmagnet/internal/protocol/dht/dhtfx"
	"github.com/amanverasia/bitmagnet/internal/protocol/metainfo/metainfofx"
	"github.com/amanverasia/bitmagnet/internal/queue/queuefx"
	"github.com/amanverasia/bitmagnet/internal/telemetry/telemetryfx"
	"github.com/amanverasia/bitmagnet/internal/tmdb/tmdbfx"
	"github.com/amanverasia/bitmagnet/internal/torznab/torznabfx"
	"github.com/amanverasia/bitmagnet/internal/validation/validationfx"
	"github.com/amanverasia/bitmagnet/internal/version/versionfx"
	"github.com/amanverasia/bitmagnet/internal/webui"
	"github.com/amanverasia/bitmagnet/internal/worker/workerfx"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"app",
		blockingfx.New(),
		classifierfx.New(),
		configfx.New(),
		dhtcrawlerfx.New(),
		dhtfx.New(),
		databasefx.New(),
		gqlfx.New(),
		healthfx.New(),
		httpserverfx.New(),
		importerfx.New(),
		loggingfx.New(),
		metainfofx.New(),
		metricsfx.New(),
		processorfx.New(),
		queuefx.New(),
		telemetryfx.New(),
		tmdbfx.New(),
		torznabfx.New(),
		validationfx.New(),
		versionfx.New(),
		workerfx.New(),
		fx.Provide(
			args.New,
			cli.New,
			hooks.New,
			// cli commands:
			classifiercmd.New,
			configcmd.New,
			healthcheckcmd.New,
			reprocesscmd.New,
			processcmd.New,
			workercmd.New,
		),
		fx.Provide(webui.New),
		fx.Decorate(migrations.NewDecorator),
	)
}
