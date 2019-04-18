package log

import (
	"fmt"
	"os"
)

// Console simple consile log
type Console struct {
}

// NewConsoleLog default log print into stderr
func NewConsoleLog(opt ...interface{}) *Console {
	return &Console{}
}

// Debugf debugf
func (*Console) Debugf(s string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", v...)
}

// Infof infof
func (*Console) Infof(s string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", v...)
}

// Warnf warnf
func (*Console) Warnf(s string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", v...)
}

// Errorf errorf
func (*Console) Errorf(s string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", v...)
}

// Fatalf fatalf
func (*Console) Fatalf(s string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", v...)
}
