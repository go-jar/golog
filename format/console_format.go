package format

import (
	"golog/base"

	"github.com/go-jar/gocolor"
)

type colorFunc func(msg []byte) []byte

type ConsoleFormat struct {
	f        base.IFormat
	colorMsg map[int]colorFunc
}

func NewConsoleFormat(format base.IFormat) *ConsoleFormat {
	return &ConsoleFormat{
		f: format,
		colorMsg: map[int]colorFunc{
			base.LEVEL_DEBUG:     gocolor.GetYellowMsg,
			base.LEVEL_INFO:      gocolor.GetBlueMsg,
			base.LEVEL_NOTICE:    gocolor.GetCyanMsg,
			base.LEVEL_WARN:      gocolor.GetMagantaMsg,
			base.LEVEL_ERROR:     gocolor.GetRedMsg,
			base.LEVEL_CRITICAL:  gocolor.GetBlackMsg,
			base.LEVEL_ALERT:     gocolor.GetWhiteMsg,
			base.LEVEL_EMERGENCY: gocolor.GetGreenMsg,
		},
	}
}

func (cf *ConsoleFormat) SetColorFunc(level int, colorF colorFunc) {
	cf.colorMsg[level] = colorF
}

func (cf *ConsoleFormat) Format(level int, msg []byte) []byte {
	return cf.colorMsg[level](cf.f.Format(level, msg))
}
