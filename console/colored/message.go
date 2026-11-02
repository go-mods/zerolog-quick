package colored

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

// Message is a logger that writes colored messages to os.Stdout.
var Message = NewMessage(os.Stdout)

// NewMessage returns a logger formatted like Message that writes to w.
// A nil w discards the output.
func NewMessage(w io.Writer) zerolog.Logger {
	return zerolog.New(
		zerolog.ConsoleWriter{
			Out:     output(w),
			NoColor: true,
			PartsOrder: []string{
				zerolog.MessageFieldName,
			},
			FormatExtra: Colorize,
		},
	)
}
