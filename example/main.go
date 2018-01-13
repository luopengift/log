package main

import (
	"github.com/luopengift/log"
	"os"
)

func main() {
	//最简单的使用方式
	log.Debug("This is a Debug.")
	log.Info("This is a Info.")
	log.Warn("This is a Warn.")
	log.Error("This is a Error.")
	log.SetLogger("gohttp", log.NewLog("Stdout", os.Stdout))
	log.GetLogger("gohttp").Info("This is a gohttp info")
	//	log.Display("%#v", log.Logger)

	// filehandler
	filehandler := log.NewLog("file", log.NewFile("/tmp/test.log", 10))
	for _, _ = range []int{1, 2, 3, 4} {
		filehandler.Info("1111111")
	}
}
