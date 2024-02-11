package log

import (
	"os"

	zlog "github.com/rs/zerolog"
	logger "github.com/rs/zerolog/log"
)

const (
	Debug = zlog.DebugLevel
	Info  = zlog.InfoLevel
	Warn  = zlog.WarnLevel
	Error = zlog.ErrorLevel
	Fatal = zlog.FatalLevel
	Panic = zlog.PanicLevel
)

type LogContext map[string]interface{}

func Log(
	msg string,
	level zlog.Level,
	context map[string]interface{},
) {
	if context == nil {
		logger.WithLevel(level).Msg(msg)
	} else {
		logger.WithLevel(level).Fields(context).Msg(msg)
	}
}

func Init() {
	zlog.TimeFieldFormat = zlog.TimeFormatUnix
	zlog.SetGlobalLevel(zlog.InfoLevel)
	logger.Logger = logger.Output(zlog.ConsoleWriter{Out: os.Stderr})
}
