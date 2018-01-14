package log

import "fmt"

func Debug(format string, v ...interface{}) {
	GetLogger(__default__).Debug(format, v...)
}
func Info(format string, v ...interface{}) {
	GetLogger(__default__).Info(format, v...)
}
func Warn(format string, v ...interface{}) {
	GetLogger(__default__).Warn(format, v...)
}
func Error(format string, v ...interface{}) {
	GetLogger(__default__).Error(format, v...)
}
func Fatal(format string, v ...interface{}) {
	GetLogger(__default__).Error(format, v...)
}

func ConsoleWithRed(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Println(red(msg))
}

func ConsoleWithGreen(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Println(green(msg))
}

func ConsoleWithYellow(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Println(yellow(msg))
}

func ConsoleWithBlue(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Println(blue(msg))
}

func ConsoleWithMagenta(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Println(magenta(msg))
}
