package format

import (
	"bytes"
	"golog/base"
	"time"
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
	logLevel, ok := base.LogLevel[level]
	if !ok {
		logLevel = []byte("-")
	}

	return appendToBuf([]byte("["), logLevel, []byte("]\t["), []byte(time.Now().Format(sf.timePattern)), []byte("]\t"), msg, []byte("\n"))
}

func appendToBuf(elem []byte, elemsRest ...[]byte) []byte {
	buf := bytes.NewBuffer(elem)
	for _, e := range elemsRest {
		buf.Write(e)
	}

	return buf.Bytes()
}
