package user

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-demo/user-api/internal/logic/user"
	"zero-demo/user-api/internal/svc"
	"zero-demo/user-api/internal/types"
)

func ProjectInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProjectInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewProjectInfoLogic(r.Context(), svcCtx)
		resp, err := l.ProjectInfo(&req)
		logx.Infof("project:%v err:%s", resp, err)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
