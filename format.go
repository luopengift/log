package log

import (
	"encoding/json"
	"fmt"
)

const DEFAULT_FORMAT = "TIME [LEVEL] FILE:LINE MESSAGE"

const (
	ModeColor = 1 << iota
)

// Formatter interface
type Formatter interface {
	Format(rcd *LogRecord) string
}

// NullFormat implements Formatter interface.
type NullFormat struct{}

// Format only output log msg.
func (f *NullFormat) Format(rcd *LogRecord) string {
	return rcd.Msg
}

//JSONFormat implements Formatter interface.
type JSONFormat struct{}

// Format marshal log record.
func (f *JSONFormat) Format(rcd *LogRecord) string {
	b, _ := json.Marshal(rcd)
	return string(b)
}

// TextFormat implements Formatter interface.
type TextFormat struct {
	format string
	mode   int
}

// NewTextFormat defines TextFormat. log record format and color.
func NewTextFormat(f string, mode int) Formatter {
	return &TextFormat{format: f, mode: mode}
}

// Format format log record to requird format.
func (f *TextFormat) Format(rcd *LogRecord) string {
	msg := rcd.Format(f.format)
	if f.mode&ModeColor != 0 {
		return setColor(rcd.Level, msg)
	}
	return msg
}

// KvFormat implements Formatter interface.
type KvFormat struct{}

// Format TODO.
func (f *KvFormat) Format(rcd *LogRecord) string {
	return fmt.Sprintf("TODO:kv format, %s", rcd.Msg)
}
