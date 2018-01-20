package log

import (
	"fmt"
)

const (
	red     = uint8(iota + 91) //红色
	green                      //绿色
	yellow                     //黄色
	blue                       //蓝色
	magenta                    //洋红
	blue2                      //湖蓝
	none    = uint8(0)
)

func color(col uint8, s interface{}) string {
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", col, s)
}
