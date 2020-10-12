package format

import (
	"golog/base"

	"github.com/go-jar/color"
)

type colorFunc func(msg []byte) []byte

type ConsoleFormat struct {
	formatter base.IFormat
	colorMsg  map[int]colorFunc
}

func NewConsoleFormat(format base.IFormat) *ConsoleFormat {
	return &ConsoleFormat{
		formatter: format,
		colorMsg: map[int]colorFunc{
			base.LEVEL_DEBUG:     color.RenderMsgYellow,
			base.LEVEL_INFO:      color.RenderMsgBlue,
			base.LEVEL_NOTICE:    color.RenderMsgCyan,
			base.LEVEL_WARN:      color.RenderMsgMaganta,
			base.LEVEL_ERROR:     color.RenderMsgRed,
			base.LEVEL_CRITICAL:  color.RenderMsgBlack,
			base.LEVEL_ALERT:     color.RenderMsgWhite,
			base.LEVEL_EMERGENCY: color.RenderMsgGreen,
		},
	}
}

func (cf *ConsoleFormat) SetColorFunc(level int, colorF colorFunc) {
	cf.colorMsg[level] = colorF
}

func (cf *ConsoleFormat) Format(level int, msg []byte) []byte {
	return cf.colorMsg[level](cf.formatter.Format(level, msg))
}
