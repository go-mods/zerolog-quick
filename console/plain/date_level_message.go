package plain

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

// DateLevelMessage is a logger that writes the date, the level and the plain message to os.Stdout.
var DateLevelMessage = NewDateLevelMessage(os.Stdout)

// NewDateLevelMessage returns a logger formatted like DateLevelMessage that writes to w.
// A nil w discards the output.
func NewDateLevelMessage(w io.Writer) zerolog.Logger {
	return zerolog.New(
		zerolog.ConsoleWriter{
			Out:        output(w),
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
		},
	).Level(zerolog.TraceLevel).With().Timestamp().Logger()
}
