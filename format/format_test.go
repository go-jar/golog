package format

import (
	"fmt"
	"testing"

	"github.com/go-jar/golog"
)

func TestConsoleFormat(t *testing.T) {
	sf := NewSimpleFormat()
	cf := NewConsoleFormat(sf)
	fmt.Println(string(cf.Format(golog.LEVEL_INFO, []byte("Hello, World!"))))
}
