package models

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ThingsModel = (*customThingsModel)(nil)

type (
	// ThingsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customThingsModel.
	ThingsModel interface {
		thingsModel
	}

	customThingsModel struct {
		*defaultThingsModel
	}
)

// NewThingsModel returns a model for the database table.
func NewThingsModel(conn sqlx.SqlConn, c cache.CacheConf) ThingsModel {
	return &customThingsModel{
		defaultThingsModel: newThingsModel(conn, c),
	}
}
