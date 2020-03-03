package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"goApiFrame/web/common"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

/**
SugaredLogger 与 Logger 对比
在性能很好但不是很关键的上下文中，使用SugaredLogger。它比其他结构化日志记录包快4-10倍，并且支持结构化和printf风格的日志记录。
在每一微秒和每一次内存分配都很重要的上下文中，使用Logger。它甚至比SugaredLogger更快，内存分配次数也更少，但它只支持强类型的结构化日志记录。
打印示例
{"level":"ERROR","time":"2020-03-03T10:37:48.862+0800","caller":"C:/Users/xdf/go/src/goApiFrame/web/main.go:24","msg":"main test err"}
*/
var SugarLogger *zap.SugaredLogger
var timeFormat = "20060102"

func getHook(path string) lumberjack.Logger {
	response_hook := lumberjack.Logger{
		Filename:   path + time.Now().Format(timeFormat) + ".log", // 日志文件路径
		MaxSize:    common.MyConfig.HookMaxSize,                   // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: common.MyConfig.HookMaxBackups,                // 日志文件最多保存多少个备份
		MaxAge:     common.MyConfig.HookMaxAge,                    // 文件最多保存多少天
		Compress:   common.MyConfig.HookCompress,                  // 是否压缩
	}
	return response_hook
}

func getEncoderConfig() zapcore.EncoderConfig {
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
	errHook := getHook("./logs/err/")
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(getEncoderConfig()),                                         // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&errHook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 构造日志
	SugarLogger = zap.New(core, caller, development).Sugar()
}
