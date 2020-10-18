package golog

import (
	"github.com/go-jar/color"
)

type colorFunc func(msg []byte) []byte

type ConsoleFormat struct {
	formatter IFormat
	colorMsg  map[int]colorFunc
}

func NewConsoleFormat(format IFormat) *ConsoleFormat {
	return &ConsoleFormat{
		formatter: format,
		colorMsg: map[int]colorFunc{
			LEVEL_DEBUG:     color.RenderMsgYellow,
			LEVEL_INFO:      color.RenderMsgBlue,
			LEVEL_NOTICE:    color.RenderMsgCyan,
			LEVEL_WARN:      color.RenderMsgMaganta,
			LEVEL_ERROR:     color.RenderMsgRed,
			LEVEL_CRITICAL:  color.RenderMsgBlack,
			LEVEL_ALERT:     color.RenderMsgWhite,
			LEVEL_EMERGENCY: color.RenderMsgGreen,
		},
	}
}

func (cf *ConsoleFormat) SetColorFunc(level int, colorF colorFunc) {
	cf.colorMsg[level] = colorF
}

func (cf *ConsoleFormat) Format(level int, msg []byte) []byte {
	return cf.colorMsg[level](cf.formatter.Format(level, msg))
}
