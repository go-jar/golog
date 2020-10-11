package logger

import (
	"fmt"
	"golog/base"
	"golog/format"
	"golog/writer"
	"testing"
)

func TestSimpleLogger(t *testing.T) {
	sl, err := NewSimpleLogger("/data/test.log", base.LEVEL_INFO)
	if err != nil {
		fmt.Println(err)
	}

	msg := []byte("Hello, World!")
	for i := 0; i < 10; i++ {
		sl.Emergency(msg)
		sl.Alert(msg)
		sl.Critical(msg)
		sl.Error(msg)
		sl.Warn(msg)
		sl.Notice(msg)
		sl.Info(msg)
		sl.Debug(msg)
	}

	sl.W.Free()

	sl.SetWriter(writer.NewConsoleWriter())
	sl.SetFormat(format.NewConsoleFormat(sl.formater))
	sl.SetLevel(base.LEVEL_DEBUG)

	for i := 0; i < 10; i++ {
		sl.Emergency(msg)
		sl.Alert(msg)
		sl.Critical(msg)
		sl.Error(msg)
		sl.Warn(msg)
		sl.Notice(msg)
		sl.Info(msg)
		sl.Debug(msg)
	}
}
