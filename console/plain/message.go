package plain

import (
	"os"

	"github.com/rs/zerolog"
)

var Message = zerolog.New(
	zerolog.ConsoleWriter{
		Out:     os.Stdout,
		NoColor: true,
		PartsOrder: []string{
			zerolog.MessageFieldName,
		},
	},
)
