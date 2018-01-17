// Package brannigan contains helper methods to redirect CoreOS capnslog output to an Uber Zap logger
package brannigan

import (
	"fmt"

	"github.com/coreos/pkg/capnslog"
	"go.uber.org/zap"
)

// ZapFormatter implements the capnslog.Formatter interface
type ZapFormatter struct {
	logger *zap.Logger
}

func (z *ZapFormatter) Format(pkg string, level capnslog.LogLevel, depth int, entries ...interface{}) {
	msg := fmt.Sprint(entries...)
	field := zap.String("pkg", pkg)

	switch level {
	case capnslog.CRITICAL:
		z.logger.Fatal(msg, field)
	case capnslog.ERROR:
		z.logger.Error(msg, field)
	case capnslog.WARNING:
		z.logger.Warn(msg, field)
	case capnslog.NOTICE:
		z.logger.Info(msg, field, zap.Bool("notice", true))
	case capnslog.INFO:
		z.logger.Info(msg, field)
	case capnslog.DEBUG:
		z.logger.Debug(msg, field)
	case capnslog.TRACE:
		z.logger.Debug(msg, field, zap.Bool("trace", true))
	}
}

func (z *ZapFormatter) Flush() {
	z.logger.Sync()
}

// NewZapFormatter creates a capnslog Formatter that delegates to the provided zap logger
func NewZapFormatter(logger *zap.Logger) capnslog.Formatter {
	return &ZapFormatter{logger: logger}
}

// RedirectCapnslog sets the capnslog formatter to the provided zap logger
func RedirectCapnslog(logger *zap.Logger) {
	capnslog.SetFormatter(NewZapFormatter(logger))
}

// RedirectCapnslogToGlobalZapLogger sets the capnslog formatter to the global zap logger
func RedirectCapnslogToGlobalZapLogger() {
	capnslog.SetFormatter(NewZapFormatter(zap.L()))
}
