package main

import (
	"GoLearn/eazy-chat/pkg/interceptor/rpcserver"
	"flag"
	"fmt"

	"GoLearn/eazy-chat/apps/user/rpc/internal/config"
	"GoLearn/eazy-chat/apps/user/rpc/internal/server"
	"GoLearn/eazy-chat/apps/user/rpc/internal/svc"
	"GoLearn/eazy-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/dev/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	// 日志错误记录中间件
	s.AddUnaryInterceptors(rpcserver.LogInterceptor)
	defer s.Stop()
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
