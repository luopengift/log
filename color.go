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

func color(col uint8, s interface{}) string {
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", col, s)
}

func none(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func red(v interface{}) string {
	return color(__RED__, v)
}

func green(v interface{}) string {
	return color(__GREEN__, v)
}

func yellow(v interface{}) string {
	return color(__YELLOW__, v)
}

func blue(v interface{}) string {
	return color(__BLUE__, v)
}

func magenta(v interface{}) string {
	return color(__MAGENTA__, v)
}

func blue2(v interface{}) string {
	return color(__BLUE2__, v)
}

func setColor(lv uint8, v interface{}) string {
	switch lv {
	case TRACE:
		return green(v)
	case DEBUG:
		return blue(v)
	case INFO:
		return none(v)
	case WARN:
		return yellow(v)
	case ERROR:
		return magenta(v)
	case PANIC:
		return red(v)
	default:
		return blue2(v)
	}
}
