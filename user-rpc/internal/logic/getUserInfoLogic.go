package logic

import (
	"context"
	"google.golang.org/grpc/metadata"
	"strconv"
	"zero-demo/user-rpc/internal/svc"
	"zero-demo/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	logx.Infof("GetUserInfo start...")

	// meta 传递参数
	var userId string
	if md, ok := metadata.FromIncomingContext(l.ctx); ok {
		logx.Infof("md:%+v", md)
		name := md.Get("nickname")
		userId = md.Get("userId")[0]
		logx.Infof("nickname:%s userId:%s", name, userId)
	}

	tokenUserId := l.ctx.Value("user_id")
	logx.Infof("tokenUserId:%v",tokenUserId)

	uid, _ := strconv.ParseInt(userId, 10, 64)

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, uid)
	logx.Infof("user:%v err:%s", user, err)
	if user == nil {
		return &pb.GetUserInfoResp{Id: in.Id, Nickname: ""}, nil
	}

	return &pb.GetUserInfoResp{Id: in.Id, Nickname: user.Nickname.String}, nil
}
