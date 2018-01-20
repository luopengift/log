package log

import (
	"fmt"
)

const (
	_RED     = uint8(iota + 91) //红色
	_GREEN                      //绿色
	_YELLOW                     //黄色
	_BLUE                       //蓝色
	_MAGENTA                    //洋红
	_BLUE2                      //湖蓝
)

func color(col uint8, s interface{}) string {
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", col, s)
}

func none(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func red(v interface{}) string {
	return color(_RED, v)
}

func green(v interface{}) string {
	return color(_GREEN, v)
}

func yellow(v interface{}) string {
	return color(_YELLOW, v)
}

func blue(v interface{}) string {
	return color(_BLUE, v)
}

func magenta(v interface{}) string {
	return color(_MAGENTA, v)
}

func blue2(v interface{}) string {
	return color(_BLUE2, v)
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
