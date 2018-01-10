package log

import (
	"fmt"
	"os"
	"testing"
)

func Test_File(t *testing.T) {
	filelog := NewFile("test-%Y%M%D.log", 1000)
	l := New("file_test", filelog)
	l.SetFormatter(&JSONFormat{})
	l.Debug("12345")
}


func Test_Stdout(t *testing.T) {
	fmt.Println("0&0=", 0|0)
	fmt.Println("0&1=", 0|1)
	fmt.Println("1&0=", 1|0)
	fmt.Println("1&1=", 1|1)

	//filelog := New("log", NewFile("test-%Y%M%D.log", 1000))
	log := New("console", os.Stdout)
	//log.SetFormatter(&TextFormat{})
	fmt.Println(LogMap)

	log.Debug("DEUBG")
	log.Info("INFO")
	log.Warn("WARN")
	log.Error("ERROR")
	log.Fatal("FATAL")

	Error("%s", "$$$$$")
}
