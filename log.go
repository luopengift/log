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
	ModeSync  = 1 << iota //同步
	ModeAsync             //异步
)

//日志的基本配置
type Log struct {
	mux        *sync.Mutex
	pool       sync.Pool //临时对象池
	name       string
	mode       int       //日志模式
	level      uint8     //最低的告警级别
	depth      int       //runtime.Call(depth)
	delim      string    //每条日志的分隔符
	timeFormat string    //日志时间格式
	Formatter            //日志格式化接口,实现Format() string
	out        io.Writer //日志输出点
}

func NewLog(name string, out io.Writer) *Log {
	return &Log{
		mux:        new(sync.Mutex),
		pool:       sync.Pool{New: func() interface{} { return new(LogRecord) }},
		name:       name,
		mode:       ModeSync,
		level:      DEBUG,
		depth:      2,
		delim:      "\n",
		timeFormat: time.RFC3339,
		Formatter:  NewTextFormat(DEFAULT_FORMAT, 0),
		out:        out,
	}
}

// SetOutput sets the output destination for Log.
func (l *Log) SetOutput(out io.Writer) {
	l.out = out
}

// SetTimeFormat sets the output time format for Log.
// Default is RFC3339.
func (l *Log) SetTimeFormat(timeFormat string) {
	l.timeFormat = timeFormat
}

// SetFormatter sets the output log format for Log.
// Input params must implement Formatter interface.
// Default is NewTextFormat(DEFAULT_FORMAT, 0).
func (l *Log) SetFormatter(format Formatter) {
	l.Formatter = format
}

// SetLevel sets the output level for Log.
// Default is DEBUG
func (l *Log) SetLevel(level uint8) {
	l.level = level
}

// SetMode sets the output mode for Log.
// Default is. TODO.
func (l *Log) SetMode(mode int) {
	l.mode = mode
}

// SetDelim sets the output split of Log.
// Default is "\n".
func (l *Log) SetDelim(delim string) {
	l.delim = delim
}

// if warp this package, please reset call depth. Default is 2.
func (l *Log) SetCallDepth(depth int) {
	l.depth = depth
}

// Output writes the output for a logging event.
func (l *Log) Output(lv uint8, format string, v ...interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()
	if lv < l.level {
		return
	}
	ctn := l.pool.Get().(*LogRecord)
	ctn.Time = time.Now().Format(l.timeFormat)
	ctn.Level = lv
	ctn.Module = l.name
	ctn.Msg = fmt.Sprintf(format, v...)
	ctn.FuncPtr, ctn.File, ctn.Line, _ = runtime.Caller(l.depth)
	msg := l.Format(ctn) + l.delim
	fmt.Fprint(l.out, msg)
	l.pool.Put(ctn)
}

// Debug calls l.Output to write the log as level debug.
func (l *Log) Debug(format string, v ...interface{}) {
	l.Output(DEBUG, format, v...)
}

// Info calls l.Output to write the log as level info.
func (l *Log) Info(format string, v ...interface{}) {
	l.Output(INFO, format, v...)
}

// Warn calls l.Output to write the log as level warn.
func (l *Log) Warn(format string, v ...interface{}) {
	l.Output(WARN, format, v...)
}

// Error calls l.Output to write the log as level error.
func (l *Log) Error(format string, v ...interface{}) {
	l.Output(ERROR, format, v...)
}

// Fatal calls l.Output to write the log as level fatal.
func (l *Log) Fatal(format string, v ...interface{}) {
	l.Output(FATAL, format, v...)
}

// 输出当前堆栈信息
func (l *Log) Trace(format string, v ...interface{}) {
	debug.PrintStack()
	l.Output(TRACE, format, v...)
}

// panic: 并且打印当前堆栈信息
func (l *Log) Panic(format string, v ...interface{}) {
	l.Output(PANIC, format, v...)
	panic(fmt.Sprintf(format, v...))
}

// Errorf implement error interface
func (l *Log) Errorf(format string, v ...interface{}) error {
	return fmt.Errorf(format, v...)
}
