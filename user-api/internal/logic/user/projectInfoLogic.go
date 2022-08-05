package user

import (
	"context"
	"zero-demo/user-rpc/pb"
	"zero-demo/user-rpc/usercenter"

	"zero-demo/user-api/internal/svc"
	"zero-demo/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProjectInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectInfoLogic {
	return &ProjectInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProjectInfoLogic) ProjectInfo(req *types.ProjectInfoReq) (resp *types.ProjectInfoResp, err error) {
	reqInfo := &pb.GetProjectReq{
		ProKey: req.ProKey,
	}
	projectInfo, err := l.svcCtx.UserRpcClient.GetProjectInfo(l.ctx, reqInfo)
	if err != nil {
		logx.Infof("get project info failed:%s", err)
		return nil, err
	}
	resp = new(types.ProjectInfoResp)
	resp = fillProjectInfoResp(projectInfo, resp)
	return resp, nil
}

func fillProjectInfoResp(info *usercenter.ProjectResp, resp *types.ProjectInfoResp) *types.ProjectInfoResp {
	resp.Name = info.Name
	resp.AccessKey = info.AccessKey
	resp.Status = int(info.Status)
	resp.AccessSecret = info.AccessSecret
	resp.Types = info.Type
	resp.CreateAt = info.CreateAt
	return resp
}
