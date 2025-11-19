package utils

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Logger = zerolog.Logger{}

func InitLogger() {
	logger := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.DateTime},
	).Level(Config.LogLevel).With().Timestamp().Caller().Logger()

	Logger = logger

	logger.Debug().Msg("Logger setup complete")
}
