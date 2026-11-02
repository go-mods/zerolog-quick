package plain

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

// Message is a logger that writes plain messages to os.Stdout.
var Message = NewMessage(os.Stdout)

// NewMessage returns a logger formatted like Message that writes to w.
// w must not be nil.
func NewMessage(w io.Writer) zerolog.Logger {
	return zerolog.New(
		zerolog.ConsoleWriter{
			Out:     w,
			NoColor: true,
			PartsOrder: []string{
				zerolog.MessageFieldName,
			},
		},
	)
}
