package main

import (
	"github.com/luopengift/log"
	"os"
)

func main() {
	log.Debug("This is a Debug.")
	log.SetLogger("gohttp", log.NewLog("Stdout",os.Stdout))
	log.GetLogger("gohttp").Info("This is a gohttp info")
}
