package search

import (
	"github.com/amanverasia/bitmagnet/internal/database/query"
	"github.com/amanverasia/bitmagnet/internal/model"
	"github.com/amanverasia/bitmagnet/internal/protocol"
)

func TorrentFileInfoHashCriteria(infoHashes ...protocol.ID) query.Criteria {
	return infoHashCriteria(model.TableNameTorrentFile, infoHashes...)
}
