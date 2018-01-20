package log

import (
	"fmt"
	"io/ioutil"
)

// SetTextFormat sets log message format
func SetTextFormat(format string, mode int) {
	GetLogger(_Default).SetFormatter(NewTextFormat(format, mode))
}

// SetTimeFormat sets time format, if TIME is enabled.
func SetTimeFormat(format string) {
	GetLogger(_Default).SetTimeFormat(format)
}

// Trace calls default output to write the log as level trace.
func Trace(format string, v ...interface{}) {
	GetLogger(_Default).Trace(format, v...)
}

// Debug calls default output to write the log as level debug.
func Debug(format string, v ...interface{}) {
	GetLogger(_Default).Debug(format, v...)
}

// Info calls default output to write the log as level info.
func Info(format string, v ...interface{}) {
	GetLogger(_Default).Info(format, v...)
}

// Warn calls default output to write the log as level warn.
func Warn(format string, v ...interface{}) {
	GetLogger(_Default).Warn(format, v...)
}

// Error calls default output to write the log as level error.
func Error(format string, v ...interface{}) {
	GetLogger(_Default).Error(format, v...)
}

// Fatal calls default output to write the log as level fatal.
func Fatal(format string, v ...interface{}) {
	GetLogger(_Default).Error(format, v...)
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

// ConsoleWithRed write message to stdout with red color
func ConsoleWithRed(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Println(color(red, msg))
}

// ConsoleWithGreen write message to stdout with red color
func ConsoleWithGreen(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Println(color(green, msg))
}

// ConsoleWithYellow write message to stdout with yellow color.
func ConsoleWithYellow(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Println(color(yellow, msg))
}

// ConsoleWithBlue write message to stdout with blue color.
func ConsoleWithBlue(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Println(color(blue, msg))
}

// ConsoleWithMagenta write message to stdout with magenta color.
func ConsoleWithMagenta(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Println(color(magenta, msg))
}
