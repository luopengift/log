package log

const (
	// DEBUG level debug
	DEBUG = uint8(1 << iota)
	TRACE
	INFO
	WARN
	ERROR
	FATAL
	PANIC
)

// LevelMap
var LevelMap = map[uint8]string{
	DEBUG: "D",
	TRACE: "T",
	INFO:  "I",
	WARN:  "W",
	ERROR: "E",
	FATAL: "F",
	PANIC: "P",
}
