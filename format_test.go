package golog

import (
	"fmt"
	"testing"
)

func TestConsoleFormat(t *testing.T) {
	sf := NewSimpleFormat()
	cf := NewConsoleFormat(sf)
	fmt.Println(string(cf.Format(LEVEL_INFO, []byte("Hello, World!"))))
}
