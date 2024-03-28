package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

const logPath = "./data/blog.log"

var logfile = &lumberjack.Logger{
	MaxSize:    8,
	MaxBackups: 5,
	MaxAge:     30,
	Compress:   true,
}

func init() {
	logfile.Filename = logPath
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(io.MultiWriter(logfile, zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}))
}

func Stop() {
	err := logfile.Close()
	if err != nil {
		log.Error().Err(err).Msg("[logger-Stop-1] could not close logfile")
	}
}

func Rotate() {
	err := logfile.Rotate()
	if err != nil {
		log.Error().Err(err).Msg("[logger-Rotate-1] could not rotate logfile")
	}
}
