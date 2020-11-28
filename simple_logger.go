package golog

type SimpleLogger struct {
	w         IWriter
	formatter IFormat
	level     int
}

func NewSimpleLogger(w IWriter, f IFormat) (*SimpleLogger, error) {
	return &SimpleLogger{
		w:         w,
		formatter: f,
		level:     LEVEL_INFO,
	}, nil
}

func NewFileLogger(path string, bufSize, level int) (*SimpleLogger, error) {
	fw, err := NewFileWriter(path, bufSize)
	if err != nil {
		return nil, err
	}

	return &SimpleLogger{
		w:         fw,
		formatter: NewSimpleFormat(),
		level:     level,
	}, nil
}

func NewAsyncLogger(path string, bufSize, queueSize, level int) (*SimpleLogger, error) {
	fw, err := NewFileWriter(path, bufSize)
	if err != nil {
		return nil, err
	}

	return &SimpleLogger{
		w:         NewAsyncWriter(fw, queueSize),
		formatter: NewSimpleFormat(),
		level:     level,
	}, nil
}

func NewConsoleLogger(level int) (*SimpleLogger, error) {
	return &SimpleLogger{
		w:         NewConsoleWriter(),
		formatter: NewConsoleFormat(NewSimpleFormat()),
		level:     level,
	}, nil
}

func (sl *SimpleLogger) SetWriter(w IWriter) {
	if sl.w != nil {
		sl.w.Free()
	}

	sl.w = w
}

func (sl *SimpleLogger) SetFormat(f IFormat) {
	sl.formatter = f
}

func (sl *SimpleLogger) SetLevel(level int) {
	sl.level = level
}

func (sl *SimpleLogger) Emergency(msg []byte) {
	sl.Log(LEVEL_EMERGENCY, msg)
}

func (sl *SimpleLogger) Alert(msg []byte) {
	sl.Log(LEVEL_ALERT, msg)
}

func (sl *SimpleLogger) Critical(msg []byte) {
	sl.Log(LEVEL_CRITICAL, msg)
}

func (sl *SimpleLogger) Error(msg []byte) {
	sl.Log(LEVEL_ERROR, msg)
}

func (sl *SimpleLogger) Warn(msg []byte) {
	sl.Log(LEVEL_WARN, msg)
}

func (sl *SimpleLogger) Notice(msg []byte) {
	sl.Log(LEVEL_NOTICE, msg)
}

func (sl *SimpleLogger) Info(msg []byte) {
	sl.Log(LEVEL_INFO, msg)
}

func (sl *SimpleLogger) Debug(msg []byte) {
	sl.Log(LEVEL_DEBUG, msg)
}

func (sl *SimpleLogger) Close() {
	sl.w.Free()
}

func (sl *SimpleLogger) Log(level int, msg []byte) error {
	if level > sl.level {
		return nil
	}

	if _, err := sl.w.Write(sl.formatter.Format(level, msg)); err != nil {
		return err
	}

	return nil
}
