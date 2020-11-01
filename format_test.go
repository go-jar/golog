package golog

import (
	"fmt"
	"testing"
)

func TestConsoleFormat(t *testing.T) {
	sf := NewSimpleFormat()
	cf := NewConsoleFormat(sf)
	fmt.Println(string(cf.Format(LEVEL_INFO, []byte("Hello, World!"))))

	fif := NewFileInfoFormat()
	cf = NewConsoleFormat(fif)
	fmt.Println(string(cf.Format(LEVEL_DEBUG, []byte("Hello, World!"))))
}
