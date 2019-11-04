package common

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var Log *zap.Logger

//lumberjack切割日志，不足之处，在于日志切割后日志，其文件名过于复杂，可能不便后后继处理
//日志名示例：system-2019-11-04T06-21-46.400
func InitLogger() {
	hook := lumberjack.Logger{
		Filename:   MyConfig.HookFilename,   // 日志文件路径
		MaxSize:    MyConfig.HookMaxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: MyConfig.HookMaxBackups, // 日志文件最多保存多少个备份
		MaxAge:     MyConfig.HookMaxAge,     // 文件最多保存多少天
		Compress:   MyConfig.HookCompress,   // 是否压缩
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
	switch MyConfig.Loglevel {
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
	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", MyConfig.ServiceName))
	// 构造日志
	Log = zap.New(core, caller, development, filed)
	Log.Info("DefaultLogger init success")
}
