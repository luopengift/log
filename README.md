## log

### useage example

```
package main

import "github.com/luopengift/log"

func main() {
	log.Info("%s", "Hello") //默认用法

	file := log.NewFile("filename")
	file.SetMaxBytes(100)
	mylogs := log.NewLog("file", file)
	mylogs.SetFormatter(&log.TextFormat{})
	mylogs.SetLevel(log.DEBUG)
	mylogs.Info("%s", "This is a log test")

	log.AddLogger(mylogs)
	log.GetLogger("file").Info("%s", "This is a logger test")
	log.GetLogger("file").Warn("This is a warn test")
	log.GetLogger("file").Error("This is a error test")

}
```



