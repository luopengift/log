package log

import (
	"fmt"
	"os"
)

var Logger map[string]*Log

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

func DelLogger(name string) {
	delete(Logger, name)
}

func init() {
	Logger = make(map[string]*Log)
	SetLogger("__ROOT__", NewLog("__ROOT__", os.Stdout))
	GetLogger("__ROOT__").SetCallDepth(3)
}
