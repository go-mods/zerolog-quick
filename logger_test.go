package log_test

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/gookit/color"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/natefinch/lumberjack.v2"

	log "github.com/go-mods/zerolog-quick"
	"github.com/go-mods/zerolog-quick/console/colored"
	"github.com/go-mods/zerolog-quick/console/plain"
	jsonlog "github.com/go-mods/zerolog-quick/json"
)

const (
	// childEnv makes the test binary log and exit instead of running the tests.
	// Its value selects what the child process logs.
	childEnv = "ZEROLOG_QUICK_TEST_CHILD"
	// childJSON selects the default JSON logger.
	childJSON = "json"
	// childToStderr is appended to the name of a console logger to build it on os.Stderr.
	childToStderr = "/stderr"

	datePattern     = `\d{4}-\d{2}-\d{2} `
	dateTimePattern = `\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2} `
	ansiReset       = "\x1b[0m"
)

// event is a message logged by logEvents and what a console logger must print for it.
type event struct {
	level   string // as printed by the loggers with a level
	message string
	ansi    string // color sequence opening the line of a colored logger
}

var events = []event{
	{"DEBUG", "Printed message", "\x1b[0;39m"},
	{"TRACE", "trace message", "\x1b[0;37m"},
	{"DEBUG", "debug message", "\x1b[0;39m"},
	{"INFO ", "info message", "\x1b[0;97m"},
	{"WARN ", "warn message", "\x1b[1;93m"},
	{"ERROR", "error message", "\x1b[0;91m"},
	{"FATAL", "fatal message", "\x1b[1;91m"},
	{"PANIC", "panic message", "\x1b[1;96m"},
}

// consoleCase describes a console logger and the format of its lines.
type consoleCase struct {
	name        string
	logger      zerolog.Logger
	constructor func(io.Writer) zerolog.Logger
	timestamp   string // pattern of the timestamp, empty if the logger has none
	level       bool
	colored     bool
}

var consoleCases = []consoleCase{
	{
		name:        "ConsolePlainMessage",
		logger:      plain.Message,
		constructor: plain.NewMessage,
	},
	{
		name:        "ConsolePlainDateMessage",
		logger:      plain.DateMessage,
		constructor: plain.NewDateMessage,
		timestamp:   datePattern,
	},
	{
		name:        "ConsolePlainDateTimeMessage",
		logger:      plain.DateTimeMessage,
		constructor: plain.NewDateTimeMessage,
		timestamp:   dateTimePattern,
	},
	{
		name:        "ConsolePlainDateLevelMessage",
		logger:      plain.DateLevelMessage,
		constructor: plain.NewDateLevelMessage,
		timestamp:   datePattern,
		level:       true,
	},
	{
		name:        "ConsolePlainDateTimeLevelMessage",
		logger:      plain.DateTimeLevelMessage,
		constructor: plain.NewDateTimeLevelMessage,
		timestamp:   dateTimePattern,
		level:       true,
	},
	{
		name:        "ConsoleColoredMessage",
		logger:      colored.Message,
		constructor: colored.NewMessage,
		colored:     true,
	},
	{
		name:        "ConsoleColoredDateMessage",
		logger:      colored.DateMessage,
		constructor: colored.NewDateMessage,
		timestamp:   datePattern,
		colored:     true,
	},
	{
		name:        "ConsoleColoredDateTimeMessage",
		logger:      colored.DateTimeMessage,
		constructor: colored.NewDateTimeMessage,
		timestamp:   dateTimePattern,
		colored:     true,
	},
	{
		name:        "ConsoleColoredDateLevelMessage",
		logger:      colored.DateLevelMessage,
		constructor: colored.NewDateLevelMessage,
		timestamp:   datePattern,
		level:       true,
		colored:     true,
	},
	{
		name:        "ConsoleColoredDateTimeLevelMessage",
		logger:      colored.DateTimeLevelMessage,
		constructor: colored.NewDateTimeLevelMessage,
		timestamp:   dateTimePattern,
		level:       true,
		colored:     true,
	},
}

