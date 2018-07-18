package log

// DEBUG level debug
const (
	DEBUG = uint8(1 << iota)
	TRACE
	INFO
	WARN
	ERROR
	FATAL
	PANIC
)

// LevelMap level map
var LevelMap = map[uint8]string{
	DEBUG: "D",
	TRACE: "T",
	INFO:  "I",
	WARN:  "W",
	ERROR: "E",
	FATAL: "F",
	PANIC: "P",
}

// LevelColor Level color
var LevelColor = map[uint8]uint8{
	DEBUG: blue,
	TRACE: green,
	INFO:  none,
	WARN:  yellow,
	ERROR: magenta,
	FATAL: red,
	PANIC: blue2,
}
