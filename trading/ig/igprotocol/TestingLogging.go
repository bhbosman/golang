package igprotocol

import "testing"

// TestingLogging ...
type TestingLogging struct {
	LogEnabled bool
	TestUnit   *testing.T
}

// GetLogEnabled ...
func (ctx *TestingLogging) GetLogEnabled() bool {
	return ctx.LogEnabled
}

// Logf ...
func (ctx *TestingLogging) Logf(format string, args ...interface{}) {
	if ctx.LogEnabled {
		ctx.TestUnit.Logf(format, args)
	}
}

// Log ...
func (ctx *TestingLogging) Log(s string) {
	if ctx.LogEnabled {
		ctx.TestUnit.Log(s)

	}
}
