package logic

import (
	"context"

	"zero-demo/user-rpc/internal/svc"
	"zero-demo/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectInfoByProKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProjectInfoByProKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectInfoByProKeyLogic {
	return &GetProjectInfoByProKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProjectInfoByProKeyLogic) GetProjectInfoByProKey(in *pb.GetProjectReq) (*pb.ProjectResp, error) {
	project, err := l.svcCtx.ProjectsModel.FindOne(l.ctx, 1)
	if err != nil {
		logx.Infof("[GetProjectInfo] err:%s", err)
		return nil, err
	}
	resp := fillProjectResp(project)

	return resp, nil
}
