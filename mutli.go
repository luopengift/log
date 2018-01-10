package log

import (
	"fmt"
	"os"
)

var LogMap map[string]*Log

func GetLogger(name string) *Log {
	if l, ok := LogMap[name]; ok {
		return l
	}
	return LogMap["__ROOT__"]
}

func SetLogger(name string, l *Log) error {
	if _, ok := LogMap[name]; ok {
		return fmt.Errorf("%s is exist in logger", name)
	}
	LogMap[name] = l
	return nil
}

func DelLogger(name string) {
	delete(LogMap, name)
}

func init() {
	LogMap = make(map[string]*Log)
	SetLogger("__ROOT__", New("__ROOT__", os.Stdout))
	GetLogger("__ROOT__").SetCallDepth(3)
}
