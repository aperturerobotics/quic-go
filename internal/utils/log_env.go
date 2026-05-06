//go:build !tinygo

package utils

import (
	"fmt"
	"os"
	"strings"
)

func readLoggingEnv() LogLevel {
	switch strings.ToLower(os.Getenv(logEnv)) {
	case "":
		return LogLevelNothing
	case "debug":
		return LogLevelDebug
	case "info":
		return LogLevelInfo
	case "error":
		return LogLevelError
	default:
		fmt.Fprintln(os.Stderr, "invalid quic-go log level, see https://github.com/quic-go/quic-go/wiki/Logging")
		return LogLevelNothing
	}
}
