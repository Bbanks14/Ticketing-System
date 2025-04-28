package log

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/Bbanks14/Ticketing-System/structs"
	"github.com/Bbanks14/ticketing-system/globals"
	"github.com/Bbanks14/ticketing-system/structs"
)

// logger is the default logger which writes its output to standard format.
// The standard flags are set for the logger, so each output is prepended with date and time
var logger = log.New(os.Stdout, "", log.LstdFlags)

// timeFormat is the representation of the time format the default logger in the `log` package uses.
// It is used to precede the log messages written by the server error log with date and time.
const timeFormat string = "2006-01-02 15:04:05"

// Info writes a standardized lof info message represented through the parameter v
// to the standard log with the INFO log level.
// The message preceded by the log level string and appended by
// the function location (package, filename and line number) where the logging took place. The message is only
// written to the log if the log level is INFO.
func INFO(v ...interface{}) {
	if canLog(structs.LevelInfo) {
		logln(structs.LevelInfo, v...)
	}
}

// Infof writes a standardized lof info message represented through a format string and the corresponding arguments
// to the standard log with the INFO log level.
// The message preceded by the log level string and appended by
// the function location (package, filename and line number) where the logging took place. The message is only
// written to the log if the log level is INFO.
func Infof(format string, v ...interface{}) {
	if canLog(structs.LevelInfo) {
		logf(structs.LevelInfo, format, v...)
	}
}

// Warn writes a standardized lof info message represented through the parameter v
// to the standard log with the WARNING log level.
// The message preceded by the log level string and appended by
// the function location (package, filename and line number) where the logging took place. The message is only
// written to the log if the log level is WARNING or higher.
func Warn(v ...interface{}) {
	if canLog(structs.LevelWarning) {
		logln(structs.LevelWarning, v...)
	}
}

// Warnf writes a standardized lof info message represented through a format string and the corresponding arguments
// to the standard log with the WARNING log level.
// The message preceded by the log level string and appended by
// the function location (package, filename and line number) where the logging took place. The message is only
// written to the log if the log level is WARNING or higher.
func Warnf(format string, v ...interface{}) {
	if canLog(structs.LevelWarning) {
		logf(structs.LevelWarning, format, v...)
	}
}

// Error writes a standardized lof info message represented through the parameter v
// to the standard log with the ERROR log level.
// The message preceded by the log level string and appended by
// the function location (package, filename and line number) where the logging took place. The message is only
// written to the log if the log level is ERROR or higher.
func Error(v ...interface{}) {
	if canLog(structs.LevelError) {
		logln(structs.LevelError, v...)
	}
}

// Errorf writes a standardized lof info message represented through a format string and the corresponding arguments
// to the standard log with the ERROR log level.
// The message preceded by the log level string and appended by
// the function location (package, filename and line number) where the logging took place. The message is only
// written to the log if the log level is ERROR or higher.
func Errorf(format string, v ...interface{}) {
	if canLog(structs.LevelError) {
		logf(structs.LevelError, format, v...)
	}
}
