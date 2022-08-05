package models

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProjectsModel = (*customProjectsModel)(nil)

type (
	// ProjectsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProjectsModel.
	ProjectsModel interface {
		projectsModel
		RowBuilder() squirrel.SelectBuilder
		FindOneByQuery(ctx context.Context, query string) (*Projects, error)
	}

	customProjectsModel struct {
		*defaultProjectsModel
	}
)

func (c customProjectsModel) FindOneByQuery(ctx context.Context,query string) (*Projects, error) {

	query, values, err := c.RowBuilder().Where(" name = $1 ",query).ToSql()
	logx.Infof("query:%s  values:%s err:%s", query, values, err)

	var resp Projects

	err = c.QueryRowNoCacheCtx(ctx, &resp, query, values...)

	logx.Infof("resp===================>:%v",resp)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

func (c customProjectsModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(projectsRows).From(c.table)
}

// NewProjectsModel returns a model for the database table.
func NewProjectsModel(conn sqlx.SqlConn, c cache.CacheConf) ProjectsModel {
	return &customProjectsModel{
		defaultProjectsModel: newProjectsModel(conn, c),
	}
}
