package main

import (
	"os"

	"github.com/rs/zerolog"

	log "github.com/go-mods/zerolog-quick"
	"github.com/go-mods/zerolog-quick/console/colored"
	"github.com/go-mods/zerolog-quick/console/plain"
)

func main() {
	// Default JSON Logger
	log.Info().Msg("This is a default JSON log message")

	// Plain Console Loggers
	log.Logger = plain.Message
	log.Info().Msg("Plain message only")

	log.Logger = plain.DateMessage
	log.Info().Msg("Plain date and message")

	log.Logger = plain.DateTimeMessage
	log.Info().Msg("Plain date, time, and message")

	log.Logger = plain.DateLevelMessage
	log.Info().Msg("Plain date, level, and message")

	log.Logger = plain.DateTimeLevelMessage
	log.Info().Msg("Plain date, time, level, and message")

	// Colored Console Loggers
	log.Logger = colored.Message
	log.Info().Msg("Colored message only")

	log.Logger = colored.DateMessage
	log.Info().Msg("Colored date and message")

	log.Logger = colored.DateTimeMessage
	log.Info().Msg("Colored date, time, and message")

	log.Logger = colored.DateLevelMessage
	log.Info().Msg("Colored date, level, and message")

	log.Logger = colored.DateTimeLevelMessage
	log.Info().Msg("Colored date, time, level, and message")

	// Custom Logger Configuration
	customLogger := zerolog.New(
		zerolog.ConsoleWriter{
			Out:        os.Stdout,
			NoColor:    false,
			TimeFormat: "2006-01-02 15:04:05",
			PartsOrder: []string{
				zerolog.TimestampFieldName,
				zerolog.LevelFieldName,
				zerolog.MessageFieldName,
			},
		},
	).With().Timestamp().Logger()

	log.Logger = customLogger
	log.Info().Msg("Custom logger message")

	// Demonstrating different log levels
	log.Debug().Msg("This is a debug message")
	log.Info().Msg("This is an info message")
	log.Warn().Msg("This is a warning message")
	log.Error().Msg("This is an error message")
	// log.Fatal().Msg("This is a fatal message") // Uncomment to test, but note that it will terminate the program

	// Logging with additional fields
	log.Info().
		Str("string_field", "value").
		Int("int_field", 123).
		Float64("float_field", 3.14).
		Bool("bool_field", true).
		Msg("Log message with additional fields")

	// Logging an error
	err := someFunction()
	if err != nil {
		log.Error().Err(err).Msg("An error occurred")
	}
}

func someFunction() error {
	// Simulating an error
	return os.ErrNotExist
}
