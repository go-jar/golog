package golog

type SimpleLogger struct {
	w         IWriter
	formatter IFormat
	level     int
}

func NewSimpleLogger(w IWriter, f IFormat) *SimpleLogger {
	return &SimpleLogger{
		w:         w,
		formatter: f,
		level:     LevelInfo,
	}
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

func (sl *SimpleLogger) SetWriter(w IWriter) *SimpleLogger {
	if sl.w != nil {
		sl.w.Free()
	}

	sl.w = w
	return sl
}

func (sl *SimpleLogger) SetFormat(f IFormat) *SimpleLogger {
	sl.formatter = f
	return sl
}

func (sl *SimpleLogger) SetLevel(level int) *SimpleLogger {
	sl.level = level
	return sl
}

func (sl *SimpleLogger) Emergency(msg []byte) {
	sl.Log(LevelEmergency, msg)
}

func (sl *SimpleLogger) Alert(msg []byte) {
	sl.Log(LevelAlert, msg)
}

func (sl *SimpleLogger) Critical(msg []byte) {
	sl.Log(LevelCritical, msg)
}

func (sl *SimpleLogger) Error(msg []byte) {
	sl.Log(LevelError, msg)
}

func (sl *SimpleLogger) Warn(msg []byte) {
	sl.Log(LevelWarn, msg)
}

func (sl *SimpleLogger) Notice(msg []byte) {
	sl.Log(LevelNotice, msg)
}

func (sl *SimpleLogger) Info(msg []byte) {
	sl.Log(LevelInfo, msg)
}

func (sl *SimpleLogger) Debug(msg []byte) {
	sl.Log(LevelDebug, msg)
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
