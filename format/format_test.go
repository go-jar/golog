package format

import (
	"fmt"
	"golog/base"
	"testing"
)

func TestConsoleFormat(t *testing.T) {
	sf := NewSimpleFormat()
	cf := NewConsoleFormat(sf)
	fmt.Println(string(cf.Format(base.LEVEL_INFO, []byte("Hello, World!"))))
}
