package logwrapper

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Options struct {
	ServiceName string
	Environment string
}

type StandardLogger struct {
	*logrus.Entry
}

func CreateLogger(opts *Options) (*StandardLogger, error) {
	baseLogger := logrus.New()
	baseLogger.Formatter = &logrus.JSONFormatter{}
	baseLogger.SetReportCaller(true)

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
