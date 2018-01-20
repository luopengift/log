package log

import (
	"fmt"
	"os"
)

var (
	// Logger manage mutli log handler
	Logger map[string]*Log
	// _Default name with default log handler
	_Default string
)

// GetLogger get Log with name,
// if not set then returns default Log.
func GetLogger(name string) *Log {
	if l, ok := Logger[name]; ok {
		return l
	}
	return Logger["__ROOT__"]
}

// SetLogger update Log.
func SetLogger(name string, l *Log) error {
	if _, ok := Logger[name]; ok {
		return fmt.Errorf("%s is exist in logger", name)
	}
	Logger[name] = l
	return nil
}

// SetDefaultLogger set default Log with named Log.
func SetDefaultLogger(name string) {
	_Default = name
}

// ResetDefaultLogger set default Log with __ROOT__.
func ResetDefaultLogger() {
	_Default = "__ROOT__"
}

// AddLogger add a Log to Logger.
func AddLogger(l *Log) error {
	Logger[l.name] = l
	return nil
}

// DelLogger delete a named Log.
func DelLogger(name string) {
	delete(Logger, name)
}

func init() {
	_Default = "__ROOT__"
	Logger = make(map[string]*Log)

	root := NewLog("__ROOT__", os.Stdout)
	root.SetFormatter(NewTextFormat(DEFAULT_FORMAT, 1))
	root.SetCallDepth(3)
	AddLogger(root)
}
