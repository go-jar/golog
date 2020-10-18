package golog

import (
	"fmt"
	"testing"
)

func TestSimpleLogger(t *testing.T) {
	sl, err := NewFileLogger("/data/test.log", 1024, LEVEL_INFO)
	if err != nil {
		fmt.Println(err)
	}

	testLog(sl)

	sl, err = NewAsyncLogger("/data/test1.log", 1024, 1024, LEVEL_ERROR)
	if err != nil {
		fmt.Println(err)
	}

	testLog(sl)

	sl, err = NewConsoleLogger(LEVEL_NOTICE)
	if err != nil {
		fmt.Println(err)
	}

	testLog(sl)
}

func testLog(sl *SimpleLogger) {
	msg := []byte("Hello, World!")

	for i := 0; i < 1000; i++ {
		sl.Emergency(msg)
		sl.Alert(msg)
		sl.Critical(msg)
		sl.Error(msg)
		sl.Warn(msg)
		sl.Notice(msg)
		sl.Info(msg)
		sl.Debug(msg)
	}

	sl.Close()
}
