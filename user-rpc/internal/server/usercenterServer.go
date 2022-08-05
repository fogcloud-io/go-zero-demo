// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"zero-demo/user-rpc/internal/logic"
	"zero-demo/user-rpc/internal/svc"
	"zero-demo/user-rpc/pb"
)

type UsercenterServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUsercenterServer
}

func NewUsercenterServer(svcCtx *svc.ServiceContext) *UsercenterServer {
	return &UsercenterServer{
		svcCtx: svcCtx,
	}
}

func (s *UsercenterServer) GetUserInfo(ctx context.Context, in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

func (s *UsercenterServer) GetProjectInfo(ctx context.Context, in *pb.GetProjectReq) (*pb.ProjectResp, error) {
	l := logic.NewGetProjectInfoLogic(ctx, s.svcCtx)
	return l.GetProjectInfo(in)
}

func (s *UsercenterServer) GetProjectInfoByProKey(ctx context.Context, in *pb.GetProjectReq) (*pb.ProjectResp, error) {
	l := logic.NewGetProjectInfoByProKeyLogic(ctx, s.svcCtx)
	return l.GetProjectInfoByProKey(in)
}
