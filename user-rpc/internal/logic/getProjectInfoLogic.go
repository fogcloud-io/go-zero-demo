package logic

import (
	"context"
	"time"
	"zero-demo/user-rpc/models"

	"zero-demo/user-rpc/internal/svc"
	"zero-demo/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProjectInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectInfoLogic {
	return &GetProjectInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//func (l *GetProjectInfoLogic) GetProjectInfoByProName(in *pb.GetProjectReq) (*pb.ProjectResp, error) {
//	project, err := l.svcCtx.ProjectsModel.FindOne(l.ctx, 1)
//	if err != nil {
//		logx.Infof("[GetProjectInfo] err:%s", err)
//		return nil, err
//	}
//	resp := fillProjectResp(project)
//	return resp, nil
//}

func (l *GetProjectInfoLogic) GetProjectInfo(in *pb.GetProjectReq) (*pb.ProjectResp, error) {
	//whereBuild := l.svcCtx.ProjectsModel.RowBuilder().Where(squirrel.Eq{"name":in.ProKey})
	resp, err := l.svcCtx.ProjectsModel.FindOneByQuery(l.ctx, in.ProKey)
	logx.Infof("resp:%v err:%s", resp, err)
	if err == nil {
		return fillProjectResp(resp),nil
	}
	return nil, err
}

func fillProjectResp(project *models.Projects) *pb.ProjectResp {
	proResp := new(pb.ProjectResp)
	proResp.Status = uint32(project.Status)
	proResp.AccessKey = project.AccessKey
	proResp.AccessSecret = project.AccessSecret
	proResp.Type = project.Type
	proResp.CreateAt = project.CreatedAt.Format(time.RFC3339)
	proResp.Name = project.Name
	return proResp
}
