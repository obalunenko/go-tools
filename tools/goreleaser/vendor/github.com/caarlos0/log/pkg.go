package log

import (
	"io"
	"os"

	"github.com/charmbracelet/colorprofile"
)

// Log is a singleton instance of [Logger].
var Log Interface = New(os.Stderr)

// New creates a new logger.
func New(w io.Writer) *Logger {
	env := os.Environ()
	if os.Getenv("CI") != "" {
		env = append(env, "CLICOLOR_FORCE=1")
	}
	return &Logger{
		Writer:  colorprofile.NewWriter(w, env),
		Padding: defaultPadding,
		Level:   InfoLevel,
	}
}

// SetLevel sets the log level. This is not thread-safe.
func SetLevel(l Level) {
	if logger, ok := Log.(*Logger); ok {
		logger.Level = l
	}
}

// SetLevelFromString sets the log level from a string, panicing when invalid. This is not thread-safe.
func SetLevelFromString(s string) {
	if logger, ok := Log.(*Logger); ok {
		logger.Level = MustParseLevel(s)
	}
}

// ResetPadding resets the padding to default.
func ResetPadding() {
	Log.ResetPadding()
}

// IncreasePadding increases the padding 1 times.
func IncreasePadding() {
	Log.IncreasePadding()
}

// DecreasePadding decreases the padding 1 times.
func DecreasePadding() {
	Log.DecreasePadding()
}

// WithField returns a new entry with the `key` and `value` set.
func WithField(key string, value any) *Entry {
	return Log.WithField(key, value)
}

// WithError returns a new entry with the "error" set to `err`.
func WithError(err error) *Entry {
	return Log.WithError(err)
}

// WithoutPadding returns a new entry with padding set to default.
func WithoutPadding() *Entry {
	return Log.WithoutPadding()
}

// Debug level message.
func Debug(msg string) {
	Log.Debug(msg)
}

// Info level message.
func Info(msg string) {
	Log.Info(msg)
}

// Warn level message.
func Warn(msg string) {
	Log.Warn(msg)
}

// Error level message.
func Error(msg string) {
	Log.Error(msg)
}

// Fatal level message, followed by an exit.
func Fatal(msg string) {
	Log.Fatal(msg)
}

// Debugf level formatted message.
func Debugf(msg string, v ...any) {
	Log.Debugf(msg, v...)
}

// Infof level formatted message.
func Infof(msg string, v ...any) {
	Log.Infof(msg, v...)
}

// Warnf level formatted message.
func Warnf(msg string, v ...any) {
	Log.Warnf(msg, v...)
}

// Errorf level formatted message.
func Errorf(msg string, v ...any) {
	Log.Errorf(msg, v...)
}

// Fatalf level formatted message, followed by an exit.
func Fatalf(msg string, v ...any) {
	Log.Fatalf(msg, v...)
}
