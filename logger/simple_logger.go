package logger

import (
	"golog/base"
	"golog/format"
	"golog/writer"
)

type SimpleLogger struct {
	W        base.IWriter
	formater base.IFormat
	level    int
}

func NewSimpleLogger(path string, level int) (*SimpleLogger, error) {
	fw, err := writer.NewFileWriter(path, 1024)
	if err != nil {
		return nil, err
	}

	return &SimpleLogger{
		W:        writer.NewAsyncWriter(fw, 1024),
		formater: format.NewSimpleFormat(),
		level:    level,
	}, nil
}

func (sl *SimpleLogger) SetWriter(w base.IWriter) {
	sl.W = w
}

func (sl *SimpleLogger) SetFormat(f base.IFormat) {
	sl.formater = f
}

func (sl *SimpleLogger) SetLevel(level int) {
	sl.level = level
}

func (sl *SimpleLogger) Emergency(msg []byte) {
	sl.Log(base.LEVEL_EMERGENCY, msg)
}

func (sl *SimpleLogger) Alert(msg []byte) {
	sl.Log(base.LEVEL_ALERT, msg)
}

func (sl *SimpleLogger) Critical(msg []byte) {
	sl.Log(base.LEVEL_CRITICAL, msg)
}

func (sl *SimpleLogger) Error(msg []byte) {
	sl.Log(base.LEVEL_ERROR, msg)
}

func (sl *SimpleLogger) Warn(msg []byte) {
	sl.Log(base.LEVEL_WARN, msg)
}

func (sl *SimpleLogger) Notice(msg []byte) {
	sl.Log(base.LEVEL_NOTICE, msg)
}

func (sl *SimpleLogger) Info(msg []byte) {
	sl.Log(base.LEVEL_INFO, msg)
}

func (sl *SimpleLogger) Debug(msg []byte) {
	sl.Log(base.LEVEL_DEBUG, msg)
}

func (sl *SimpleLogger) Log(level int, msg []byte) error {
	if level > sl.level {
		return nil
	}

	if _, err := sl.W.Write(sl.formater.Format(level, msg)); err != nil {
		return err
	}

	return nil
}
