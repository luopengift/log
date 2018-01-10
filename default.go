package log

import "fmt"

func Debug(format string, v ...interface{}) {
	GetLogger("__ROOT__").Debug(format, v...)
}
func Info(format string, v ...interface{}) {
	GetLogger("__ROOT__").Debug(format, v...)
}
func Warn(format string, v ...interface{}) {
	GetLogger("__ROOT__").Debug(format, v...)
}
func Error(format string, v ...interface{}) {
	GetLogger("__ROOT__").Debug(format, v...)
}
func Fatal(format string, v ...interface{}) {
	GetLogger("__ROOT__").Debug(format, v...)
}

func ConsoleWithRed(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Println(Red(msg))
}

func ConsoleWithGreen(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Println(Green(msg))
}

func ConsoleWithYellow(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Println(Yellow(msg))
}

func ConsoleWithBlue(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Println(Blue(msg))
}

func ConsoleWithMagenta(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	fmt.Println(Magenta(msg))
}
