package websocket

import (
	"go.uber.org/zap"
)

var (
	wLog Logger = zap.NewExample().Sugar().With("pkg", "metrics", "internal", "dionysus")
)

type Lwriter struct {
}

func (writer Lwriter) Write(p []byte) (n int, err error) {
	wLog.Errorf(string(p))
	return len(p), nil
}

type Logger interface {
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
}

func SetLog(logger Logger) {
	wLog = logger
	wLog.Infof("websocket log is set")
}
