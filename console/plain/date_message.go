package plain

import (
	"os"

	"github.com/rs/zerolog"
)

var DateMessage = zerolog.New(
	zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    true,
		TimeFormat: "2006-01-02",
		PartsOrder: []string{
			zerolog.TimestampFieldName,
			zerolog.MessageFieldName,
		},
	},
).With().Timestamp().Logger()
