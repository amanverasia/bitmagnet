package banning

import (
	"errors"

	"github.com/amanverasia/bitmagnet/internal/protocol/metainfo"
)

type sizeChecker struct {
	min int64
}

func (c sizeChecker) Check(info metainfo.Info) error {
	if info.TotalLength() < c.min {
		return errors.New("size too small")
	}

	return nil
}
