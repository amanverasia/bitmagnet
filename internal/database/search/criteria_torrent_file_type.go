package search

import (
	"github.com/amanverasia/bitmagnet/internal/database/query"
	"github.com/amanverasia/bitmagnet/internal/model"
)

func TorrentFileTypeCriteria(fileTypes ...model.FileType) query.Criteria {
	var extensions []string
	for _, fileType := range fileTypes {
		extensions = append(extensions, fileType.Extensions()...)
	}

	return TorrentFileExtensionCriteria(extensions...)
}
