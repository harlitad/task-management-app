package logger

import (
	"os"

	"github.com/harlitad/task-management-app/config"
	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	log *logrus.Entry
}

// var levels map[string]logrus.Level
var levels = map[string]logrus.Level{
	"info":    logrus.InfoLevel,
	"warning": logrus.WarnLevel,
	"error":   logrus.ErrorLevel,
	"debug":   logrus.DebugLevel,
}

func NewLogrusLogger(conf config.Config) ILogger {
	log := logrus.New()
	log.SetLevel(levels[conf.LogLevel])
	log.Formatter = &logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	}
	log.SetOutput(os.Stdout)
	return &LogrusLogger{log: logrus.NewEntry(log)}
}

func (log *LogrusLogger) Info(args ...interface{}) {
	log.log.Info(args...)
}

func (log *LogrusLogger) Debug(args ...interface{}) {
	log.log.Debug(args...)
}

func (log *LogrusLogger) Warn(args ...interface{}) {
	log.log.Warn(args...)
}

func (log *LogrusLogger) Error(args ...interface{}) {
	log.log.Error(args...)
}

func (log *LogrusLogger) Fatal(args ...interface{}) {
	log.log.Fatal(args...)
}

func (log *LogrusLogger) Panic(args ...interface{}) {
	log.log.Panic(args...)
}

func (log *LogrusLogger) Fatalf(template string, args ...interface{}) {
	log.log.Fatalf(template, args...)
}

func (log *LogrusLogger) Panicf(template string, args ...interface{}) {
	log.log.Panicf(template, args...)
}

func (log *LogrusLogger) Infof(template string, args ...interface{}) {
	log.log.Infof(template, args...)
}

func (log *LogrusLogger) Debugf(template string, args ...interface{}) {
	log.log.Debugf(template, args...)
}

func (log *LogrusLogger) Warnf(template string, args ...interface{}) {
	log.log.Warnf(template, args...)
}

func (log *LogrusLogger) Errorf(template string, args ...interface{}) {
	log.log.Errorf(template, args...)
}

func (l *LogrusLogger) WithFields(fields Fields) ILogger {
	newLog := l.log.WithFields(logrus.Fields(fields))
	return &LogrusLogger{log: newLog}
}
