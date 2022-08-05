package models

import (
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductsModel = (*customProductsModel)(nil)

type (
	// ProductsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductsModel.
	ProductsModel interface {
		productsModel
		RowBuilder() squirrel.SelectBuilder
	}

	customProductsModel struct {
		*defaultProductsModel
	}
)

func (c customProductsModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(productsRows).From(c.table)
}

// NewProductsModel returns a model for the database table.
func NewProductsModel(conn sqlx.SqlConn, c cache.CacheConf) ProductsModel {
	return &customProductsModel{
		defaultProductsModel: newProductsModel(conn, c),
	}
}
