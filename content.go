package log

import (
	"encoding/json"
	"fmt"
	"path"
)

type Content struct {
	Time   string `json:"timesatmp,omitempty"` //时间
	Level  uint8  `json:"level,omitempty"`     //日志级别
	Module string `json:"module,omitempty"`    //日志钩子(loghandler),默认是"__DEFAULT__"
	File   string `json:"file,omitempty"`      //文件名
	Line   int    `json:"line,omitempty"`      //日志行号
	body   string `json:"body,omitempty"`
}

type Formatter interface {
	Format(ctn *Content) string
}

// This is equal fmt.Sprintf()
type NullFormat struct{}

func (f *NullFormat) Format(ctn *Content) string {
	return ctn.body
}

//json
type JSONFormat struct{}

func (f *JSONFormat) Format(ctn *Content) string {
	b, _ := json.Marshal(ctn)
	return string(b)
}

type TextFormat struct{}

func (f *TextFormat) Format(ctn *Content) string {
	return fmt.Sprintf("%s| %-5s [%s] %s:%d %s", ctn.Time, LevelMap[ctn.Level], ctn.Module, path.Base(ctn.File), ctn.Line, ctn.body)
}

type ConsoleFormat struct {
}

func (f *ConsoleFormat) Format(ctn *Content) string {
	return setColor(ctn.Level, fmt.Sprintf(
		"%s [%s] %s %s:%d %s",
		ctn.Time, LevelMap[ctn.Level], ctn.Module, path.Base(ctn.File), ctn.Line, ctn.body,
	))
}

type KvFormat struct{}

func (f *KvFormat) Format(ctn *Content) string {
	return fmt.Sprintf("TODO:kv format, %s", ctn.body)
}
