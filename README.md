## log

### useage example

```
package main

import (
	"os"
	"github.com/luopengift/log"
)

func main() {
	log.Debug("%s", "Hello Debug") //默认用法
	log.Info("%s", "Hello Info") //默认用法
	log.Warn("%s", "Hello Warn") //默认用法

	file := log.NewFile("filename")
	file.SetMaxBytes(100)
	mylogs := log.NewLog("file", file)
	mylogs.SetFormatter(log.NewTextFormat("MESSAGE", log.ModeColor))
	mylogs.SetLevel(log.DEBUG)
	mylogs.Info("%s", "This is a log test")

	std := log.NewLog("std", os.Stdout)
	std.SetFormatter(log.NewTextFormat("TIME LEVEL MODULE FILE:LINE MESSAGE", log.ModeColor))
	std.Warn("THIS IS WARN TEST")

	log.AddLogger(mylogs)
	log.GetLogger("file").Info("%s", "This is a logger test")
	log.GetLogger("file").Warn("This is a warn test")
	log.GetLogger("file").Error("This is a error test")

}
```



