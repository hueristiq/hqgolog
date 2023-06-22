package hqgolog

import (
	"github.com/hueristiq/hqgolog/formatter"
	"github.com/hueristiq/hqgolog/levels"
	"github.com/hueristiq/hqgolog/writer"
)

var (
	DefaultLogger *Logger
)

func init() {
	DefaultLogger = &Logger{}
	DefaultLogger.SetMaxLevel(levels.LevelDebug)
	DefaultLogger.SetFormatter(formatter.NewCLI(&formatter.CLIOptions{
		Colorize: true,
	}))
	DefaultLogger.SetWriter(writer.NewCLI())
}

func Print() (event *Event) {
	event = &Event{
		logger:   DefaultLogger,
		level:    levels.LevelInt(-1),
		metadata: make(map[string]string),
	}

	return event
}

func Debug() (event *Event) {
	level := levels.Levels[levels.LevelDebug]

	event = &Event{
		logger:   DefaultLogger,
		level:    level,
		metadata: make(map[string]string),
	}

	return
}

func Info() (event *Event) {
	level := levels.Levels[levels.LevelInfo]

	event = &Event{
		logger:   DefaultLogger,
		level:    level,
		metadata: make(map[string]string),
	}

	return
}

func Warn() (event *Event) {
	level := levels.Levels[levels.LevelWarn]

	event = &Event{
		logger:   DefaultLogger,
		level:    level,
		metadata: make(map[string]string),
	}

	return
}

func Error() (event *Event) {
	level := levels.Levels[levels.LevelError]

	event = &Event{
		logger:   DefaultLogger,
		level:    level,
		metadata: make(map[string]string),
	}

	return
}

func Fatal() (event *Event) {
	level := levels.Levels[levels.LevelFatal]

	event = &Event{
		logger:   DefaultLogger,
		level:    level,
		metadata: make(map[string]string),
	}

	return
}
