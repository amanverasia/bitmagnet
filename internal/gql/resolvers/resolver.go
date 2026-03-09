package resolvers

import (
	"github.com/amanverasia/bitmagnet/internal/blocking"
	"github.com/amanverasia/bitmagnet/internal/database/dao"
	"github.com/amanverasia/bitmagnet/internal/database/search"
	"github.com/amanverasia/bitmagnet/internal/health"
	"github.com/amanverasia/bitmagnet/internal/metrics/queuemetrics"
	"github.com/amanverasia/bitmagnet/internal/metrics/torrentmetrics"
	"github.com/amanverasia/bitmagnet/internal/processor"
	"github.com/amanverasia/bitmagnet/internal/queue/manager"
	"github.com/amanverasia/bitmagnet/internal/worker"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Dao                  *dao.Query
	Search               search.Search
	Workers              worker.Registry
	Checker              health.Checker
	QueueMetricsClient   queuemetrics.Client
	QueueManager         manager.Manager
	TorrentMetricsClient torrentmetrics.Client
	Processor            processor.Processor
	BlockingManager      blocking.Manager
}
