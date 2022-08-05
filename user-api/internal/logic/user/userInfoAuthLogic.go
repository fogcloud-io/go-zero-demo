package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/stringx"
	"zero-demo/user-rpc/pb"

	"zero-demo/user-api/internal/svc"
	"zero-demo/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoAuthLogic {
	return &UserInfoAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoAuthLogic) UserInfoAuth(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	logx.Infof("UserInfo req:%v rand:%s", req,stringx.Rand())
	infoReq := &pb.GetUserInfoReq{Id: req.UserId}

	userInfo, err := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, infoReq)

	if err != nil {
		return nil, err
	}

	return &types.UserInfoResp{
		UserId:   userInfo.Id,
		Nickname: userInfo.Nickname,
	}, nil
}
