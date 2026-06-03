package logger

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// Log is the global logger instance
var Log *logrus.Logger

// Init initializes the global logger with the given verbosity
func Init(verbose bool) {
	Log = logrus.New()
	Log.SetOutput(os.Stderr)

	// Set formatter
	if verbose {
		Log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			ForceColors:     true,
		})
		Log.SetLevel(logrus.DebugLevel)
	} else {
		Log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   false,
			TimestampFormat: "",
			ForceColors:     true,
		})
		Log.SetLevel(logrus.InfoLevel)
	}
}

// SetLevel sets the log level from a string
func SetLevel(level string) {
	if Log == nil {
		Init(false)
	}

	switch strings.ToLower(level) {
	case "debug":
		Log.SetLevel(logrus.DebugLevel)
	case "info":
		Log.SetLevel(logrus.InfoLevel)
	case "warn", "warning":
		Log.SetLevel(logrus.WarnLevel)
	case "error":
		Log.SetLevel(logrus.ErrorLevel)
	default:
		Log.SetLevel(logrus.InfoLevel)
	}
}
