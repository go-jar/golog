package golog

import (
	"fmt"
	"strconv"
	"testing"
)

func TestConsoleWriter(t *testing.T) {
	cw := NewConsoleWriter()

	for i := 0; i < 1000; i++ {
		cw.Write([]byte("Hello, " + strconv.Itoa(i) + "\n"))
	}

	cw.Free()
}

func TestFileWriter(t *testing.T) {
	fw, err := NewFileWriter("/tmp/test.log", 1024)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 1000; i++ {
		fw.Write([]byte("Hello, " + strconv.Itoa(i) + "\n"))
	}

	fw.Free()
}

func TestAsyncWriter(t *testing.T) {
	fw, err := NewFileWriter("/tmp/test1.log", 1024)
	if err != nil {
		fmt.Println(err)
	}

	aw := NewAsyncWriter(fw, 10)

	for i := 0; i < 1000; i++ {
		aw.Write([]byte("Hello, " + strconv.Itoa(i) + "\n"))
	}

	aw.Free()
}
