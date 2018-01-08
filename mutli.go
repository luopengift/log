package log

import (
//"os"
)

var LogMap map[string]*Log

func Register(name string) {
	LogMap[name] = New(name, &Stdout{})
}

func GetLogger(name string) *Log {
	if l, ok := LogMap[name]; ok {
		return l
	}
	return LogMap["__DEFAULT__"]
}

func SetLogger(name string, config string) error {
	if _, ok := LogMap[name]; ok {
		return nil //l.Init(config)
	}
	return nil
}

func DelLogger(name string) {
	delete(LogMap, name)
}

func init() {
	LogMap = make(map[string]*Log)
	Register("__DEFAULT__")
}
