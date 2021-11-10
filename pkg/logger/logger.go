package logger

import (
	"fmt"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type HTTPLog struct {
	IP,
	Method,
	Path,
	RequestBody,
	Query,
	ResponseBody string
	Status int
}

var (
	zapLog    *zap.Logger
	AppLogger *Logger
)

type Logger struct {
	logger *zap.SugaredLogger
}

func Init() {
	initZap()
	AppLogger = &Logger{logger: zapLog.Sugar()}
}

func initZap() {
	var err error
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapLog, err = loggerConfig.Build()
	if err != nil {
		log.Fatalln("Got error while building zap config.")
		return
	}
	return
}

func (l *Logger) HTTPLog(info HTTPLog) {
	l.logger.Infow(fmt.Sprintf("request method %s path %s", info.Method, info.Path),
		"IP", info.IP,
		"method", info.Method,
		"path", info.Path,
		"requestBody", info.RequestBody,
		"query", info.Query,
		"responseBody", info.ResponseBody,
	)
	defer l.logger.Sync()
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.logger.Infof(format, v...)
	defer l.logger.Sync()
}

func (l *Logger) Info(msg string) {
	l.logger.Info(msg)
	defer l.logger.Sync()
}

func (l *Logger) Error(v ...interface{}) {
	l.logger.Error(v)
	defer l.logger.Sync()
}
