package interceptors

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

// ServerInterceptors 服务拦截器
func ServerInterceptors(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Println("服务拦截器开始了：", req)

	return handler(ctx, req)
}
