package logging

import (
	"TVHelper/global"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type logItem struct {
	FileName string
	Level    zap.LevelEnablerFunc
}

// Init 初始化日志.
func Init() {
	logPath := global.LogSetting.Path
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		err := os.Mkdir(logPath, 0666)
		if err != nil {
			log.Fatal("目录已存在或无法创建此目录", logPath)
		}
	}

	items := []logItem{
		{
			FileName: getLogFileName("info"),
			Level: func(level zapcore.Level) bool {
				return level <= zap.InfoLevel
			},
		},
		{
			FileName: getLogFileName("err"),
			Level: func(level zapcore.Level) bool {
				return level > zap.InfoLevel
			},
		},
	}

	NewLogger(items)
}

// NewLogger 日志.
func NewLogger(items []logItem) {
	var (
		cfg   zapcore.Encoder
		cores []zapcore.Core
	)
	if global.LogSetting.Encoder == "json" {
		cfg = zapcore.NewJSONEncoder(getEncoderConfig())
	} else {
		cfg = zapcore.NewConsoleEncoder(getEncoderConfig())
	}

	for _, v := range items {
		hook := lumberjack.Logger{
			Filename:   v.FileName,
			MaxSize:    global.LogSetting.LumberJack.MaxSize,
			MaxBackups: global.LogSetting.LumberJack.MaxBackups,
			MaxAge:     global.LogSetting.LumberJack.MaxAge,
			Compress:   global.LogSetting.LumberJack.Compress,
			LocalTime:  true, // 备份文件名本地/UTC时间
		}
		var writeSyncer zapcore.WriteSyncer
		switch global.LogSetting.Output {
		case "both":
			writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))
		case "file":
			writeSyncer = zapcore.AddSync(&hook)
		case "console":
			writeSyncer = zapcore.AddSync(os.Stdout)
		default:
			writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))
		}
		core := zapcore.NewCore(
			cfg,         // 编码器配置;
			writeSyncer, // 打印到控制台和文件
			v.Level,     // 日志级别
		)
		cores = append(cores, core)
	}

	logger := zap.New(zapcore.NewTee(cores...), zap.AddCaller())
	global.Logger = logger
}

func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zap.NewProductionEncoderConfig()
	// 时间格式自定义
	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("01-02 15:04:05.000"))
	}
	// 打印路径自定义
	config.EncodeCaller = zapcore.ShortCallerEncoder
	return config
}

func getLogFileName(suffix string) string {
	fileName := strings.Join([]string{global.LogSetting.FilePrefix, suffix, "log"}, ".")
	return path.Join(global.LogSetting.Path, fileName)
}
