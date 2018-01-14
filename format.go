package log

import (
	"encoding/json"
	"fmt"
)

const DEFAULT_FORMAT = "TIME [LEVEL] FILE:LINE MESSAGE"

const (
	ModeColor = 1 << iota
)

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
	mode   int
}

func NewTextFormat(f string, mode int) Formatter {
	return &TextFormat{format: f, mode: mode}
}

func (f *TextFormat) Format(rcd *LogRecord) string {
	msg := rcd.Format(f.format)
	if f.mode&ModeColor != 0 {
		return setColor(rcd.Level, msg)
	}
	return msg
}

type KvFormat struct{}

func (f *KvFormat) Format(rcd *LogRecord) string {
	return fmt.Sprintf("TODO:kv format, %s", rcd.Msg)
}
