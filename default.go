package log

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// SetOutput sets default log output.
func SetOutput(out io.Writer) *Log {
	return GetLogger(_Default).SetOutput(out)
}

// SetTextFormat sets default log message format.
func SetTextFormat(format string, mode int) *Log {
	return GetLogger(_Default).SetFormatter(NewTextFormat(format, mode))
}

// SetTimeFormat sets default time format, if TIME is enabled.
func SetTimeFormat(format string) *Log {
	return GetLogger(_Default).SetTimeFormat(format)
}

// SetLevel set default level.
func SetLevel(level uint8) *Log {
	return GetLogger(_Default).SetLevel(level)
}

// Trace calls default output to write the log as level trace.
func Trace(format string, v ...interface{}) *Log {
	return GetLogger(_Default).Trace(format, v...)
}

// Debug calls default output to write the log as level debug.
func Debug(format string, v ...interface{}) *Log {
	return GetLogger(_Default).Debug(format, v...)
}

// Info calls default output to write the log as level info.
func Info(format string, v ...interface{}) *Log {
	return GetLogger(_Default).Info(format, v...)
}

// Warn calls default output to write the log as level warn.
func Warn(format string, v ...interface{}) *Log {
	return GetLogger(_Default).Warn(format, v...)
}

// Error calls default output to write the log as level error.
func Error(format string, v ...interface{}) *Log {
	return GetLogger(_Default).Error(format, v...)
}

// Fatal calls default output to write the log as level fatal.
func Fatal(format string, v ...interface{}) *Log {
	return GetLogger(_Default).Error(format, v...)
}

// Panic calls default output to write the log as level panic.
func Panic(format string, v ...interface{}) {
	GetLogger(_Default).Panic(format, v...)
}

// Errorf returns error interface with error message.
func Errorf(format string, v ...interface{}) error {
	return GetLogger(_Default).Errorf(format, v...)
}

// OutputWithFile write to file, It truncates file before writing.
func OutputWithFile(name string, format string, v ...interface{}) error {
	msg := fmt.Sprintf(format, v...)
	return ioutil.WriteFile(name, []byte(msg), 0644)
}

// ConsoleWithRed write message to stderr with red color
func ConsoleWithRed(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Fprintln(os.Stderr, color(red, msg))
}

// ConsoleWithGreen write message to stderr with red color
func ConsoleWithGreen(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Fprintln(os.Stderr, color(green, msg))
}

// ConsoleWithYellow write message to stderr with yellow color.
func ConsoleWithYellow(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Fprintln(os.Stderr, color(yellow, msg))
}

// ConsoleWithBlue write message to stderr with blue color.
func ConsoleWithBlue(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Fprintln(os.Stderr, color(blue, msg))
}

// ConsoleWithMagenta write message to stderr with magenta color.
func ConsoleWithMagenta(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Fprintln(os.Stderr, color(magenta, msg))
}
