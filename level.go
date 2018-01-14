package log

const (
	DEBUG = uint8(1 << iota)
	INFO
	WARN
	ERROR
	FATAL
	PANIC
	TRACE
)

var LevelMap = map[uint8]string{
	TRACE: "T",
	DEBUG: "D",
	INFO:  "I",
	WARN:  "W",
	ERROR: "E",
	FATAL: "F",
	PANIC: "P",
}

