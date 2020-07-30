package grpc

import (
	"context"
	"github.com/fengberlin/go-grpc-http-rest-microservice-tutorial/pkg/api/v1"
	"github.com/fengberlin/go-grpc-http-rest-microservice-tutorial/pkg/logger"
	"github.com/fengberlin/go-grpc-http-rest-microservice-tutorial/pkg/protocol/middleware"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
)

// RunServer运行gRPC服务以发布ToDo服务
func RunServer(ctx context.Context, v1API v1.ToDoServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// gRPC服务器statup选项
	opts := []grpc.ServerOption{}

	// 添加中间件
	opts = middleware.AddLogging(logger.Log, opts)

	// 注册服务
	server := grpc.NewServer(opts...)
	v1.RegisterToDoServiceServer(server, v1API)

	// 优雅地关闭
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// 信号是CTRL+C
			logger.Log.Warn("shutting down gRPC server...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	// 启动gRPC服务器
	logger.Log.Info("starting gRPC server...")

	// 原作者的启动gRPC服务器是这样子的，但我觉得不太好，所以我改为我的方式去启动
	// return server.Serve(listen)
	if err := server.Serve(listen); err != nil {
		logger.Log.Fatal("starting gRPC server failed...")
		return err
	}

	return nil
}