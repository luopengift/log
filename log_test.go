package log

import (
	"fmt"
	"os"
	"testing"
)

func Test_File(t *testing.T) {
	filelog := NewFile("/tmp/zzz/test.log")
	filelog.SetMaxLines(2)
	l := NewLog("file_test", filelog)
	l.SetDelim("")
	l.SetFormatter(&NullFormat{})
	for i := 0; i < 10; i++ {
		l.Debug("%d", i)
	}

}

func Test_Stdout(t *testing.T) {
	log := NewLog("console", os.Stdout)
	log.SetFormatter(NewTextFormat("TIME LEVEL MODULE PATH FILE:FUNCNAME:LINE MESSAGE", 1))
	fmt.Println(Logger)

	log.Debug("DEUBG")
	log.Info("INFO")
	log.Warn("WARN")
	log.Error("ERROR")
	log.Fatal("FATAL")

	Error("%s", "$$$$$")
	Display("log", log)
}