// TestMain turns the test binary into a child process when childEnv is set.
// The predefined loggers capture os.Stdout and os.Stderr when the program starts,
// so only another process can observe which stream they really write to.
func TestMain(m *testing.M) {
	if child := os.Getenv(childEnv); child != "" {
		logFromChild(child)
		os.Exit(0)
	}
	os.Exit(m.Run())
}

func logFromChild(child string) {
	if child == childJSON {
		log.Info().Msg("Info message")
		return
	}

	// Color support is detected from the environment: force it for a stable output.
	color.ForceColor()
	color.Enable = true

	name, toStderr := strings.CutSuffix(child, childToStderr)
	for _, c := range consoleCases {
		if c.name != name {
			continue
		}
		log.Logger = c.logger
		if toStderr {
			log.Logger = c.constructor(os.Stderr)
		}
		logEvents()
		return
	}
	panic("unknown child: " + child)
}

// logEvents logs the messages of events through the global logger.
func logEvents() {
	log.Print("Printed message")
	log.Trace().Msg("trace message")
	log.Debug().Msg("debug message")
	log.Info().Msg("info message")
	log.Warn().Msg("warn message")
	log.Error().Msg("error message")
	log.WithLevel(zerolog.FatalLevel).Msg("fatal message")
	log.WithLevel(zerolog.PanicLevel).Msg("panic message")
}

// runChild runs the test binary as a child process and returns its stdout and stderr.
func runChild(t *testing.T, child string) (stdout, stderr string) {
	t.Helper()

	executable, err := os.Executable()
	require.NoError(t, err)

	var out, errOut bytes.Buffer
	cmd := exec.Command(executable)
	cmd.Env = append(os.Environ(), childEnv+"="+child)
	cmd.Stdout = &out
	cmd.Stderr = &errOut
	require.NoError(t, cmd.Run(), "stderr: %s", errOut.String())

	return out.String(), errOut.String()
}

// forceColor enables colors for the duration of the test.
func forceColor(t *testing.T) {
	t.Helper()

	level, enable := color.ForceColor(), color.Enable
	color.Enable = true
	t.Cleanup(func() {
		color.ForceSetColorLevel(level)
		color.Enable = enable
	})
}

// useLogger sets the global logger for the duration of the test.
func useLogger(t *testing.T, logger zerolog.Logger) {
	t.Helper()

	previous := log.Logger
	log.Logger = logger
	t.Cleanup(func() { log.Logger = previous })
}

// assertEvents checks that output holds exactly one line per event, in the format of c.
func assertEvents(t *testing.T, c consoleCase, output string) {
	t.Helper()

	lines := strings.Split(strings.TrimSuffix(output, "\n"), "\n")
	require.Len(t, lines, len(events), "output: %q", output)

	for i, e := range events {
		pattern := c.timestamp
		if c.level {
			pattern += `\[` + e.level + `\] `
		}
		pattern += regexp.QuoteMeta(e.message)
		if c.colored {
			pattern = regexp.QuoteMeta(e.ansi) + pattern + regexp.QuoteMeta(ansiReset)
		}
		assert.Regexp(t, "^"+pattern+"$", lines[i])
	}
}

func TestConsoleLoggers(t *testing.T) {
	for _, c := range consoleCases {
		t.Run(c.name, func(t *testing.T) {
			t.Run("writes to stdout by default", func(t *testing.T) {
				stdout, stderr := runChild(t, c.name)

				assertEvents(t, c, stdout)
				assert.Empty(t, stderr)
			})

			t.Run("constructor writes to os.Stderr", func(t *testing.T) {
				stdout, stderr := runChild(t, c.name+childToStderr)

				assertEvents(t, c, stderr)
				assert.Empty(t, stdout)
			})

			t.Run("constructor writes to any writer", func(t *testing.T) {
				forceColor(t)
				var buf bytes.Buffer
				useLogger(t, c.constructor(&buf))

				logEvents()

				assertEvents(t, c, buf.String())
			})
		})
	}
}

