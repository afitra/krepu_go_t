package logger

import (
	"errors"
	"fmt"
	"runtime"
	"strings"

	"github.com/fatih/structs"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

const (
	timestampFormat = "2006-01-02 15:04:05.000"
	starString      = "**********"
)

var (
	errEchoContextNil = errors.New("Echo context tidak boleh nil")
)

type requestLogger struct {
	RequestID string      `json:"requestID,omitempty"`
	Payload   interface{} `json:"payload,omitempty"`
}

// Init function to make an initial logger
func Init() {
	logrus.SetReportCaller(true)
	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: timestampFormat,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			tmp := strings.Split(f.File, "/")
			filename := tmp[len(tmp)-1]
			return "", fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}

	logrus.SetFormatter(formatter)
	logrus.SetLevel(logrus.DebugLevel)
}

// Make is to get a log parameter
func Make(c echo.Context, payload interface{}) *logrus.Entry {
	var rl requestLogger
	logrus.SetReportCaller(true)

	if c != nil {
		rl.RequestID = c.Response().Header().Get(echo.HeaderXRequestID)
	}

	return logrus.WithFields(logrus.Fields{"params": structs.Map(rl)})
}
