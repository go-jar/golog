package golog

import (
	"log"
	"os"
	"sync"
)

type FileWriter struct {
	File    *os.File
	path    string
	lock    *sync.Mutex
	buf     []byte
	bufSize int
	bufPos  int
}

func NewFileWriter(path string, bufSize int) (*FileWriter, error) {
	f, err := openOrCreateFile(path)
	if err != nil {
		log.Printf("Open file error: %v\n", err)
		return nil, err
	}

	return &FileWriter{
		File:    f,
		path:    path,
		lock:    new(sync.Mutex),
		buf:     make([]byte, bufSize),
		bufSize: bufSize,
		bufPos:  0,
	}, nil
}

func (fw *FileWriter) Write(msg []byte) (int, error) {
	fw.lock.Lock()
	defer fw.lock.Unlock()

	if err := fw.ensureFileExist(); err != nil {
		return 0, err
	}

	if fw.bufSize == 0 {
		return fw.File.Write(msg)
	}

	msgLen := len(msg)
	if msgLen > fw.bufSize {
		if err := fw.flushBuf(); err != nil {
			return 0, err
		}
		return fw.File.Write(msg)
	}

	nextPos := fw.bufPos + msgLen
	if nextPos > fw.bufSize {
		if err := fw.flushBuf(); err != nil {
			return 0, err
		}
	}

	copy(fw.buf[fw.bufPos:], msg)
	fw.bufPos = fw.bufPos + msgLen

	return msgLen, nil
}

func (fw *FileWriter) Flush() error {
	if fw.bufPos == 0 || fw.bufSize == 0 {
		return nil
	}

	fw.lock.Lock()
	defer fw.lock.Unlock()

	if err := fw.ensureFileExist(); err != nil {
		return err
	}

	return fw.flushBuf()
}

func (fw *FileWriter) flushBuf() error {
	if fw.bufPos == 0 {
		return nil
	}

	_, err := fw.File.Write(fw.buf[:fw.bufPos])
	fw.bufPos = 0
	return err
}

func (fw *FileWriter) ensureFileExist() error {
	if fw.isFileExists(fw.path) {
		return nil
	}

	f, err := openOrCreateFile(fw.path)
	if err != nil {
		return err
	}

	fw.File = f
	return nil
}

func (fw *FileWriter) isFileExists(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}
	return true
}

func openOrCreateFile(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fw *FileWriter) Free() {
	fw.ensureFileExist()
	fw.flushBuf()
	fw.File.Close()
}
