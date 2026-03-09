package classifier

import (
	"github.com/amanverasia/bitmagnet/internal/tmdb"
)

type dependencies struct {
	search     LocalSearch
	tmdbClient tmdb.Client
}
