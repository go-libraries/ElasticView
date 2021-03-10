package logs

import (
	"io"
	"log"
	"path/filepath"
	"time"

	"ElasticView/platform-basic-libs/util"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// 初始化日志 logger
func InitLog(logPath string,storageDays int) {
	if logPath == ""{
		logPath = filepath.Join(util.GetCurrentDirectory(),"logs")
	}
	config := zapcore.EncoderConfig{
		MessageKey: "msg",
		TimeKey:    "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:      "file",
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	encoder := zapcore.NewConsoleEncoder(config)


	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel && lvl >= zap.InfoLevel
	})

	infoWriter := getWriter(filepath.Join(logPath, "info.log"),storageDays)
	warnWriter := getWriter(filepath.Join(logPath, "err.log"),storageDays)
	var core zapcore.Core

	core = zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel),
	)

	Logger = zap.New(core, zap.AddCaller(), zap.Development())
}

func getWriter(filename string,storageDays int) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 info.log.YYmmddHH
	hook, err := rotatelogs.New(
		filename+".%Y%m%d", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*time.Duration(storageDays)),     // 保存3天
		rotatelogs.WithRotationTime(time.Hour*24), //切割频率 24小时
	)
	if err != nil {
		log.Panic("日志启动异常", err)
	}
	return hook
}

func Debug(format string, v ...interface{}) {
	log.Println(format, v)
}
