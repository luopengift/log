package log

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var log = NewLog("root", os.Stderr)

// SetOutput sets default log output.
func SetOutput(out ...io.Writer) *Log {
	return log.SetOutput(out...)
}

// SetTextFormat sets default log message format.
func SetTextFormat(format string, mode int) *Log {
	return log.SetFormatter(NewTextFormat(format, mode))
}

// SetTimeFormat sets default time format, if TIME is enabled.
func SetTimeFormat(format string) *Log {
	return log.SetTimeFormat(format)
}

// SetLevel set default level.
func SetLevel(level uint8) *Log {
	return log.SetLevel(level)
}

// Output output
func Output(format string, v ...interface{}) *Log {
	return log.Output(format, v...)
}

// Trace calls default output to write the log as level trace.
func Trace(format string, v ...interface{}) *Log {
	return log.Trace(format, v...)
}

// Debugf calls default output to write the log as level debug.
func Debugf(format string, v ...interface{}) {
	log.Debugf(format, v...)
}

// Infof calls default output to write the log as level info.
func Infof(format string, v ...interface{}) {
	log.Infof(format, v...)
}

// Warnf calls default output to write the log as level warn.
func Warnf(format string, v ...interface{}) {
	log.Warnf(format, v...)
}

// Errorf calls default output to write the log as level error.
func Errorf(format string, v ...interface{}) {
	log.Errorf(format, v...)
}

// Fatalf calls default output to write the log as level fatal.
func Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

// Panic calls default output to write the log as level panic.
func Panic(format string, v ...interface{}) {
	log.Panic(format, v...)
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
