package log

import (
	"bytes"
	"fmt"
	"go.uber.org/zap"
	"goa.design/goa/v3/middleware"
	"strings"
)

// adapter is a thin wrapper around the stdlib logger that adapts it to
// the Logger interface.
type adapter struct {
	*zap.Logger
}

// NewLogger creates a Logger backed by a stdlib logger.
func NewLogger(l *zap.Logger) middleware.Logger {
	return &adapter{l}
}

func (a *adapter) Log(keyvals ...interface{}) error {
	n := (len(keyvals) + 1) / 2
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, "MISSING")
	}
	var fm bytes.Buffer
	vals := make([]interface{}, n)
	for i := 0; i < len(keyvals); i += 2 {
		k := keyvals[i]
		v := keyvals[i+1]
		vals[i/2] = v
		fm.WriteString(fmt.Sprintf(" %s=%%+v", k))
	}
	a.Logger.Info(fmt.Sprintf(strings.TrimSpace(fm.String()), vals...))
	return nil
}
