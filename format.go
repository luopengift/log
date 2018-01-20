package log

import (
	"encoding/json"
	"fmt"
)

const (
	// DEFAULTFORMAT default log message format.
	DEFAULTFORMAT = "TIME [LEVEL] FILE:LINE MESSAGE"
	// ModeColor color mode.
	ModeColor = 1 << iota
)

// Formatter interface
type Formatter interface {
	Format(rcd *Record) string
}

// NullFormat implements Formatter interface.
type NullFormat struct{}

// Format only output log msg.
func (f *NullFormat) Format(rcd *Record) string {
	return rcd.Msg
}

//JSONFormat implements Formatter interface.
type JSONFormat struct{}

// Format marshal log record.
func (f *JSONFormat) Format(rcd *Record) string {
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
func (f *TextFormat) Format(rcd *Record) string {
	msg := rcd.Format(f.format)
	if f.mode&ModeColor != 0 {
		return color(LevelColor[rcd.Level], msg)
	}
	return msg
}

// KvFormat implements Formatter interface.
type KvFormat struct{}

// Format TODO.
func (f *KvFormat) Format(rcd *Record) string {
	return fmt.Sprintf("TODO:kv format, %s", rcd.Msg)
}
