package logger
import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

var zapLog *zap.Logger

func init() {
    var err error
    config := zap.NewProductionConfig()
    encoderConfig := zap.NewProductionEncoderConfig()
    zapcore.TimeEncoderOfLayout("Jan _2 15:04:05.000000000")
    encoderConfig.StacktraceKey = ""
    config.EncoderConfig = encoderConfig

    zapLog, err = config.Build(zap.AddCallerSkip(1))
    if err != nil {
        panic(err)
    }
}

func Debug(message string, fields ...zap.Field) {
    zapLog.Debug(message, fields...)
    defer zapLog.Sync()
}

func Info(message string, fields ...zap.Field) {
    zapLog.Info(message, fields...)
    defer zapLog.Sync()
}

func Warn(message string, fields ...zap.Field) {
    zapLog.Warn(message, fields...)
    defer zapLog.Sync()
}

func Error(message string, fields ...zap.Field) {
    zapLog.Error(message, fields...)
    defer zapLog.Sync()
}

func Panic(message string, fields ...zap.Field) {
    zapLog.Panic(message, fields...)
    defer zapLog.Sync()
}

func Fatal(message string, fields ...zap.Field) {
    zapLog.Fatal(message, fields...)
    defer zapLog.Sync()
}
