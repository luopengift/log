// Package log implements a logging package.
package log

import (
	"fmt"
	"io"
	"runtime"
	"runtime/debug"
	"sync"
	"time"
)

const (
	// ModeSync log flush mode
	ModeSync = 1 << iota //同步
	// ModeAsync log flush mode
	ModeAsync //异步
)

// Log handler
type Log struct {
	mux        *sync.Mutex
	pool       sync.Pool //临时对象池
	name       string
	mode       int         //日志模式
	level      uint8       //最低的告警级别
	depth      int         //runtime.Call(depth)
	delim      string      //每条日志的分隔符
	timeFormat string      //日志时间格式
	Formatter              //日志格式化接口,实现Format() string
	outs       []io.Writer //日志输出点
}

// NewLog init a log with default config.
func NewLog(name string, outs ...io.Writer) *Log {
	return &Log{
		mux:        new(sync.Mutex),
		pool:       sync.Pool{New: func() interface{} { return new(Record) }},
		name:       name,
		mode:       ModeSync,
		level:      DEBUG,
		depth:      2,
		delim:      "\n",
		timeFormat: time.RFC3339,
		Formatter:  NewTextFormat(DEFAULTFORMAT, 0),
		outs:       outs,
	}
}

// SetOutput sets the output destination for Log.
func (l *Log) SetOutput(outs ...io.Writer) *Log {
	l.outs = outs
	return l
}

// SetTimeFormat sets the output time format for Log.
// Default is RFC3339.
func (l *Log) SetTimeFormat(timeFormat string) *Log {
	l.timeFormat = timeFormat
	return l
}

// SetFormatter sets the output log format for Log.
// Input params must implement Formatter interface.
// Default is NewTextFormat(DEFAULT_FORMAT, 0).
func (l *Log) SetFormatter(format Formatter) *Log {
	l.Formatter = format
	return l
}

// SetTextFormat sets log message format.
func (l *Log) SetTextFormat(format string, mode int) *Log {
	return l.SetFormatter(NewTextFormat(format, mode))
}

// SetLevel sets the output level for Log.
// Default is DEBUG
func (l *Log) SetLevel(level uint8) *Log {
	l.level = level
	return l
}

// SetMode sets the output mode for Log.
// Default is. TODO.
func (l *Log) SetMode(mode int) *Log {
	l.mode = mode
	return l
}

// SetDelim sets the output split of Log.
// Default is "\n".
func (l *Log) SetDelim(delim string) *Log {
	l.delim = delim
	return l
}

// SetCallDepth calls runtime.Caller.
// if warp this package, reset call depth. Default is 2.
func (l *Log) SetCallDepth(depth int) *Log {
	l.depth = depth
	return l
}

// Output writes the output for a logging event.
func (l *Log) output(lv uint8, format string, v ...interface{}) *Log {
	l.mux.Lock()
	defer l.mux.Unlock()
	if lv < l.level {
		return l
	}
	ctn := l.pool.Get().(*Record)
	ctn.Time = time.Now().Format(l.timeFormat)
	ctn.Level = lv
	ctn.Module = l.name
	ctn.Msg = fmt.Sprintf(format, v...)
	ctn.FuncPtr, ctn.File, ctn.Line, _ = runtime.Caller(l.depth)
	msg := l.Format(ctn) + l.delim
	for _, out := range l.outs {
		fmt.Fprint(out, msg)
	}
	l.pool.Put(ctn)
	return l
}

// Display display v
func (l *Log) Display(format string, v ...interface{}) *Log {
	l.output(WARN, format, string(dump(v)))
	return l
}

// Trace calls l.Output to write the log as level trace,
// and print stack information to stdout.
func (l *Log) Trace(format string, v ...interface{}) *Log {
	debug.PrintStack()
	l.output(TRACE, format, v...)
	return l
}

// Debug calls l.Output to write the log as level debug.
func (l *Log) Debug(format string, v ...interface{}) *Log {
	return l.output(DEBUG, format, v...)
}

// Info calls l.Output to write the log as level info.
func (l *Log) Info(format string, v ...interface{}) *Log {
	return l.output(INFO, format, v...)
}

// Warn calls l.Output to write the log as level warn.
func (l *Log) Warn(format string, v ...interface{}) *Log {
	return l.output(WARN, format, v...)
}

// Error calls l.Output to write the log as level error.
func (l *Log) Error(format string, v ...interface{}) *Log {
	return l.output(ERROR, format, v...)
}

// Fatal calls l.Output to write the log as level fatal.
func (l *Log) Fatal(format string, v ...interface{}) *Log {
	return l.output(FATAL, format, v...)
}

// Panic calls l.Output to write the log as level panic.
func (l *Log) Panic(format string, v ...interface{}) {
	l.output(PANIC, format, v...)
	panic(fmt.Sprintf(format, v...))
}

// Errorf implement error interface
func (l *Log) Errorf(format string, v ...interface{}) error {
	return fmt.Errorf(format, v...)
}
