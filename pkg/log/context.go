package log

import (
	"context"
)

var (
	// G is an alias for GetLogger.
	//
	// We may want to define this locally to a package to get package tagged log
	// messages.
	G = GetLogger

	// L is an alias for the standard logger.
	L FieldLogger
)

type loggerKey struct{}

// WithLogger returns a new context with the provided logger. Use in
// combination with logger.WithField(s) for great effect.
func WithLogger(ctx context.Context, logger FieldLogger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// GetLogger retrieves the current logger from the context. If no logger is
// available, the default logger is returned.
func GetLogger(ctx context.Context) FieldLogger {
	logger := ctx.Value(loggerKey{})
	if logger == nil {
		if L == nil {
			L = NewGoLogger()
		}

		return L
	}

	return logger.(FieldLogger)
}
