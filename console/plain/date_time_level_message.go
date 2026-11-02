package plain

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

// DateTimeLevelMessage is a logger that writes the date, the time, the level and the plain message to os.Stdout.
var DateTimeLevelMessage = NewDateTimeLevelMessage(os.Stdout)

// NewDateTimeLevelMessage returns a logger formatted like DateTimeLevelMessage that writes to w.
// w must not be nil.
func NewDateTimeLevelMessage(w io.Writer) zerolog.Logger {
	return zerolog.New(
		zerolog.ConsoleWriter{
			Out:        w,
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
}
