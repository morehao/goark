package iamdao

import (
	"context"

	"github.com/morehao/goark/pkg/dbclient"
	"gorm.io/gorm"
)

type Base struct {
	Tx *gorm.DB
}

// DB 获取DB
func (base *Base) DB(ctx context.Context) (db *gorm.DB) {
	if base.Tx != nil {
		return base.Tx.WithContext(ctx)
	}

	db = dbclient.DBIam.WithContext(ctx)
	return
}
