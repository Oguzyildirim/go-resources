package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
	"time"
)

var (
	// Log 是全局日志记录器
	Log *zap.Logger
	// customTimeFormat 是自定义时间格式
	customTimeFormat string
	// onceInit保证仅初始化全局日志记录器一次
	onceInit sync.Once
)

// customTimeEncoder 编码自定义格式的时间
// 这个例子教我们如何定制zap默认功能
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(customTimeFormat))
}

// Init 通过输入参数初始化日志
// lvl - 全局日志级别：Debug（-1），Info（0），Warn（1），Error（2），DPanic（3），Panic（4），Fatal（5）
// timeFormat - 默认使用空字符串的记录器的自定义时间格式
func Init(lvl int, timeFormat string) {

	onceInit.Do(func() {
		// 首先，定义我们的级别处理逻辑。
		globalLevel := zapcore.Level(lvl)

		// 高优先级输出是标准错误，低优先级输出是标准输出。
		// 它对Kubernetes部署很有用。
		// Kubernetes默认将os.Stdout日志项解释为INFO和os.Stderr日志项解释为ERROR。
		highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.ErrorLevel
		})

		lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= globalLevel && lvl < zapcore.ErrorLevel
		})

		consoleInfos := zapcore.Lock(os.Stdout)
		consoleErrors := zapcore.Lock(os.Stderr)

		// 配置控制台输出。
		var useCustomTimeFormat bool
		ecfg := zap.NewProductionEncoderConfig()
		if len(timeFormat) > 0 {
			customTimeFormat = timeFormat
			ecfg.EncodeTime = customTimeEncoder
			useCustomTimeFormat = true
		}
		consoleEncoder := zapcore.NewJSONEncoder(ecfg)

		// 将输出，编码器和日志级别处理功能加入zapcore
		core := zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
			zapcore.NewCore(consoleEncoder, consoleInfos, lowPriority))

		Log = zap.New(core)

		// RedirectStdLog将标准库的包全局记录器的输出重定向到InfoLevel上提供的记录器。
		// 由于zap已经处理了调用者注释，时间戳等，它会自动禁用标准库的注释和前缀。
		zap.RedirectStdLog(Log)

		if !useCustomTimeFormat {
			Log.Warn("time format for logger is not provided - use zap default")
		}
	})
}
