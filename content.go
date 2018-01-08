package log

import (
	"fmt"
	"time"
	"encoding/json"
)

//日志头
type Content struct {
	Time   time.Time `json:"timesatmp,omitempty"` //时间
	Level  uint8     `json:"level,omitempty"`     //日志级别
	Module string    `json:"module,omitempty"`    //日志钩子(loghandler),默认是"__DEFAULT__"
	File   string    `json:"file,omitempty"`      //文件名
	Line   int       `json:"line,omitempty"`      //日志行号
	body   string    `json:"body,omitempty"`
}

func (ctn *Content) String() string {
	return fmt.Sprintf("%s| %1s [%s] %s:%d %s", ctn.Time.Format("2006/01/02 15:04:05.000"), ctn.Level, ctn.Module, ctn.File, ctn.Line, ctn.body)
}

func (ctn *Content) JSON() string {
	b, _ := json.Marshal(ctn)
	return string(b)
}
