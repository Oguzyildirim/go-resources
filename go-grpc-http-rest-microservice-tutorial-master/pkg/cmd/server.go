package cmd

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/fengberlin/go-grpc-http-rest-microservice-tutorial/pkg/logger"
	"github.com/fengberlin/go-grpc-http-rest-microservice-tutorial/pkg/protocol/grpc"
	"github.com/fengberlin/go-grpc-http-rest-microservice-tutorial/pkg/protocol/rest"
	"github.com/fengberlin/go-grpc-http-rest-microservice-tutorial/pkg/service/v1"

	// mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

// Config是Server的配置
type Config struct {
	// gRPC服务器启动参数部分
	// GRPCPort是gRPC服务器监听的TCP端口
	GRPCPort string

	// HTTP/REST网关启动参数部分
	// HTTPPort是通过HTTP/REST网关监听的TCP端口
	HTTPPort string

	// 数据库数据存储参数部分
	// DatestoreDBHost是数据库的地址
	DatastoreDBHost string
	// DatastoreDBUser是用于连接数据库的用户名
	DatastoreDBUser string
	// DatastoreDBPassword是用于连接数据库的密码
	DatastoreDBPassword string
	// DatastoreDBSchema是数据库的名称
	DatastoreDBSchema string

	// 日志参数部分
	// LogLevel 是全局日志级别：Debug(-1)，Info(0)，Warn(1)，Error(2)，DPanic(3)，Panic(4)，Fatal(5)
	LogLevel int
	// LogTimeFormat 是记录器的打印时间格式，例如2006-01-02T15:04:05Z07:00
	LogTimeFormat string
}

// RunServer运行gRPC服务器和HTTP网关
func RunServer() error {
	ctx := context.Background()

	// 获取配置
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.HTTPPort, "http-port", "", "HTTP port to bind")
	flag.StringVar(&cfg.DatastoreDBHost, "db-host", "", "Database host")
	flag.StringVar(&cfg.DatastoreDBUser, "db-user", "", "Database user")
	flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "", "Database password")
	flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "", "Database schema")
	flag.IntVar(&cfg.LogLevel, "log-level", 0, "Global log level")
	flag.StringVar(&cfg.LogTimeFormat, "log-time-format", "",
		"Print time format for logger e.g. 2006-01-02T15:04:05Z07:00")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	if len(cfg.HTTPPort) == 0 {
		return fmt.Errorf("invalid TCP port for HTTP gateway: '%s'", cfg.HTTPPort)
	}

	// 初始化全局日志记录器
	logger.Init(cfg.LogLevel, cfg.LogTimeFormat)

	// 添加MySQL驱动程序特定参数来解析 date/time
	// 为另一个数据库删除它
	param := "parseTime=true"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", cfg.DatastoreDBUser,
		cfg.DatastoreDBPassword, cfg.DatastoreDBHost, cfg.DatastoreDBSchema, param)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	v1API := v1.NewToDoServiceServer(db)

	// 运行HTTP网关
	go func() {
		_ = rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
