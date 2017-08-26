package log

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
func LogInit() {
	dyn := zap.NewAtomicLevel()
	dyn.SetLevel(zap.DebugLevel)
	cfg := zap.Config{
		Level:       dyn,
		Development: true,
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "logger",
			CallerKey:      "C",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
		},
		OutputPaths:      []string{"stdout", "./info.log"},
		ErrorOutputPaths: []string{"stderr", "./error.log"},
	}
	zLog, err := cfg.Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	Log = zLog.Sugar()
}
