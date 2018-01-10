package log

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func Test_File(t *testing.T) {
	filelog := NewFile("/tmp/test-%Y%M%D%h%m%s.log", 1000)
	l := NewLog("file_test", filelog)
	l.SetFormatter(&JSONFormat{})
	for i := 0; i < 10; i++ {
		time.Sleep(400 * time.Millisecond)
		l.Debug("%d", i)
	}

}

func Test_Stdout(t *testing.T) {
	fmt.Println("0&0=", 0|0)
	fmt.Println("0&1=", 0|1)
	fmt.Println("1&0=", 1|0)
	fmt.Println("1&1=", 1|1)

	//filelog := New("log", NewFile("test-%Y%M%D.log", 1000))
	log := NewLog("console", os.Stdout)
	//log.SetFormatter(&TextFormat{})
	fmt.Println(LogMap)

	log.Debug("DEUBG")
	log.Info("INFO")
	log.Warn("WARN")
	log.Error("ERROR")
	log.Fatal("FATAL")

	Error("%s", "$$$$$")
}

