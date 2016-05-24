package igprotocol

import (
    "fmt"
    "os"
)

// StdOutLogging ...
type StdOutLogging struct {
	LogEnabled bool
}

// GetLogEnabled ...
func (ctx StdOutLogging) GetLogEnabled() bool {
	return ctx.LogEnabled
}

// Logf ...
func (ctx *StdOutLogging) Logf(format string, args ...interface{}) {
	if ctx.LogEnabled {
		s := fmt.Sprintf(format, args)
		os.Stdout.WriteString(s)

	}
}

// Log ...
func (ctx *StdOutLogging) Log(s string) {
	if ctx.LogEnabled {
		os.Stdout.WriteString(s)

	}
}
