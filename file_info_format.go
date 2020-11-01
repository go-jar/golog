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
	callLevel   int
}

func NewFileInfoFormat() *FileInfoFormat {
	return &FileInfoFormat{
		timePattern: "2006-01-02 15:04:05",
	}
}

func (fif *FileInfoFormat) SetCallLevel(callLevel int) {
	fif.callLevel = callLevel
}

func (fif *FileInfoFormat) SetTimePattern(tp string) {
	fif.timePattern = tp
}

func (fif *FileInfoFormat) Format(level int, msg []byte) []byte {
	logLevel, ok := LogLevel[level]
	if !ok {
		logLevel = []byte("-")
	}

	_, fullFileName, line, _ := runtime.Caller(3 + fif.callLevel)
	fileName := path.Base(fullFileName)

	return gomisc.AppendBytes([]byte("["), logLevel, []byte("]\t["),
		[]byte(time.Now().Format(fif.timePattern)), []byte("]\t["),
		[]byte(fileName), []byte(": "), []byte(strconv.Itoa(line)), []byte("] "),
		msg, []byte("\n"))
}
