// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "zero-demo/user-api/internal/handler/user"
	"zero-demo/user-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/info",
				Handler: user.UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/update",
				Handler: user.UserUpdateHandler(serverCtx),
			},
		},
		rest.WithPrefix("/userapi/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/info",
				Handler: user.UserInfoAuthHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/project/info",
				Handler: user.ProjectInfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/userapi/v1/auth"),
	)
}
