package colored

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

// DateMessage is a logger that writes the date and the colored message to os.Stdout.
var DateMessage = NewDateMessage(os.Stdout)

// NewDateMessage returns a logger formatted like DateMessage that writes to w.
// w must not be nil.
func NewDateMessage(w io.Writer) zerolog.Logger {
	return zerolog.New(
		zerolog.ConsoleWriter{
			Out:        w,
			NoColor:    true,
			TimeFormat: "2006-01-02",
			PartsOrder: []string{
				zerolog.TimestampFieldName,
				zerolog.MessageFieldName,
			},
			FormatExtra: Colorize,
		},
	).With().Timestamp().Logger()
}
