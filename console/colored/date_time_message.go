package colored

import (
	"os"

	"github.com/rs/zerolog"
)

var DateTimeMessage = zerolog.New(
	zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    true,
		TimeFormat: "2006-01-02 15:04:05",
		PartsOrder: []string{
			zerolog.TimestampFieldName,
			zerolog.MessageFieldName,
		},
		FormatExtra: Colorize,
	},
).With().Timestamp().Logger()
