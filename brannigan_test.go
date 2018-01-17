package brannigan_test

import (
	"testing"

	"github.com/charithe/brannigan"
	"github.com/coreos/pkg/capnslog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestRedirection(t *testing.T) {
	levelEnabler := func(lvl zapcore.Level) bool { return true }
	zc, obs := observer.New(zap.LevelEnablerFunc(levelEnabler))
	wrapCore := zap.WrapCore(func(zapcore.Core) zapcore.Core {
		return zc
	})

	zapLogger, err := zap.NewDevelopment(wrapCore)
	if err != nil {
		t.Errorf("Error creating Zap logger: %+v", err)
	}

	brannigan.RedirectCapnslog(zapLogger)
	logger := capnslog.NewPackageLogger("repo", "pkg")
	logger.SetLevel(capnslog.TRACE)
	logger.Trace("trace", "entry")
	logger.Debug("debug", "entry")
	logger.Info("info", "entry")
	logger.Notice("notice", "entry")
	logger.Warning("warning", "entry")
	logger.Error("error", "entry")
	logger.Flush()

	if obs.Len() != 6 {
		t.Errorf("Expected 6 log entries. Actual = %d", obs.Len())
	}

	expectedValues := []struct {
		level      zapcore.Level
		message    string
		fieldCount int
	}{
		{
			level:      zapcore.DebugLevel,
			message:    "traceentry",
			fieldCount: 2,
		},
		{
			level:      zapcore.DebugLevel,
			message:    "debugentry",
			fieldCount: 1,
		},
		{
			level:      zapcore.InfoLevel,
			message:    "infoentry",
			fieldCount: 1,
		},
		{
			level:      zapcore.InfoLevel,
			message:    "noticeentry",
			fieldCount: 2,
		},
		{
			level:      zapcore.WarnLevel,
			message:    "warningentry",
			fieldCount: 1,
		},
		{
			level:      zapcore.ErrorLevel,
			message:    "errorentry",
			fieldCount: 1,
		},
	}

	for i, logEntry := range obs.All() {
		expectedVal := expectedValues[i]
		if logEntry.Entry.Level != expectedVal.level {
			t.Errorf("Expected level: %d. Actual: %d", expectedVal.level, logEntry.Entry.Level)
		}

		if logEntry.Entry.Message != expectedVal.message {
			t.Errorf("Expected message: '%s'. Actual: '%s'", expectedVal.message, logEntry.Entry.Message)
		}

		if len(logEntry.Context) != expectedVal.fieldCount {
			t.Errorf("Expected %d context values. Actual: %d", expectedVal.fieldCount, len(logEntry.Context))
		}
	}
}
