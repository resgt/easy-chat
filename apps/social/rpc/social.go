package main

import (
	"GoLearn/eazy-chat/pkg/interceptor/rpcserver"
	"flag"
	"fmt"

	"GoLearn/eazy-chat/apps/social/rpc/internal/config"
	"GoLearn/eazy-chat/apps/social/rpc/internal/server"
	"GoLearn/eazy-chat/apps/social/rpc/internal/svc"
	"GoLearn/eazy-chat/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/dev/social.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		social.RegisterSocialServer(grpcServer, server.NewSocialServer(ctx))

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
