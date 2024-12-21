package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"kotakemail.id/config"
	appcontext "kotakemail.id/pkg/context"
)

type Logger struct {
	cfg    *config.LoggingConfig
	logger zerolog.Logger
}

func NewLogger(ctx *appcontext.AppContext, cfg *config.LoggingConfig) *Logger {
	switch cfg.Level {
	case config.LOG_TRACE:
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case config.LOG_DEBUG:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case config.LOG_INFO:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case config.LOG_WARNING:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case config.LOG_ERROR:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case config.LOG_FATAL:
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	}
	var writers []io.Writer
	if ctx.GetStr(appcontext.EnvironmentKey) == "development" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if cfg.Output == config.LOG_OUTPUT_FILE {
		writers = append(writers, initLogRotator(cfg))
	}

	if len(writers) == 0 {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}

	return &Logger{
		cfg: cfg,
		logger: zerolog.New(zerolog.MultiLevelWriter(writers...)).With().
			Str("service", ctx.GetStr(appcontext.AppNameKey)).
			Str("ENV", ctx.GetStr(appcontext.EnvironmentKey)).
			Timestamp().
			Caller().
			Logger(),
	}
}

func initLogRotator(cfg *config.LoggingConfig) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   cfg.File,
		MaxSize:    cfg.MaxSize,    // Max size in megabytes before log is rotated
		MaxBackups: cfg.MaxBackups, // Max number of old log files to retain
		MaxAge:     cfg.MaxAge,     // Max number of days to retain old log files
		Compress:   cfg.Compress,   // Compress old log files
	}

}
