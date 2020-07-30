package middleware

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// codeToLevel 将OK重定向到DEBUG级别日志记录而不是INFO
// 这是如何log几个gRPC code结果的示例
func codeToLevel(code codes.Code) zapcore.Level {

	if code == codes.OK {
		// 它是Debug级别
		return zapcore.DebugLevel
	}

	return grpc_zap.DefaultCodeToLevel(code)
}

// AddLogging 返回打开日志记录的grpc.Server配置选项。
func AddLogging(logger *zap.Logger, opts []grpc.ServerOption) []grpc.ServerOption {
	// 日志记录器的共享选项，具有自定义gRPC code以记录级别功能。
	o := []grpc_zap.Option{
		grpc_zap.WithLevels(codeToLevel),
	}

	// 确保使用zapLogger记录gRPC库内部的日志语句。
	grpc_zap.ReplaceGrpcLogger(logger)

	// 添加一元拦截器
	opts = append(opts, grpc_middleware.WithUnaryServerChain(
		grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		grpc_zap.UnaryServerInterceptor(logger, o...),
	))

	// 添加流拦截器（此处作为示例添加）
	opts = append(opts, grpc_middleware.WithStreamServerChain(
		grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		grpc_zap.StreamServerInterceptor(logger, o...),
	))

	return opts
}


