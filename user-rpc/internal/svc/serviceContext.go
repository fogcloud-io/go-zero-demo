package svc

import (
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-demo/user-rpc/internal/config"
	"zero-demo/user-rpc/models"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     models.UserModel
	ProjectsModel models.ProjectsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserModel:     models.NewUserModel(sqlx.NewSqlConn("postgres", c.DB.DataSource)),
		ProjectsModel: models.NewProjectsModel(sqlx.NewSqlConn("postgres", c.DB.DataSource), c.Cache),
	}
}
