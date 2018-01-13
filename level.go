package log

const (
	TRACE = uint8(1 << iota) //1
	DEBUG                    //2
	INFO                     //4
	WARN                     //8
	ERROR                    //16
	FATAL                    //32
	PANIC                    //64
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
