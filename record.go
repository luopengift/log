package log

import (
	"path"
	"strconv"
	"strings"
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

func (rcd *LogRecord) Format(str string) string {
	str = strings.Replace(str, "TIME", rcd.Time, -1)
	str = strings.Replace(str, "LEVEL", LevelMap[rcd.Level], -1)
	str = strings.Replace(str, "MODULE", rcd.Module, -1)
	str = strings.Replace(str, "FUNCNAME", FuncName(rcd.FuncPtr), -1)
	str = strings.Replace(str, "PATH", path.Dir(rcd.File), -1)
	str = strings.Replace(str, "FILE", path.Base(rcd.File), -1)
	str = strings.Replace(str, "LINE", strconv.Itoa(rcd.Line), -1)
	str = strings.Replace(str, "MESSAGE", rcd.Msg, -1)
	return str
}
