package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-demo/user-api/internal/logic/user"
	"zero-demo/user-api/internal/svc"
	"zero-demo/user-api/internal/types"

)

func UserInfoAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserInfoAuthLogic(r.Context(), svcCtx)
		resp, err := l.UserInfoAuth(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
