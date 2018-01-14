package log

import (
	"encoding/json"
	"fmt"
)

const DEFAULT_FORMAT = "TIME [LEVEL] FILE:LINE MESSAGE"

type Formatter interface {
	Format(rcd *LogRecord) string
}

// This is equal fmt.Sprintf()
type NullFormat struct{}

func (f *NullFormat) Format(rcd *LogRecord) string {
	return rcd.Msg
}

//json
type JSONFormat struct{}

func (f *JSONFormat) Format(rcd *LogRecord) string {
	b, _ := json.Marshal(rcd)
	return string(b)
}

type TextFormat struct {
	format string
}

func NewTextFormat(f ...string) Formatter {
	format := ""
	if len(f) == 0 {
		format = DEFAULT_FORMAT
	} else {
		format = f[0]
	}

	return &TextFormat{format: format}
}

func (f *TextFormat) Format(rcd *LogRecord) string {
	return rcd.Format(f.format)
}

type ConsoleFormat struct {
	format string
}

func NewConsoleFormat(f ...string) Formatter {
	format := ""
	if len(f) == 0 {
		format = DEFAULT_FORMAT
	} else {
		format = f[0]
	}

	return &ConsoleFormat{format: format}
}
func (f *ConsoleFormat) Format(rcd *LogRecord) string {
	msg := rcd.Format(f.format)
	return setColor(rcd.Level, msg)
}

type KvFormat struct{}

func (f *KvFormat) Format(rcd *LogRecord) string {
	return fmt.Sprintf("TODO:kv format, %s", rcd.Msg)
}
