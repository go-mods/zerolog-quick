package colored

import (
	"bytes"

	"github.com/gookit/color"
	"github.com/rs/zerolog"
)

var (
	ThemeTrace = &color.Theme{Name: "trace", Style: color.Style{color.OpReset, color.FgWhite}}
	ThemeDebug = &color.Theme{Name: "debug", Style: color.Style{color.OpReset, color.FgDefault}}
	ThemeInfo  = &color.Theme{Name: "info", Style: color.Style{color.OpReset, color.FgLightWhite}}
	ThemeWarn  = &color.Theme{Name: "warn", Style: color.Style{color.OpBold, color.FgLightYellow}}
	ThemeError = &color.Theme{Name: "error", Style: color.Style{color.OpReset, color.FgLightRed}}
	ThemeFatal = &color.Theme{Name: "fatal", Style: color.Style{color.OpBold, color.FgLightRed}}
	ThemePanic = &color.Theme{Name: "panic", Style: color.Style{color.OpBold, color.FgLightCyan}}
)

func Colorize(evt map[string]interface{}, buffer *bytes.Buffer) error {
	var m string
	if level, ok := evt[zerolog.LevelFieldName].(string); ok {
		switch level {
		case zerolog.LevelTraceValue:
			m = Trace(buffer.String())
		case zerolog.LevelDebugValue:
			m = Debug(buffer.String())
		case zerolog.LevelInfoValue:
			m = Info(buffer.String())
		case zerolog.LevelWarnValue:
			m = Warn(buffer.String())
		case zerolog.LevelErrorValue:
			m = Error(buffer.String())
		case zerolog.LevelFatalValue:
			m = Fatal(buffer.String())
		case zerolog.LevelPanicValue:
			m = Panic(buffer.String())
		default:
			m = Debug(buffer.String())
		}
	}

	if len(m) > 0 {
		buffer.Reset()
		buffer.WriteString(m)
	}

	return nil
}

func Trace(m string) string {
	return ThemeTrace.Sprint(m)
}

func Debug(m string) string {
	return ThemeDebug.Sprint(m)
}

func Info(m string) string {
	return ThemeInfo.Sprint(m)
}

func Warn(m string) string {
	return ThemeWarn.Sprint(m)
}

func Error(m string) string {
	return ThemeError.Sprint(m)
}

func Fatal(m string) string {
	return ThemeFatal.Sprint(m)
}

func Panic(m string) string {
	return ThemePanic.Sprint(m)
}
