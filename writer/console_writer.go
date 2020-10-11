package writer

import (
	"os"
	"sync"
)

type ConsoleWriter struct {
	lock *sync.Mutex
}

func NewConsoleWriter() *ConsoleWriter {
	return &ConsoleWriter{
		lock: new(sync.Mutex),
	}
}

func (cw *ConsoleWriter) Write(msg []byte) (int, error) {
	cw.lock.Lock()
	defer cw.lock.Unlock()

	return os.Stdout.Write(msg)
}

func (cw *ConsoleWriter) Flush() error {
	return nil
}

func (cw *ConsoleWriter) Free() {
}
