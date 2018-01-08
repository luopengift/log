/**
  打印字符颜色
*/
package log

import (
	"fmt"
)

const (
	__RED__     = uint8(iota + 91) //红色
	__GREEN__                      //绿色
	__YELLOW__                     //黄色
	__BLUE__                       //蓝色
	__MAGENTA__                    //洋红
	__BLUE2__                      //湖蓝
)

func Color(col uint8, s interface{}) string {
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", col, s)
}

func None(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func Red(v interface{}) string {
	return Color(__RED__, v)
}

func Green(v interface{}) string {
	return Color(__GREEN__, v)
}

func Yellow(v interface{}) string {
	return Color(__YELLOW__, v)
}

func Blue(v interface{}) string {
	return Color(__BLUE__, v)
}

func Magenta(v interface{}) string {
	return Color(__MAGENTA__, v)
}

func Blue2(v interface{}) string {
	return Color(__BLUE2__, v)
}

func setColor(lv uint8, v interface{}) string {
	switch lv {
	case TRACE:
		return Green(v)
	case DEBUG:
		return Blue(v)
	case INFO:
		return None(v)
	case WARN:
		return Yellow(v)
	case ERROR:
		return Magenta(v)
	case PANIC:
		return Red(v)
	default:
		return Blue2(v)
	}
}
