package format

import (
	"golog/base"

	"github.com/go-jar/gocolor"
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
			base.LEVEL_DEBUG:     gocolor.RenderMsgYellow,
			base.LEVEL_INFO:      gocolor.RenderMsgBlue,
			base.LEVEL_NOTICE:    gocolor.RenderMsgCyan,
			base.LEVEL_WARN:      gocolor.RenderMsgMaganta,
			base.LEVEL_ERROR:     gocolor.RenderMsgRed,
			base.LEVEL_CRITICAL:  gocolor.RenderMsgBlack,
			base.LEVEL_ALERT:     gocolor.RenderMsgWhite,
			base.LEVEL_EMERGENCY: gocolor.RenderMsgGreen,
		},
	}
}

func (cf *ConsoleFormat) SetColorFunc(level int, colorF colorFunc) {
	cf.colorMsg[level] = colorF
}

func (cf *ConsoleFormat) Format(level int, msg []byte) []byte {
	return cf.colorMsg[level](cf.formatter.Format(level, msg))
}
