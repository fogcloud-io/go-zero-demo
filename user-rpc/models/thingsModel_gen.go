// Code generated by goctl. DO NOT EDIT!

package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"zero-demo/genModel"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	thingsFieldNames          = builder.RawFieldNames(&Things{}, true)
	thingsRows                = strings.Join(thingsFieldNames, ",")
	thingsRowsExpectAutoSet   = strings.Join(stringx.Remove(thingsFieldNames, "id", "create_time", "update_time", "create_t", "update_at"), ",")
	thingsRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(thingsFieldNames, "id", "create_time", "update_time", "create_at", "update_at"))

	cachePublicThingsIdPrefix                   = "cache:public:things:id:"
	cachePublicThingsDeviceIdPrefix             = "cache:public:things:deviceId:"
	cachePublicThingsDeviceNameProductKeyPrefix = "cache:public:things:deviceName:productKey:"
)

type (
	thingsModel interface {
		Insert(ctx context.Context, data *Things) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Things, error)
		FindOneByDeviceId(ctx context.Context, deviceId string) (*Things, error)
		FindOneByDeviceNameProductKey(ctx context.Context, deviceName string, productKey string) (*Things, error)
		Update(ctx context.Context, data *Things) error
		Delete(ctx context.Context, id int64) error
	}

	defaultThingsModel struct {
		sqlc.CachedConn
		table string
	}

	Things struct {
		Id              int64          `db:"id"`
		DeviceId        string         `db:"device_id"`
		ProductKey      string         `db:"product_key"`
		DeviceName      string         `db:"device_name"`
		DeviceSecret    string         `db:"device_secret"`
		Nickname        sql.NullString `db:"nickname"`
		Status          sql.NullString `db:"status"`
		CreatedAt       sql.NullTime   `db:"created_at"`
		RegisterAt      sql.NullTime   `db:"register_at"`
		UpdatedAt       sql.NullTime   `db:"updated_at"`
		FirmwareVersion sql.NullString `db:"firmware_version"`
		Ip              sql.NullString `db:"ip"`
		ClientId        sql.NullString `db:"client_id"`
		Mac             sql.NullString `db:"mac"`
		IsDeleted       bool           `db:"is_deleted"`
		Extra           sql.NullString `db:"extra"`
		IsMocker        bool           `db:"is_mocker"`
		ProductId       sql.NullInt64  `db:"product_id"`
		BatchId         sql.NullInt64  `db:"batch_id"`
		EnabledStatus   string         `db:"enabled_status"`
		ProjectId       sql.NullInt64  `db:"project_id"`
	}
)

func newThingsModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultThingsModel {
	return &defaultThingsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      `"public"."things"`,
	}
}

func (m *defaultThingsModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	publicThingsDeviceIdKey := fmt.Sprintf("%s%v", cachePublicThingsDeviceIdPrefix, data.DeviceId)
	publicThingsDeviceNameProductKeyKey := fmt.Sprintf("%s%v:%v", cachePublicThingsDeviceNameProductKeyPrefix, data.DeviceName, data.ProductKey)
	publicThingsIdKey := fmt.Sprintf("%s%v", cachePublicThingsIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicThingsDeviceIdKey, publicThingsDeviceNameProductKeyKey, publicThingsIdKey)
	return err
}

func (m *defaultThingsModel) FindOne(ctx context.Context, id int64) (*Things, error) {
	publicThingsIdKey := fmt.Sprintf("%s%v", cachePublicThingsIdPrefix, id)
	var resp Things
	err := m.QueryRowCtx(ctx, &resp, publicThingsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", thingsRows, m.table)
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

func (m *defaultThingsModel) FindOneByDeviceId(ctx context.Context, deviceId string) (*Things, error) {
	publicThingsDeviceIdKey := fmt.Sprintf("%s%v", cachePublicThingsDeviceIdPrefix, deviceId)
	var resp Things
	err := m.QueryRowIndexCtx(ctx, &resp, publicThingsDeviceIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where device_id = $1 limit 1", thingsRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, deviceId); err != nil {
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

func (m *defaultThingsModel) FindOneByDeviceNameProductKey(ctx context.Context, deviceName string, productKey string) (*Things, error) {
	publicThingsDeviceNameProductKeyKey := fmt.Sprintf("%s%v:%v", cachePublicThingsDeviceNameProductKeyPrefix, deviceName, productKey)
	var resp Things
	err := m.QueryRowIndexCtx(ctx, &resp, publicThingsDeviceNameProductKeyKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where device_name = $1 and product_key = $2 limit 1", thingsRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, deviceName, productKey); err != nil {
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

func (m *defaultThingsModel) Insert(ctx context.Context, data *Things) (sql.Result, error) {
	publicThingsDeviceIdKey := fmt.Sprintf("%s%v", cachePublicThingsDeviceIdPrefix, data.DeviceId)
	publicThingsDeviceNameProductKeyKey := fmt.Sprintf("%s%v:%v", cachePublicThingsDeviceNameProductKeyPrefix, data.DeviceName, data.ProductKey)
	publicThingsIdKey := fmt.Sprintf("%s%v", cachePublicThingsIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)", m.table, thingsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeviceId, data.ProductKey, data.DeviceName, data.DeviceSecret, data.Nickname, data.Status, data.CreatedAt, data.RegisterAt, data.UpdatedAt, data.FirmwareVersion, data.Ip, data.ClientId, data.Mac, data.IsDeleted, data.Extra, data.IsMocker, data.ProductId, data.BatchId, data.EnabledStatus, data.ProjectId)
	}, publicThingsDeviceIdKey, publicThingsDeviceNameProductKeyKey, publicThingsIdKey)
	return ret, err
}

func (m *defaultThingsModel) Update(ctx context.Context, newData *Things) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	publicThingsDeviceIdKey := fmt.Sprintf("%s%v", cachePublicThingsDeviceIdPrefix, data.DeviceId)
	publicThingsDeviceNameProductKeyKey := fmt.Sprintf("%s%v:%v", cachePublicThingsDeviceNameProductKeyPrefix, data.DeviceName, data.ProductKey)
	publicThingsIdKey := fmt.Sprintf("%s%v", cachePublicThingsIdPrefix, data.Id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, thingsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Id, newData.DeviceId, newData.ProductKey, newData.DeviceName, newData.DeviceSecret, newData.Nickname, newData.Status, newData.CreatedAt, newData.RegisterAt, newData.UpdatedAt, newData.FirmwareVersion, newData.Ip, newData.ClientId, newData.Mac, newData.IsDeleted, newData.Extra, newData.IsMocker, newData.ProductId, newData.BatchId, newData.EnabledStatus, newData.ProjectId)
	}, publicThingsDeviceIdKey, publicThingsDeviceNameProductKeyKey, publicThingsIdKey)
	return err
}

func (m *defaultThingsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePublicThingsIdPrefix, primary)
}

func (m *defaultThingsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", thingsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultThingsModel) tableName() string {
	return m.table
}
