package logger

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/labstack/echo"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type Log struct {
	Logger *logrus.Logger
}

func NewLogger(logpath string, logname string, loglevel string) (*Log, error) {
	l := logrus.New()

	if strings.Trim(logname, " ") == "" {
		logname = "app"
	}

	if strings.Trim(logpath, " ") == "" {
		logpath = "/var/log"
	}

	logf, err := rotatelogs.New(
		fmt.Sprintf("%s/%s.%%Y%%m%%d", logpath, logname),

		// symlink current log to this file
		rotatelogs.WithLinkName(fmt.Sprintf("%s/%s.log", logpath, logname)),

		// max : 7 days to keep
		rotatelogs.WithMaxAge(24*7*time.Hour),

		// rotate every day
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return nil, err
	}

	l.Formatter = &logrus.JSONFormatter{}
	l.Out = logf

	switch loglevel {
	case "info":
		l.Level = logrus.InfoLevel
	case "panic":
		l.Level = logrus.PanicLevel
	case "fatal":
		l.Level = logrus.FatalLevel
	case "error":
		l.Level = logrus.ErrorLevel
	default:
		l.Level = logrus.DebugLevel
	}

	return &Log{
		Logger: l,
	}, nil
}

func (l *Log) LogRequest(c echo.Context, req interface{}, resp interface{}) {
	l.Logger.WithFields(logrus.Fields{
		"remote_ip": c.RealIP(),
		"protocol":  c.Request().Proto,
		"host":      c.Request().Host,
		"uri":       c.Request().RequestURI,
		"headers":   c.Request().Header,
		"method":    c.Request().Method,
		"request":   req,
		"response":  resp,
	}).Info("request log")
}
