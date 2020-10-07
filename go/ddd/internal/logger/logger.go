package logger

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type (
	// StandardLog is wrap log data
	StandardLog struct {
		*logrus.Entry
	}
)

// NewLogRequest  is create log request
func NewLogRequest(param *http.Request) *logrus.Entry {
	var log logrus.Logger
	log.SetFormatter(&logrus.JSONFormatter{})

	return log.WithFields(
		logrus.Fields{
			"start_at": time.Now().Format("2006-01-02 15:04:05"),
			"method":   param.Method,
			"uri":      param.RequestURI,
			"ip":       param.Host,
		},
	)
}

// NewLogger is function to create new log
func NewLogger(appname string) *StandardLog {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	var baselog = logrus.New().WithFields(logrus.Fields{
		"syslog": map[string]string{
			"hostname": hostname,
			"appname":  appname,
		},
	})

	var standardLog = &StandardLog{baselog}
	standardLog.Logger.Formatter = &logrus.JSONFormatter{}
	standardLog.Entry = standardLog.WithFields(
		logrus.Fields{
			"start_at": time.Now().Format("2006-01-02 15:04:05"),
		},
	)
	return standardLog
}

// DataLog is
type DataLog struct {
	Request      interface{} `json:"-"`
	RctraceID    string      `json:"rctrace_id"`
	Error        string      `json:"error"`
	File         string      `json:"-"`
	FileFunction string      `json:"-"`
	Event        string      `json:"-"`
}

// DataLog is when add data to logger
func (l *StandardLog) DataLog(d DataLog) *logrus.Entry {
	data, _ := json.Marshal(d.Request)
	return l.WithFields(logrus.Fields{
		"errors": map[string]interface{}{
			"path": map[string]string{
				"file":     d.File,
				"function": d.FileFunction,
			},
			"event":      d.Event,
			"rctrace_id": d.RctraceID,
			"error":      d.Error,
		},
		"data": fmt.Sprintf("%+v", string(data)),
	})
}

// Fatal Any error that is forcing a shutdown of the service or application to prevent data loss (or further data loss)
func (l *StandardLog) Fatal(msg string, data DataLog) {
	log := l.DataLog(data)
	log.Fatal(msg)
}

// Error Any error which is fatal to the operation, but not the service or application (can't open a required file. missing data, etc.).These errors will force user (administrator, or direct user) intervention.
func (l *StandardLog) Error(msg string, data DataLog) {
	log := l.DataLog(data)
	log.Error(msg)
}

// Info Generally useful information to log (service start/stop, configuration assumptions, etc). Info I want to always have available but usually don't care about under normal circumstances
func (l *StandardLog) Info(msg string, data DataLog) {
	log := l.DataLog(data)
	log.Info(msg)
}

// Debug  Information that is diagnostically helpful to people more than just developers (IT, sysadmins, etc.).
func (l *StandardLog) Debug(msg string, data DataLog) {
	log := l.DataLog(data)
	log.Debug(msg)
}

// Trace  Only when I would be "tracing" the code and trying to find one part of a function specifically
func (l *StandardLog) Trace(msg string, data DataLog) {
	log := l.DataLog(data)
	log.Trace(msg)
}
