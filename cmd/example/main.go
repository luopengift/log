package main

import "github.com/luopengift/log"

func main() {
	filelog := log.NewFile("logs/test.log")
	filelog.SetMaxLines(2)
	filelog.SetMaxIndex(30)
	l := log.NewLog("file_test", filelog)
	//l.SetDelim("")
	l.SetFormatter(&log.NullFormat{})
	for i := 0; i < 20; i++ {
		l.Debugf("%d", i)
	}
}
