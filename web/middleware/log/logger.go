package log

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"goApiFrame/web/common"
	"goApiFrame/web/util"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var Log *zap.Logger

//自定义日志中间件
func Logger() gin.HandlerFunc {
	hook := lumberjack.Logger{
		Filename:   "./logs/" + time.Now().Format(util.DateFormat) + ".log", // 日志文件路径
		MaxSize:    common.MyConfig.HookMaxSize,                             // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: common.MyConfig.HookMaxBackups,                          // 日志文件最多保存多少个备份
		MaxAge:     common.MyConfig.HookMaxAge,                              // 文件最多保存多少天
		Compress:   common.MyConfig.HookCompress,                            // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",  //输出时间的key名
		LevelKey:       "level", //输出日志级别的key名
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg", //输入信息的key名
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,      //每行的分隔符。基本zapcore.DefaultLineEnding 即"\n"
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // 输出的时间格式 ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别,debug可以打印出info,debug,warn；info级别可以打印warn，info；warn只能打印warn
	var level zapcore.Level
	switch common.MyConfig.Loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	atomicLevel := zap.NewAtomicLevelAt(level)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()

	return func(context *gin.Context) {
		startTime := time.Now()
		context.Next()
		endTime := time.Now()
		//执行时间
		useTime := endTime.Sub(startTime)
		requestMethod := context.Request.Method
		requestUrl := context.Request.RequestURI
		statusCode := context.Writer.Status()
		clientIP := context.ClientIP()
		// 设置初始化字段
		filed := zap.Fields(
			zap.String("service_name", common.MyConfig.ServiceName),
			zap.Any("status_code", statusCode),
			zap.Duration("cost_time", useTime),
			zap.String("request_method", requestMethod),
			zap.String("request_url", requestUrl),
			zap.String("client_ip", clientIP),
		)
		// 构造日志
		Log = zap.New(core, caller, development, filed)
		Log.Info("")
	}
}
