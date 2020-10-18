package format

import (
	"github.com/go-jar/color"
	"github.com/go-jar/golog"
)

type colorFunc func(msg []byte) []byte

type ConsoleFormat struct {
	formatter golog.IFormat
	colorMsg  map[int]colorFunc
}

func NewConsoleFormat(format golog.IFormat) *ConsoleFormat {
	return &ConsoleFormat{
		formatter: format,
		colorMsg: map[int]colorFunc{
			golog.LEVEL_DEBUG:     color.RenderMsgYellow,
			golog.LEVEL_INFO:      color.RenderMsgBlue,
			golog.LEVEL_NOTICE:    color.RenderMsgCyan,
			golog.LEVEL_WARN:      color.RenderMsgMaganta,
			golog.LEVEL_ERROR:     color.RenderMsgRed,
			golog.LEVEL_CRITICAL:  color.RenderMsgBlack,
			golog.LEVEL_ALERT:     color.RenderMsgWhite,
			golog.LEVEL_EMERGENCY: color.RenderMsgGreen,
		},
	}
}

func (cf *ConsoleFormat) SetColorFunc(level int, colorF colorFunc) {
	cf.colorMsg[level] = colorF
}

func (cf *ConsoleFormat) Format(level int, msg []byte) []byte {
	return cf.colorMsg[level](cf.formatter.Format(level, msg))
}
