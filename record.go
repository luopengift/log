package log

import (
	"encoding/json"
	"fmt"
	"path"
)

type LogRecord struct {
	Time    string  `json:"timesatmp,omitempty"` //时间
	Level   uint8   `json:"level,omitempty"`     //日志级别
	Module  string  `json:"module,omitempty"`    //日志钩子(loghandler),默认是"__ROOT__"
	FuncPtr uintptr `json:"funcPtr,omitempty"`   //函数名
	File    string  `json:"file,omitempty"`      //文件名
	Line    int     `json:"line,omitempty"`      //日志行号
	Msg     string  `json:"message,omitempty"`
}

type Formatter interface {
	Format(ctn *LogRecord) string
}

// This is equal fmt.Sprintf()
type NullFormat struct{}

func (f *NullFormat) Format(ctn *LogRecord) string {
	return ctn.Msg
}

//json
type JSONFormat struct{}

func (f *JSONFormat) Format(ctn *LogRecord) string {
	b, _ := json.Marshal(ctn)
	return string(b)
}

type TextFormat struct{}

func (f *TextFormat) Format(ctn *LogRecord) string {
	return fmt.Sprintf("%s| %s [%s] %s %s:%d %s", ctn.Time, LevelMap[ctn.Level], ctn.Module, FuncName(ctn.FuncPtr), path.Base(ctn.File), ctn.Line, ctn.Msg)
}

type ConsoleFormat struct {}

func (f *ConsoleFormat) Format(ctn *LogRecord) string {
	return setColor(ctn.Level, fmt.Sprintf(
		"%s [%s] %s %s %s:%d %s",
		ctn.Time, LevelMap[ctn.Level], ctn.Module, FuncName(ctn.FuncPtr), path.Base(ctn.File), ctn.Line, ctn.Msg,
	))
}

type KvFormat struct{}

func (f *KvFormat) Format(ctn *LogRecord) string {
	return fmt.Sprintf("TODO:kv format, %s", ctn.Msg)
}


