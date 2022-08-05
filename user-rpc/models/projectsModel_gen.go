// Code generated by goctl. DO NOT EDIT!

package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
	"zero-demo/genModel"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	projectsFieldNames          = builder.RawFieldNames(&Projects{}, true)
	projectsRows                = strings.Join(projectsFieldNames, ",")
	projectsRowsExpectAutoSet   = strings.Join(stringx.Remove(projectsFieldNames, "id", "create_time", "update_time", "create_t", "update_at"), ",")
	projectsRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(projectsFieldNames, "id", "create_time", "update_time", "create_at", "update_at"))

	cachePublicProjectsIdPrefix        = "cache:public:projects:id:"
	cachePublicProjectsAccessKeyPrefix = "cache:public:projects:accessKey:"
)

type (
	projectsModel interface {
		Insert(ctx context.Context, data *Projects) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Projects, error)
		FindOneByAccessKey(ctx context.Context, accessKey string) (*Projects, error)
		Update(ctx context.Context, data *Projects) error
		Delete(ctx context.Context, id int64) error
	}

	defaultProjectsModel struct {
		sqlc.CachedConn
		table string
	}

	Projects struct {
		Id           int64          `db:"id"`
		Name         string         `db:"name"`
		Vid          sql.NullInt64  `db:"vid"`
		Type         string         `db:"type"`
		Description  sql.NullString `db:"description"`
		Slug         sql.NullString `db:"slug"`
		Status       int64          `db:"status"`
		AccessKey    string         `db:"access_key"`
		AccessSecret string         `db:"access_secret"`
		CreatedAt    time.Time      `db:"created_at"`
		UpdatedAt    time.Time      `db:"updated_at"`
		IsDeleted    bool           `db:"is_deleted"`
		GroupId      sql.NullInt64  `db:"group_id"`
		OwnerId      sql.NullInt64  `db:"owner_id"`
		MaxThingNum  int64          `db:"max_thing_num"`
	}
)

func newProjectsModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultProjectsModel {
	return &defaultProjectsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      `"public"."projects"`,
	}
}

func (m *defaultProjectsModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	publicProjectsAccessKeyKey := fmt.Sprintf("%s%v", cachePublicProjectsAccessKeyPrefix, data.AccessKey)
	publicProjectsIdKey := fmt.Sprintf("%s%v", cachePublicProjectsIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicProjectsAccessKeyKey, publicProjectsIdKey)
	return err
}

func (m *defaultProjectsModel) FindOne(ctx context.Context, id int64) (*Projects, error) {
	publicProjectsIdKey := fmt.Sprintf("%s%v", cachePublicProjectsIdPrefix, id)
	var resp Projects
	err := m.QueryRowCtx(ctx, &resp, publicProjectsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", projectsRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, genModel.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProjectsModel) FindOneByAccessKey(ctx context.Context, accessKey string) (*Projects, error) {
	publicProjectsAccessKeyKey := fmt.Sprintf("%s%v", cachePublicProjectsAccessKeyPrefix, accessKey)
	var resp Projects
	err := m.QueryRowIndexCtx(ctx, &resp, publicProjectsAccessKeyKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where access_key = $1 limit 1", projectsRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, accessKey); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, genModel.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProjectsModel) Insert(ctx context.Context, data *Projects) (sql.Result, error) {
	publicProjectsAccessKeyKey := fmt.Sprintf("%s%v", cachePublicProjectsAccessKeyPrefix, data.AccessKey)
	publicProjectsIdKey := fmt.Sprintf("%s%v", cachePublicProjectsIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)", m.table, projectsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.Vid, data.Type, data.Description, data.Slug, data.Status, data.AccessKey, data.AccessSecret, data.CreatedAt, data.UpdatedAt, data.IsDeleted, data.GroupId, data.OwnerId, data.MaxThingNum)
	}, publicProjectsAccessKeyKey, publicProjectsIdKey)
	return ret, err
}

func (m *defaultProjectsModel) Update(ctx context.Context, newData *Projects) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	publicProjectsAccessKeyKey := fmt.Sprintf("%s%v", cachePublicProjectsAccessKeyPrefix, data.AccessKey)
	publicProjectsIdKey := fmt.Sprintf("%s%v", cachePublicProjectsIdPrefix, data.Id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, projectsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Id, newData.Name, newData.Vid, newData.Type, newData.Description, newData.Slug, newData.Status, newData.AccessKey, newData.AccessSecret, newData.CreatedAt, newData.UpdatedAt, newData.IsDeleted, newData.GroupId, newData.OwnerId, newData.MaxThingNum)
	}, publicProjectsAccessKeyKey, publicProjectsIdKey)
	return err
}

func (m *defaultProjectsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePublicProjectsIdPrefix, primary)
}

func (m *defaultProjectsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", projectsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultProjectsModel) tableName() string {
	return m.table
}
