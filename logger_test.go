package log_test

import (
	"io"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"

	log "github.com/go-mods/zerolog-quick"
	"github.com/go-mods/zerolog-quick/console/colored"
	"github.com/go-mods/zerolog-quick/console/plain"
)

func TestConsolePlainMessage(t *testing.T) {
	log.Logger = plain.Message
	log.Print("Printed message")
	log.Trace().Msg("trace message")
	log.Debug().Msg("debug message")
	log.Info().Msg("info message")
	log.Warn().Msg("warn message")
	log.Error().Msg("error message")
	log.WithLevel(zerolog.FatalLevel).Msg("fatal message")
	log.WithLevel(zerolog.PanicLevel).Msg("panic message")
}

func TestConsolePlainDateMessage(t *testing.T) {
	log.Logger = plain.DateMessage
	log.Print("Printed message")
	log.Trace().Msg("trace message")
	log.Debug().Msg("debug message")
	log.Info().Msg("info message")
	log.Warn().Msg("warn message")
	log.Error().Msg("error message")
	log.WithLevel(zerolog.FatalLevel).Msg("fatal message")
	log.WithLevel(zerolog.PanicLevel).Msg("panic message")
}

func TestConsolePlainDateTimeMessage(t *testing.T) {
	log.Logger = plain.DateTimeMessage
	log.Print("Printed message")
	log.Trace().Msg("trace message")
	log.Debug().Msg("debug message")
	log.Info().Msg("info message")
	log.Warn().Msg("warn message")
	log.Error().Msg("error message")
	log.WithLevel(zerolog.FatalLevel).Msg("fatal message")
	log.WithLevel(zerolog.PanicLevel).Msg("panic message")
}

func TestConsolePlainDateLevelMessage(t *testing.T) {
	log.Logger = plain.DateLevelMessage
	log.Print("Printed message")
	log.Trace().Msg("trace message")
	log.Debug().Msg("debug message")
	log.Info().Msg("info message")
	log.Warn().Msg("warn message")
	log.Error().Msg("error message")
	log.WithLevel(zerolog.FatalLevel).Msg("fatal message")
	log.WithLevel(zerolog.PanicLevel).Msg("panic message")
}

func TestConsolePlainDateTimeLevelMessage(t *testing.T) {
	log.Logger = plain.DateTimeLevelMessage
	log.Print("Printed message")
	log.Trace().Msg("trace message")
	log.Debug().Msg("debug message")
	log.Info().Msg("info message")
	log.Warn().Msg("warn message")
	log.Error().Msg("error message")
	log.WithLevel(zerolog.FatalLevel).Msg("fatal message")
	log.WithLevel(zerolog.PanicLevel).Msg("panic message")
}

func TestConsoleColoredMessage(t *testing.T) {
	log.Logger = colored.Message
	log.Print("Printed message")
	log.Trace().Msg("trace message")
	log.Debug().Msg("debug message")
	log.Info().Msg("info message")
	log.Warn().Msg("warn message")
	log.Error().Msg("error message")
	log.WithLevel(zerolog.FatalLevel).Msg("fatal message")
	log.WithLevel(zerolog.PanicLevel).Msg("panic message")
}

func TestConsoleColoredDateMessage(t *testing.T) {
	log.Logger = colored.DateMessage
	log.Print("Printed message")
	log.Trace().Msg("trace message")
	log.Debug().Msg("debug message")
	log.Info().Msg("info message")
	log.Warn().Msg("warn message")
	log.Error().Msg("error message")
	log.WithLevel(zerolog.FatalLevel).Msg("fatal message")
	log.WithLevel(zerolog.PanicLevel).Msg("panic message")
}

func TestConsoleColoredDateTimeMessage(t *testing.T) {
	log.Logger = colored.DateTimeMessage
	log.Print("Printed message")
	log.Trace().Msg("trace message")
	log.Debug().Msg("debug message")
	log.Info().Msg("info message")
	log.Warn().Msg("warn message")
	log.Error().Msg("error message")
	log.WithLevel(zerolog.FatalLevel).Msg("fatal message")
	log.WithLevel(zerolog.PanicLevel).Msg("panic message")
}

func TestConsoleColoredDateLevelMessage(t *testing.T) {
	log.Logger = colored.DateLevelMessage
	log.Print("Printed message")
	log.Trace().Msg("trace message")
	log.Debug().Msg("debug message")
	log.Info().Msg("info message")
	log.Warn().Msg("warn message")
	log.Error().Msg("error message")
	log.WithLevel(zerolog.FatalLevel).Msg("fatal message")
	log.WithLevel(zerolog.PanicLevel).Msg("panic message")
}

func TestConsoleColoredDateTimeLevelMessage(t *testing.T) {
	log.Logger = colored.DateTimeLevelMessage
	log.Print("Printed message")
	log.Trace().Msg("trace message")
	log.Debug().Msg("debug message")
	log.Info().Msg("info message")
	log.Warn().Msg("warn message")
	log.Error().Msg("error message")
	log.WithLevel(zerolog.FatalLevel).Msg("fatal message")
	log.WithLevel(zerolog.PanicLevel).Msg("panic message")
}

func TestJsonDefault(t *testing.T) {
	log.Info().Msg("Info message")
}

func TestMulti(t *testing.T) {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = time.RFC3339Nano

	var consoleOut io.Writer = zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    true,
		TimeFormat: "2006-01-02",
		PartsOrder: []string{
			zerolog.TimestampFieldName,
			zerolog.MessageFieldName,
		},
		FormatExtra: colored.Colorize,
	}

	fileOut := &lumberjack.Logger{
		Filename:   "test.log",
		MaxSize:    5, //
		MaxBackups: 10,
		MaxAge:     14,
		Compress:   true,
	}

	multi := zerolog.MultiLevelWriter(consoleOut, fileOut)

	logger := zerolog.New(multi).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	logger.Print("Printed message")
	logger.Trace().Msg("trace message")
	logger.Debug().Msg("debug message")
	logger.Info().Msg("info message")
	logger.Warn().Msg("warn message")
	logger.Error().Msg("error message")
	logger.WithLevel(zerolog.FatalLevel).Msg("fatal message")
	logger.WithLevel(zerolog.PanicLevel).Msg("panic message")
}
