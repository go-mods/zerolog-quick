package colored

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

var DateLevelMessage = zerolog.New(
	zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    true,
		TimeFormat: "2006-01-02",
		PartsOrder: []string{
			zerolog.TimestampFieldName,
			zerolog.LevelFieldName,
			zerolog.MessageFieldName,
		},
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[%-5s]", i))
		},
		FormatExtra: Colorize,
	},
).Level(zerolog.TraceLevel).With().Timestamp().Logger()
