package log

import (
	"fmt"
	"os"
)

var Logger map[string]*Log
var __default__ string

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
	__default__ = name
}

func ResetDefaultLogger() {
	__default__ = "__ROOT__"
}

func AddLogger(l *Log) error {
	Logger[l.name] = l
	return nil
}

func DelLogger(name string) {
	delete(Logger, name)
}

func init() {
	__default__ = "__ROOT__"
	Logger = make(map[string]*Log)

	root := NewLog("__ROOT__", os.Stdout)
	root.SetFormatter(NewTextFormat(DEFAULT_FORMAT, 1))
	root.SetCallDepth(3)
	AddLogger(root)
}
