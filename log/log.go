package logwrapper

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Options struct {
	ServiceName string
	Environment string
	LogPath		string
}

type StandardLogger struct {
	*logrus.Entry
}

func CreateLogger(opts *Options) (*StandardLogger, error) {
	logFile, err := os.OpenFile(opts.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	baseLogger := logrus.New()
	baseLogger.Formatter = &logrus.JSONFormatter{}
	baseLogger.SetReportCaller(true)
	baseLogger.SetOutput(logFile)

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	loggerWithStdFields := baseLogger.WithFields(logrus.Fields{
		"service_name": opts.ServiceName,
		"environment":  opts.Environment,
		"hostname":     hostname,
	})

	standardLogger := &StandardLogger{loggerWithStdFields}

	return standardLogger, nil
}
