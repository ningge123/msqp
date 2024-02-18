package logs

import (
	"github.com/charmbracelet/log"
	"msqp/common/config"
	"os"
	"time"
)

var logger *log.Logger

func Init(appName string) {
	logger = log.New(os.Stderr)

	logger.SetLevel(log.InfoLevel)
	if config.Conf.Log.Level == "DEBUG" {
		logger.SetLevel(log.DebugLevel)
	}

	log.SetPrefix(appName)
	log.SetReportTimestamp(true)
	log.SetTimeFormat(time.DateTime)
}

func Fatal(format string, values ...any) {
	if len(values) == 0 {
		logger.Fatal(format)
	} else {
		logger.Fatalf(format, values)
	}
}

func Info(format string, values ...any) {
	if len(values) == 0 {
		logger.Info(format)
	} else {
		logger.Infof(format, values)
	}
}

func Error(format string, values ...any) {
	if len(values) == 0 {
		logger.Error(format)
	} else {
		logger.Errorf(format, values)
	}
}

func Debug(format string, values ...any) {
	if len(values) == 0 {
		logger.Debug(format)
	} else {
		logger.Debugf(format, values)
	}
}
