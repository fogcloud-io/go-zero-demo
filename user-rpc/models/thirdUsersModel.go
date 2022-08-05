package models

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ThirdUsersModel = (*customThirdUsersModel)(nil)

type (
	// ThirdUsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customThirdUsersModel.
	ThirdUsersModel interface {
		thirdUsersModel
	}

	customThirdUsersModel struct {
		*defaultThirdUsersModel
	}
)

// NewThirdUsersModel returns a model for the database table.
func NewThirdUsersModel(conn sqlx.SqlConn, c cache.CacheConf) ThirdUsersModel {
	return &customThirdUsersModel{
		defaultThirdUsersModel: newThirdUsersModel(conn, c),
	}
}
