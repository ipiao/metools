package melogger

import "testing"

func TestLogger(t *testing.T) {
	logger := New("Test")
	// logger.SetHeader(simpleHeader)
	logger.Debug("debug")
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
}
