package golog

import "io"

const (
	LevelEmergency = 0
	LevelAlert     = 1
	LevelCritical  = 2
	LevelError     = 3
	LevelWarn      = 4
	LevelNotice    = 5
	LevelInfo      = 6
	LevelDebug     = 7
)

var LogLevel map[int][]byte = map[int][]byte{
	LevelEmergency: []byte("Emerg"),
	LevelAlert:     []byte("Alert"),
	LevelCritical:  []byte("Crit"),
	LevelError:     []byte("Error"),
	LevelWarn:      []byte("Warn"),
	LevelNotice:    []byte("Noti"),
	LevelInfo:      []byte("Info"),
	LevelDebug:     []byte("Debug"),
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
	Log(level int, msg []byte) error
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
