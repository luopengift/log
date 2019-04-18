package log

import (
	"os"
	"testing"
)

func Test_Default(t *testing.T) {
	Infof("Debug")
	Warnf("watn")
}

func Test_File(t *testing.T) {
	filelog := NewFile("/tmp/zzz/test.log")
	filelog.SetMaxLines(2)
	l := NewLog("file_test", filelog)
	l.SetDelim("")
	l.SetFormatter(&NullFormat{})
	for i := 0; i < 10; i++ {
		l.Debugf("%d", i)
	}

}

func Test_Stdout(t *testing.T) {
	log := NewLog("console", os.Stdout)
	log.SetFormatter(NewTextFormat("TIME LEVEL MODULE PATH FILE:FUNCNAME:LINE MESSAGE", ModeColor))
	// fmt.Println(Logger)

	log.Debugf("DEUBG")
	log.Infof("INFO")
	log.Warnf("WARN")
	log.Errorf("ERROR")
	log.Fatalf("FATAL")

	Errorf("%s", "$$$$$")
	//dump("log", log)
}
