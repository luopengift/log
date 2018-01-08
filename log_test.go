package log

import (
	"fmt"
	"testing"
)

/*
func Test_File(t *testing.T) {
	log := New("service")
	log.Error("hello")
}
*/
func Test_Stdout(t *testing.T) {
	fmt.Println("0&0=", 0|0)
	fmt.Println("0&1=", 0|1)
	fmt.Println("1&0=", 1|0)
	fmt.Println("1&1=", 1|1)

	log := New("log", &Stdout{})
	SetLogger("hello", "")
	log.SetFormat(&JSONFormat{})
	log.SetLevel(WARN)
	log.format(0, "")
	fmt.Println(LogMap)
	log.Debug("DEUBG")
	log.Info("INFO")
	log.Warn("WARN")
	log.Error("ERROR")
	log.Fatal("FATAL")
	//log.Panic("PANIC")
}
