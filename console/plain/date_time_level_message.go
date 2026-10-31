package plain

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

var DateTimeLevelMessage = zerolog.New(
	zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    true,
		TimeFormat: "2006-01-02 15:04:05",
		PartsOrder: []string{
			zerolog.TimestampFieldName,
			zerolog.LevelFieldName,
			zerolog.MessageFieldName,
		},
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[%-5s]", i))
		},
	},
).Level(zerolog.TraceLevel).With().Timestamp().Logger()
