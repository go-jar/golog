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
			LevelDebug:     color.RenderMsgYellow,
			LevelInfo:      color.RenderMsgBlue,
			LevelNotice:    color.RenderMsgCyan,
			LevelWarn:      color.RenderMsgMaganta,
			LevelError:     color.RenderMsgRed,
			LevelCritical:  color.RenderMsgBlack,
			LevelAlert:     color.RenderMsgWhite,
			LevelEmergency: color.RenderMsgGreen,
		},
	}
}

func (cf *ConsoleFormat) SetColorFunc(level int, colorF colorFunc) {
	cf.colorMsg[level] = colorF
}

func (cf *ConsoleFormat) Format(level int, msg []byte) []byte {
	return cf.colorMsg[level](cf.formatter.Format(level, msg))
}
