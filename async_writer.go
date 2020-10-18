package golog

import (
	"sync"
)

const (
	WRITE = 0
	FLUSH = 1
	FREE  = 2
)

type AsyncWriter struct {
	w         IWriter
	msgChan   chan *asyncMsg
	queueSize int
	wg        *sync.WaitGroup
}

type asyncMsg struct {
	kind int
	msg  []byte
}

func NewAsyncWriter(w IWriter, queueSize int) *AsyncWriter {
	aw := &AsyncWriter{
		w:         w,
		msgChan:   make(chan *asyncMsg, queueSize),
		queueSize: queueSize,
		wg:        new(sync.WaitGroup),
	}

	go aw.asyncLogRoutine()
	aw.wg.Add(1)
	return aw
}

func (aw *AsyncWriter) asyncLogRoutine() {
	defer aw.wg.Done()

	for {
		am := <-aw.msgChan
		switch am.kind {
		case WRITE:
			aw.w.Write(am.msg)
		case FLUSH:
			aw.w.Flush()
		case FREE:
			aw.w.Free()
			return
		}
	}
}

func (aw *AsyncWriter) Write(msg []byte) (int, error) {
	aw.msgChan <- &asyncMsg{kind: WRITE, msg: msg}
	return len(msg), nil
}

func (aw *AsyncWriter) Flush() error {
	aw.msgChan <- &asyncMsg{kind: FLUSH, msg: nil}
	return nil
}

func (aw *AsyncWriter) Free() {
	aw.msgChan <- &asyncMsg{kind: FREE, msg: nil}
	aw.wg.Wait()
}
