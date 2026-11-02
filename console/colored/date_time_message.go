package colored

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

// DateTimeMessage is a logger that writes the date, the time and the colored message to os.Stdout.
var DateTimeMessage = NewDateTimeMessage(os.Stdout)

// NewDateTimeMessage returns a logger formatted like DateTimeMessage that writes to w.
// w must not be nil.
func NewDateTimeMessage(w io.Writer) zerolog.Logger {
	return zerolog.New(
		zerolog.ConsoleWriter{
			Out:        w,
			NoColor:    true,
			TimeFormat: "2006-01-02 15:04:05",
			PartsOrder: []string{
				zerolog.TimestampFieldName,
				zerolog.MessageFieldName,
			},
			FormatExtra: Colorize,
		},
	).With().Timestamp().Logger()
}
