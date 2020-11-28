package golog

import (
	"time"

	"github.com/goinbox/gomisc"
)

type SimpleFormat struct {
	timePattern string
}

func NewSimpleFormat() *SimpleFormat {
	return &SimpleFormat{
		timePattern: "2006-01-02 15:04:05",
	}
}

func (sf *SimpleFormat) SetTimePattern(tp string) {
	sf.timePattern = tp
}

func (sf *SimpleFormat) Format(level int, msg []byte) []byte {
	logLevel, ok := LogLevel[level]
	if !ok {
		logLevel = []byte("-")
	}

	return gomisc.AppendBytes([]byte("["), logLevel, []byte("]\t["),
		[]byte(time.Now().Format(sf.timePattern)), []byte("]\t"), msg, []byte("\n"))
}
