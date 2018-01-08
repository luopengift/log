package log

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const (
	ModeSync        = 1 << iota //同步
	ModeAsync                   //异步
	ModeTime                    //时间轮转
	ModeSize                    //大小轮转
	ModeAppendDelim             //日志尾部追加delim
)

//日志的基本配置
type Log struct {
	name      string
	mode      int      //日志模式
	level     uint8    //最低的告警级别
	delim     string   //每条日志的分隔符
	Formatter          //日志格式化接口,实现Format() string
	out       []Logger //日志输出点
}

func New(name string, out ...Logger) *Log {
	return &Log{name: name, mode: ModeSync | ModeAppendDelim, level: DEBUG, delim: "\n", Formatter: &PlainFormat{}, out: out}
}

func (l *Log) SetOutput(out ...Logger) {
	l.out = out
}

func (l *Log) SetFormat(format Formatter) {
	l.Formatter = format
}

func (l *Log) SetLevel(level uint8) {
	l.level = level
}
func (l *Log) SetMode(mode int) {
	l.mode = mode
}
func (l *Log) SetDelim(delim string) {
	l.delim = delim
}

func (l *Log) format(lv uint8, format string) string {
	if l.mode&ModeAsync != 0 {
		//异步模式
		fmt.Println("异步")
	}
	if l.mode&ModeSync != 0 {
		fmt.Println("同步")
	}
	if l.mode&ModeAppendDelim != 0 {
		fmt.Println("追加delim")
	}
	return ""
}

//2018/01/08 14:34:34.000 [I] This is a log.
func (l *Log) prefix() string {
	p := make([]string, 5)
	p[0] = time.Now().Format("2006/01/02 15:04:05.000")
	return strings.Join(p, " ")

}

func (l *Log) Output(lv uint8, format string, v ...interface{}) {
	if lv < l.level {
		return
	}
	ctn := &Content{
		Time:  time.Now(),
		Level: lv,
		body:  fmt.Sprintf(format, v...),
	}
	if l.mode&ModeAppendDelim != 0 {
		_, ctn.File, ctn.Line, _ = runtime.Caller(2)
	}

	msg := l.Format(ctn)

	if l.mode&ModeAppendDelim != 0 {
		msg += l.delim
	}
	for _, out := range l.out {
		fmt.Fprint(out, msg)
	}
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
