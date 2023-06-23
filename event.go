package hqgolog

import (
	"fmt"

	"github.com/hueristiq/hqgolog/levels"
)

type Event struct { //nolint:govet // To be refactored.
	level    levels.LevelInt
	message  string
	metadata map[string]string

	logger *Logger
}

func (event *Event) Label(label string) *Event {
	event.metadata["label"] = label

	return event
}

func (event *Event) Rest(character string) *Event {
	event.metadata["rest"] = character

	return event
}

// // Str adds a string metadata item to the log
// func (event *Event) Str(key, value string) *Event {
// 	event.metadata[key] = value

// 	return event
// }

func (event *Event) Msg(message string) {
	event.message = message

	event.logger.Log(event)
}

func (event *Event) Msgf(format string, args ...interface{}) {
	event.message = fmt.Sprintf(format, args...)

	event.logger.Log(event)
}
