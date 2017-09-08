// Copyright 2012-2016 Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

import (
	"context"
	"log"
)

// Logger specifies the interface for all log operations.
type Logger interface {
	Printf(ctx context.Context, format string, v ...interface{})
}

type defaultLogger struct {
	logger *log.Logger
}

func newDefaultLogger(logger *log.Logger) Logger {
	return &defaultLogger{
		logger: logger,
	}
}

func (l *defaultLogger) Printf(ctx context.Context, format string, v ...interface{}) {
	l.logger.Printf(format, v ...)
}
