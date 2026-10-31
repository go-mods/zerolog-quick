package json

import (
	"os"

	"github.com/rs/zerolog"
)

// Default is the default global logger
var Default = zerolog.New(os.Stderr).With().Timestamp().Logger()
