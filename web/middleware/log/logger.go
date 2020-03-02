package log

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"goApiFrame/web/common"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var Log *zap.Logger
var response_core zapcore.Core
var err_core zapcore.Core
var caller zap.Option
var development zap.Option

func getResponseHook() lumberjack.Logger {
	response_hook := lumberjack.Logger{
		Filename:   "./logs/response/" + time.Now().Format(common.DateFormat) + ".log", // 日志文件路径
		MaxSize:    common.MyConfig.HookMaxSize,                                        // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: common.MyConfig.HookMaxBackups,                                     // 日志文件最多保存多少个备份
		MaxAge:     common.MyConfig.HookMaxAge,                                         // 文件最多保存多少天
		Compress:   common.MyConfig.HookCompress,                                       // 是否压缩
	}
	return response_hook
}

func getErrorHook() lumberjack.Logger {
	err_hook := lumberjack.Logger{
		Filename:   "./logs/err/" + time.Now().Format(common.DateFormat) + ".log", // 日志文件路径
		MaxSize:    common.MyConfig.HookMaxSize,                                   // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: common.MyConfig.HookMaxBackups,                                // 日志文件最多保存多少个备份
		MaxAge:     common.MyConfig.HookMaxAge,                                    // 文件最多保存多少天
		Compress:   common.MyConfig.HookCompress,                                  // 是否压缩
	}
	return err_hook
}

func getResponseEncoderConfig() zapcore.EncoderConfig {
	Config := zapcore.EncoderConfig{
		TimeKey:  "time",  //输出时间的key名
		LevelKey: "level", //输出日志级别的key名
		NameKey:  "logger",
		//CallerKey:      "caller",
		//MessageKey:     "msg", //输入信息的key名
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,      //每行的分隔符。基本zapcore.DefaultLineEnding 即"\n"
		EncodeLevel:    zapcore.CapitalLevelEncoder,    //使用大写字母记录日志级别
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // 输出的时间格式 ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	return Config
}

func getErrEncoderConfig() zapcore.EncoderConfig {
	Config := zapcore.EncoderConfig{
		TimeKey:        "time",  //输出时间的key名
		LevelKey:       "level", //输出日志级别的key名
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg", //输入信息的key名
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,      //每行的分隔符。基本zapcore.DefaultLineEnding 即"\n"
		EncodeLevel:    zapcore.CapitalLevelEncoder,    //使用大写字母记录日志级别
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // 输出的时间格式 ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	return Config
}

func InitLogger() {
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
	errHook := getErrorHook()
	err_core = zapcore.NewCore(
		zapcore.NewJSONEncoder(getErrEncoderConfig()),                                      // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&errHook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)
	respHook := getResponseHook()
	response_core = zapcore.NewCore(
		zapcore.NewJSONEncoder(getResponseEncoderConfig()),      // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&respHook)), // 打印到文件
		atomicLevel, // 日志级别
	)
	// 开启开发模式，堆栈跟踪
	caller = zap.AddCaller()
	// 开启文件及行号
	development = zap.Development()
	// 构造日志
	Log = zap.New(err_core, caller)
}

//自定义日志中间件,用于记录请求访问日志
func Logger() gin.HandlerFunc {
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
		log := zap.New(response_core, filed)
		log.Info("response_log")
	}
}
