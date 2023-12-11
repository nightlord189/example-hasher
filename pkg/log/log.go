package log

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"time"
)

func InitLogger(level, component string) error {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.LevelFieldName = "log_level" // to compatibility with Graylog

	levelParsed, err := zerolog.ParseLevel(level)
	if err != nil {
		return fmt.Errorf("err parse log level %s: %w", level, err)
	}
	zerolog.SetGlobalLevel(levelParsed)
	logger := zerolog.New(os.Stdout).With().Timestamp().
		Str("component", component).
		Logger()

	zerolog.DefaultContextLogger = &logger
	return nil
}
