package golog

import (
	"path"
	"runtime"
	"strconv"
	"time"

	"github.com/goinbox/gomisc"
)

type FileInfoFormat struct {
	timePattern string
}

func NewFileInfoFormat() *FileInfoFormat {
	return &FileInfoFormat{
		timePattern: "2006-01-02 15:04:05",
	}
}

func (sf *FileInfoFormat) SetTimePattern(tp string) {
	sf.timePattern = tp
}

func (sf *FileInfoFormat) Format(level int, msg []byte) []byte {
	logLevel, ok := LogLevel[level]
	if !ok {
		logLevel = []byte("-")
	}

	_, fullFileName, line, _ := runtime.Caller(3)
	fileName := path.Base(fullFileName)

	return gomisc.AppendBytes([]byte("["), logLevel, []byte("]\t["),
		[]byte(time.Now().Format(sf.timePattern)), []byte("]\t["),
		[]byte(fileName), []byte(": "), []byte(strconv.Itoa(line)), []byte("] "),
		msg, []byte("\n"))
}
