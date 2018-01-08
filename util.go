package log

import (
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
