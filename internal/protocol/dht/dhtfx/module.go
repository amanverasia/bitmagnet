package dhtfx

import (
	"github.com/amanverasia/bitmagnet/internal/config/configfx"
	"github.com/amanverasia/bitmagnet/internal/protocol"
	"github.com/amanverasia/bitmagnet/internal/protocol/dht/client"
	"github.com/amanverasia/bitmagnet/internal/protocol/dht/ktable"
	"github.com/amanverasia/bitmagnet/internal/protocol/dht/responder"
	"github.com/amanverasia/bitmagnet/internal/protocol/dht/server"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"dht",
		configfx.NewConfigModule[server.Config]("dht_server", server.NewDefaultConfig()),
		fx.Provide(
			fx.Annotated{
				Name:   "dht_node_id",
				Target: protocol.RandomNodeIDWithClientSuffix,
			},
			client.New,
			ktable.New,
			responder.New,
			server.New,
		),
	)
}
