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

func AddLogger(l *Log) error {
	Logger[l.name] = l
	return nil
}

func DelLogger(name string) {
	delete(Logger, name)
}

func init() {
	Logger = make(map[string]*Log)
	_root_ := NewLog("__ROOT__", os.Stdout)
	_root_.SetCallDepth(3)
	AddLogger(_root_)
}