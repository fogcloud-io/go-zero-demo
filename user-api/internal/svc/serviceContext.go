package svc

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"zero-demo/user-api/internal/config"
	"zero-demo/user-rpc/usercenter"
)

type ServiceContext struct {
	Config        config.Config
	UserRpcClient usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserRpcClient: usercenter.NewUsercenter(zrpc.MustNewClient(c.UserRpcConf, zrpc.WithUnaryClientInterceptor(MetadataInterceptor),
			zrpc.WithUnaryClientInterceptor(ParseTokenInterceptor))),
	}
}

// ParseTokenInterceptor 客户端拦截器
func ParseTokenInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	logx.Infof("ParseTokenInterceptor start ...")
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		logx.Infof("ParseTokenInterceptor err:%s", err)
		return err
	}
	logx.Infof("ParseTokenInterceptor  end...")
	return nil
}

// MetadataInterceptor API 和RPC 通过 metadata传参
func MetadataInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md := metadata.New(map[string]string{"nickname":"zhangSan","userId":"35"})
	ctx = metadata.NewOutgoingContext(ctx,md)

	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		logx.Infof("MetadataInterceptor err:%s", err)
		return err
	}
	return nil
}
