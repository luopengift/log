package log

import (
	"fmt"
	"os"
)

var Logger map[string]*Log
var _default string

func GetLogger(name string) *Log {
	if l, ok := Logger[name]; ok {
		return l
	}
	return Logger["__ROOT__"]
}

func SetLogger(name string, l *Log) error {
	if _, ok := Logger[name]; ok {
		return fmt.Errorf("%s is exist in logger", name)
	}
	Logger[name] = l
	return nil
}

func SetDefaultLogger(name string) {
	_default = name
}

func ResetDefaultLogger() {
	_default = "__ROOT__"
}

func AddLogger(l *Log) error {
	Logger[l.name] = l
	return nil
}

func DelLogger(name string) {
	delete(Logger, name)
}

func init() {
	_default = "__ROOT__"
	Logger = make(map[string]*Log)

	root := NewLog("__ROOT__", os.Stdout)
	AddLogger(root)
}
