package logger

import (
	"errors"
	"github.com/DaraDadachanji/switch-context/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var l *Logger

func init() {
	l = New()
}

type Logger struct {
	filename        string
	overRideFileLoc string
	maxSize         int
	maxBackups      int
	maxAge          int
	path            string
	logger          *zap.Logger
}

func New() *Logger {
	l := new(Logger)

	return l
}

func Log(msg string) {
	if l.logger == nil {
		var err error
		if l.logger, err = setup(); err != nil {
			return
		}

	}
	l.logger.Info(msg)
	defer l.logger.Sync()

}

func (l *Logger) SetName(name string) {
	l.filename = name
}

func OverrideFileLoc(file string) {
	l.overrideFileLoc(file)
}

func (l *Logger) overrideFileLoc(file string) {
	l.overRideFileLoc = file
}

func valuesConfig() (string, int, int, int, error) {
	err := config.Path()
	if err != nil {
		return "", 0, 0, 0, err
	}
	filename := config.Get("log.filename")
	if filename == "" {
		return "", 0, 0, 0, errors.New("log.filename not set")
	}
	maxSize := config.GetInt("log.maxSize")
	if maxSize == 0 {
		maxSize = 1
	}
	maxBackups := config.GetInt("log.maxBackups")
	maxAge := config.GetInt("log.maxAge")
	if maxAge == 0 {
		maxAge = 28
	}
	return filename, maxSize, maxBackups, maxAge, nil

}

func setup() (*zap.Logger, error) {
	filename, maxSize, maxBackups, maxAge, err := valuesConfig()
	if err != nil {
		return nil, err
	}
	if l.overRideFileLoc != "" {
		filename = l.overRideFileLoc
	}
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize, // megabytes
		MaxBackups: maxBackups,
		MaxAge:     maxAge, // days
	})
	core := zapcore.NewCore(
		// could have used here zap.NewDevelopmentEncoderConfig() but still not getting line numbers...
		zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			LineEnding: "\n----------\n",
		}),
		w,
		zap.InfoLevel,
	)

	logger := zap.New(core, zap.AddCaller())

	return logger, nil
}
