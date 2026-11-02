# zerolog-quick

[![Go Reference](https://pkg.go.dev/badge/github.com/go-mods/zerolog-quick.svg)](https://pkg.go.dev/github.com/go-mods/zerolog-quick)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-mods/zerolog-quick)](https://goreportcard.com/report/github.com/go-mods/zerolog-quick)
[![Release](https://img.shields.io/github/release/go-mods/zerolog-quick.svg?style=flat)](https://github.com/go-mods/zerolog-quick/releases)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/go-mods/zerolog-quick/blob/master/LICENSE.md)


`zerolog-quick` is a Go package that provides pre-configured loggers using the [zerolog](https://github.com/rs/zerolog) library. It offers both plain and colored console loggers with various formatting options.

## Installation

```bash
  go get github.com/go-mods/zerolog-quick
```

## Usage

### Default JSON Logger

The default logger is a JSON logger that writes to `os.Stderr`:

```go
import log "github.com/go-mods/zerolog-quick"

func main() {
    log.Info().Msg("This is an info message")
}
```

### Console Loggers

#### Plain Console Loggers

The plain console logger writes to `os.Stdout` and does not use any colors:

```go
import (
    log "github.com/go-mods/zerolog-quick"
    "github.com/go-mods/zerolog-quick/console/plain"
)

func main() {
    // Message only
    log.Logger = plain.Message
    log.Info().Msg("Info message")
    // Date and message
    log.Logger = plain.DateMessage
    log.Info().Msg("Info message with date")
    // Date, time, and message
    log.Logger = plain.DateTimeMessage
    log.Info().Msg("Info message with date and time")
    // Date, level, and message
    log.Logger = plain.DateLevelMessage
    log.Info().Msg("Info message with date and level")
    // Date, time, level, and message
    log.Logger = plain.DateTimeLevelMessage
    log.Info().Msg("Info message with date, time, and level")
}
```

#### Colored Console Loggers

The colored console logger writes to `os.Stdout` and uses colors:

```go
import (
    log "github.com/go-mods/zerolog-quick"
    "github.com/go-mods/zerolog-quick/console/colored"
)

func main() {
    // Colored message
    log.Logger = colored.Message
    log.Info().Msg("Colored info message")
    // Colored date and message
    log.Logger = colored.DateMessage
    log.Info().Msg("Colored info message with date")
    // Colored date, time, and message
    log.Logger = colored.DateTimeMessage
    log.Info().Msg("Colored info message with date and time")
    // Colored date, level, and message
    log.Logger = colored.DateLevelMessage
    log.Info().Msg("Colored info message with date and level")
    // Colored date, time, level, and message
    log.Logger = colored.DateTimeLevelMessage
    log.Info().Msg("Colored info message with date, time, and level")
}
```

### Custom Output

Every console logger has a constructor that returns the same logger writing to the `io.Writer` of your choice. It must not be nil:

| Logger | Constructor |
|--------|-------------|
| `Message` | `NewMessage(w)` |
| `DateMessage` | `NewDateMessage(w)` |
| `DateTimeMessage` | `NewDateTimeMessage(w)` |
| `DateLevelMessage` | `NewDateLevelMessage(w)` |
| `DateTimeLevelMessage` | `NewDateTimeLevelMessage(w)` |

They are available in both the `plain` and `colored` packages:

```go
import (
    "os"

    log "github.com/go-mods/zerolog-quick"
    "github.com/go-mods/zerolog-quick/console/colored"
)

func main() {
    // Keep stdout for the output of the program
    log.Logger = colored.NewDateTimeLevelMessage(os.Stderr)
    log.Info().Msg("Colored info message on stderr")
}
```

Use these constructors rather than the `Output` method of a console logger: `Output` replaces the console formatting, and the logger then writes JSON.

Colors depend on the environment of the process, not on the writer: a `colored` logger can write color codes to a file or a pipe. Use the `plain` logger when the output is not a terminal.

For the JSON logger, `Output` is all you need:

```go
import (
    "os"

    "github.com/go-mods/zerolog-quick/json"
)

func main() {
    logger := json.Default.Output(os.Stdout)
    logger.Info().Msg("JSON info message on stdout")
}
```

### Custom Logger Configuration

You can create custom logger configurations using the `zerolog.New()` function and the `zerolog.ConsoleWriter` options. For example:

```go
import (
    "os"
    "github.com/rs/zerolog"
    log "github.com/go-mods/zerolog-quick"
)

func main() {
    customLogger := zerolog.New(
        zerolog.ConsoleWriter{
            Out: os.Stdout,
            NoColor: false,
            TimeFormat: "2006-01-02 15:04:05",
            PartsOrder: []string{
                zerolog.TimestampFieldName,
                zerolog.LevelFieldName,
                zerolog.MessageFieldName,
            },
		},
    ).With().Timestamp().Logger()
    log.Logger = customLogger
    log.Info().Msg("Custom logger message")
}
```
## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE.md) file for details.
