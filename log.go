package log

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	ModeSync  = 1 << iota //同步
	ModeAsync             //异步
	ModeTime              //时间轮转
	ModeSize              //大小轮转
)

type Logger interface {
	Write([]byte) (int, error)
}

//日志的基本配置
type Log struct {
	mux        *sync.Mutex
	pool       sync.Pool //临时对象池
	name       string
	mode       int    //日志模式
	level      uint8  //最低的告警级别
	depth      int    //runtime.Call(depth)
	delim      string //每条日志的分隔符
	timeFormat string //日志时间格式
	Formatter         //日志格式化接口,实现Format() string
	out        Logger //日志输出点
}

func New(name string, out Logger) *Log {
	return &Log{
		mux:       new(sync.Mutex),
		pool:      sync.Pool{New: func() interface{} { return &Content{Module: name} }},
		name:      name,
		mode:      ModeSync,
		level:     DEBUG,
		depth:     2,
		delim:     "\n",
		timeFormat: time.RFC3339,
		Formatter: &ConsoleFormat{},
		out:       out,
	}
}

func (l *Log) SetOutput(out Logger) {
	l.out = out
}

func (l *Log) SetTimeFormat(timeFormat string) {
	l.timeFormat = timeFormat
}
func (l *Log) SetFormatter(format Formatter) {
	l.Formatter = format
}

// SetLevel Default is DEBUG
func (l *Log) SetLevel(level uint8) {
	l.level = level
}

func (l *Log) SetMode(mode int) {
	l.mode = mode
}

// SetDelim SetDelim is set the split of log. Default is "\n"
func (l *Log) SetDelim(delim string) {
	l.delim = delim
}

// if warp this package, please reset call depth. Default is 2.
func (l *Log) SetCallDepth(depth int) {
	l.depth = depth
}

func (l *Log) Output(lv uint8, format string, v ...interface{}) {
	if lv < l.level {
		return
	}
	ctn := l.pool.Get().(*Content)
	ctn.Time = time.Now().Format(l.timeFormat)
	ctn.Level = lv
	ctn.Module = l.name
	ctn.body = fmt.Sprintf(format, v...)
	_, ctn.File, ctn.Line, _ = runtime.Caller(l.depth)
	msg := l.Format(ctn) + l.delim
	fmt.Fprint(l.out, msg)
	l.pool.Put(ctn)
}

// 输出当前堆栈信息
func (l *Log) Trace(format string, v ...interface{}) {
	l.Output(TRACE, format, v...)
}
func (l *Log) Debug(format string, v ...interface{}) {
	l.Output(DEBUG, format, v...)
}
func (l *Log) Info(format string, v ...interface{}) {
	l.Output(INFO, format, v...)
}
func (l *Log) Warn(format string, v ...interface{}) {
	l.Output(WARN, format, v...)
}
func (l *Log) Error(format string, v ...interface{}) {
	l.Output(ERROR, format, v...)
}
func (l *Log) Fatal(format string, v ...interface{}) {
	l.Output(FATAL, format, v...)
}

// panic: 并且打印当前堆栈信息
func (l *Log) Panic(format string, v ...interface{}) {
	l.Output(PANIC, format, v...)
	panic(fmt.Sprintf(format, v...))
}
