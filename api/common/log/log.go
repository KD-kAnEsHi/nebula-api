// Package log provides logging utilities using the zerolog library.
// It includes functions to log errors, debug messages, and panic-level logs with additional information like code location (file, function, line number).
package log

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/rs/zerolog"
)

// output is the default writer for the logger, which is set to standard output (os.Stdout).
var output = os.Stdout //zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

// Logger is a global instance of zerolog.Logger, with timestamps enabled.
var Logger = zerolog.New(output).With().Timestamp().Logger()

// AddCodeLocation adds the file name, function name, and line number where the log event occurred.
// It captures the current code execution location and appends it to the log entry.
//
// Example usage:
// AddCodeLocation(Logger.Info()).Msg("Logging with code location")
func AddCodeLocation(e *zerolog.Event) *zerolog.Event {
	pc := make([]uintptr, 15)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	file := filepath.Base(frame.File)
	fn := filepath.Base(frame.Function)
	line := frame.Line
	return e.Str("file", file).Str("fn", fn).Int("line", line)
}

// WriteDebug logs a debug-level message. Debug logs are used for detailed internal system information.
//
// Example usage:
// WriteDebug("This is a debug message")
func WriteDebug(msg string) {
	Logger.Debug().Msg(msg)
}

// WriteError logs an error with the code location. It is used for logging runtime errors.
//
// Example usage:
//
//	err := errors.New("something went wrong")
//	WriteError(err)
func WriteError(err error) {
	AddCodeLocation(Logger.Err(err)).Send()
}

// WriteErrorWithMsg logs an error along with an additional message string.
//
// Example usage:
//
//	err := errors.New("file not found")
//	WriteErrorWithMsg(err, "Failed to open file")
func WriteErrorMsg(msg string) {
	AddCodeLocation(Logger.Error()).Msg(msg)
}

func WriteErrorWithMsg(err error, msg string) {
	AddCodeLocation(Logger.Err(err)).Msg(msg)
}

// WritePanic logs a panic-level error, which is used for critical issues that require immediate attention.
// It logs the error and includes the code location.
//
// Example usage:
//
//	err := errors.New("critical failure")
//	WritePanic(err)
func WritePanic(err error) {
	AddCodeLocation(Logger.Panic().Err(err)).Send()
}

// OffsetNotTypeInteger is a constant string used to represent an error when an offset is not of type integer.
var OffsetNotTypeInteger = "Offset is not type integer"

// Error should be used for Internal API errors. Any client side error should have a high-level log with Info
// or a low level log with Debug.

// Debug should be used for low-level logs, such as database transactions or anything detailing the API logic only.

// Info should be used for logging API requests at the high level.
