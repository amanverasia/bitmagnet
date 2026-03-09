package databasefx

import (
	"github.com/amanverasia/bitmagnet/internal/config/configfx"
	"github.com/amanverasia/bitmagnet/internal/database"
	"github.com/amanverasia/bitmagnet/internal/database/cache"
	"github.com/amanverasia/bitmagnet/internal/database/dao"
	"github.com/amanverasia/bitmagnet/internal/database/healthcheck"
	"github.com/amanverasia/bitmagnet/internal/database/migrations"
	"github.com/amanverasia/bitmagnet/internal/database/postgres"
	"github.com/amanverasia/bitmagnet/internal/database/search"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"database",
		configfx.NewConfigModule[postgres.Config]("postgres", postgres.NewDefaultConfig()),
		configfx.NewConfigModule[cache.Config]("gorm_cache", cache.NewDefaultConfig()),
		fx.Provide(
			cache.NewInMemoryCacher,
			cache.NewPlugin,
			dao.New,
			database.New,
			healthcheck.New,
			migrations.New,
			postgres.New,
			search.New,
		),
		fx.Decorate(
			cache.NewDecorator,
		),
	)
}
