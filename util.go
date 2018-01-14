package log

import (
	"runtime"
	"strings"
	"time"
)

var (
	m = map[string]string{
		"%Y": "2006",
		"%M": "01",
		"%D": "02",
		"%h": "15",
		"%m": "04",
		"%s": "05",
	}
)

func NameWithTime(str string) string {
	for k, v := range m {
		str = strings.Replace(str, k, time.Now().Format(v), -1)
	}
	return str
}

func FuncName(pc uintptr) string {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "???"
	}
	name := fn.Name()
	// The name includes the path name to the package, which is unnecessary
	// since the file name is already included.  Plus, it has center dots.
	// That is, we see
	//	runtime/debug.*T·ptrmethod
	// and want
	//	*T.ptrmethod
	if period := strings.LastIndex(name, "."); period >= 0 {
		name = name[period+1:]
	}
	return name
}

