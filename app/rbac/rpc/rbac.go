package main

import (
	"business/app/rbac/rpc/interceptors"
	"business/app/rbac/rpc/internal/config"
	"business/app/rbac/rpc/internal/server"
	"business/app/rbac/rpc/internal/svc"
	"business/app/rbac/rpc/rbac"
	"business/common/vars"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/rbac."+vars.Env+".yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewRbacCenterServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		rbac.RegisterRbacCenterServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	s.AddUnaryInterceptors(interceptors.ServerInterceptors)

	fmt.Printf("Starting client server at %s...\n", c.ListenOn)
	s.Start()
}
