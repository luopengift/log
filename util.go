package log

import (
	"os"
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

// NameWithTime formats string by time.
func NameWithTime(str string, t ...time.Time) string {
	var now time.Time
	if len(t) == 0 {
		now = time.Now()
	} else {
		now = t[1]
	}
	for k, v := range m {
		str = strings.Replace(str, k, now.Format(v), -1)
	}
	return str
}

// FuncName get function name.
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

type files []os.FileInfo

func (fs files) Len() int {
	return len(fs)
}

func (fs files) Less(i, j int) bool {
	return fs[i].ModTime().After(fs[j].ModTime())
}

func (fs files) Swap(i, j int) {
	fs[i], fs[j] = fs[j], fs[i]
}
