package manager

import (
	"github.com/amanverasia/bitmagnet/internal/database/dao"
	"gorm.io/gorm"
)

type manager struct {
	dao *dao.Query
	db  *gorm.DB
}