func TestConsoleLoggersWithNilWriter(t *testing.T) {
	constructors := map[string]func(io.Writer) zerolog.Logger{
		"plain.NewMessage":                plain.NewMessage,
		"plain.NewDateMessage":            plain.NewDateMessage,
		"plain.NewDateTimeMessage":        plain.NewDateTimeMessage,
		"plain.NewDateLevelMessage":       plain.NewDateLevelMessage,
		"plain.NewDateTimeLevelMessage":   plain.NewDateTimeLevelMessage,
		"colored.NewMessage":              colored.NewMessage,
		"colored.NewDateMessage":          colored.NewDateMessage,
		"colored.NewDateTimeMessage":      colored.NewDateTimeMessage,
		"colored.NewDateLevelMessage":     colored.NewDateLevelMessage,
		"colored.NewDateTimeLevelMessage": colored.NewDateTimeLevelMessage,
	}

	for name, build := range constructors {
		t.Run(name, func(t *testing.T) {
			logger := build(nil)

			assert.NotPanics(t, func() { logger.Info().Msg("Info message") })
		})
	}
}

func TestJsonDefault(t *testing.T) {
	stdout, stderr := runChild(t, childJSON)

	assert.Empty(t, stdout)
	assertJSONInfo(t, stderr)
}

func TestJsonDefaultOutput(t *testing.T) {
	var buf bytes.Buffer
	logger := jsonlog.Default.Output(&buf)

	logger.Info().Msg("Info message")

	assertJSONInfo(t, buf.String())
}

// assertJSONInfo checks that output is the single JSON line of an info message.
func assertJSONInfo(t *testing.T, output string) {
	t.Helper()

	var entry map[string]any
	require.NoError(t, json.Unmarshal([]byte(output), &entry), "output: %q", output)
	assert.Equal(t, "info", entry[zerolog.LevelFieldName])
	assert.Equal(t, "Info message", entry[zerolog.MessageFieldName])
	assert.NotEmpty(t, entry[zerolog.TimestampFieldName])
}

func TestMulti(t *testing.T) {
	stackMarshaler, timeFieldFormat := zerolog.ErrorStackMarshaler, zerolog.TimeFieldFormat
	t.Cleanup(func() {
		zerolog.ErrorStackMarshaler = stackMarshaler
		zerolog.TimeFieldFormat = timeFieldFormat
	})
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = time.RFC3339Nano
	forceColor(t)

	var consoleBuf bytes.Buffer
	var consoleOut io.Writer = zerolog.ConsoleWriter{
		Out:        &consoleBuf,
		NoColor:    true,
		TimeFormat: "2006-01-02",
		PartsOrder: []string{
			zerolog.TimestampFieldName,
			zerolog.MessageFieldName,
		},
		FormatExtra: colored.Colorize,
	}

	// No MaxBackups, MaxAge or Compress: lumberjack would otherwise scan the
	// directory from a goroutine that outlives Close, during the cleanup of TempDir.
	filename := filepath.Join(t.TempDir(), "test.log")
	fileOut := &lumberjack.Logger{
		Filename: filename,
		MaxSize:  5,
	}

	multi := zerolog.MultiLevelWriter(consoleOut, fileOut)

	logger := zerolog.New(multi).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	logger.Print("Printed message")
	logger.Trace().Msg("trace message")
	logger.Debug().Msg("debug message")
	logger.Info().Msg("info message")
	logger.Warn().Msg("warn message")
	logger.Error().Msg("error message")
	logger.WithLevel(zerolog.FatalLevel).Msg("fatal message")
	logger.WithLevel(zerolog.PanicLevel).Msg("panic message")

	// An open file cannot be removed on Windows: TempDir needs it closed.
	require.NoError(t, fileOut.Close())

	assertEvents(t, consoleCase{timestamp: datePattern, colored: true}, consoleBuf.String())

	content, err := os.ReadFile(filename)
	require.NoError(t, err)
	lines := strings.Split(strings.TrimSuffix(string(content), "\n"), "\n")
	require.Len(t, lines, len(events))
	for i, e := range events {
		var entry map[string]any
		require.NoError(t, json.Unmarshal([]byte(lines[i]), &entry), "line: %q", lines[i])
		assert.Equal(t, strings.ToLower(strings.TrimSpace(e.level)), entry[zerolog.LevelFieldName])
		assert.Equal(t, e.message, entry[zerolog.MessageFieldName])
		assert.Contains(t, entry[zerolog.CallerFieldName], "logger_test.go")
	}
}
