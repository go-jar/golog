package golog

import (
	"fmt"
	"testing"
)

func TestConsoleFormat(t *testing.T) {
	sf := NewSimpleFormat()
	cf := NewConsoleFormat(sf)
	fmt.Println(string(cf.Format(LevelInfo, []byte("Hello, World!"))))

	fif := NewFileInfoFormat(0)
	cf = NewConsoleFormat(fif)
	fmt.Println(string(cf.Format(LevelDebug, []byte("Hello, World!"))))
}
