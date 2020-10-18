package golog

import "io"

const (
	LEVEL_EMERGENCY = 0
	LEVEL_ALERT     = 1
	LEVEL_CRITICAL  = 2
	LEVEL_ERROR     = 3
	LEVEL_WARN      = 4
	LEVEL_NOTICE    = 5
	LEVEL_INFO      = 6
	LEVEL_DEBUG     = 7
)

var LogLevel map[int][]byte = map[int][]byte{
	LEVEL_EMERGENCY: []byte("Emerg"),
	LEVEL_ALERT:     []byte("Alert"),
	LEVEL_CRITICAL:  []byte("Crit"),
	LEVEL_ERROR:     []byte("Error"),
	LEVEL_WARN:      []byte("Warn"),
	LEVEL_NOTICE:    []byte("Noti"),
	LEVEL_INFO:      []byte("Info"),
	LEVEL_DEBUG:     []byte("Debug"),
}

type ILogger interface {
	Emergency(msg []byte)
	Alert(msg []byte)
	Critical(msg []byte)
	Error(msg []byte)
	Warn(msg []byte)
	Notice(msg []byte)
	Info(msg []byte)
	Debug(msg []byte)
	Close()
}

type IFormat interface {
	Format(level int, msg []byte) []byte
}

type IWriter interface {
	io.Writer
	Flush() error
	Free()
}
