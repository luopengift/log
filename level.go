package log

const (
	TRACE = uint8(1 << iota) //1
	DEBUG                    //2
	INFO                     //4
	WARN                     //8
	ERROR                    //16
	FATAL                    //32
	PANIC                    //64
	OFF                      //128
	NULL  = uint8(0)
)
