package formatter

import "github.com/hueristiq/hqgolog/levels"

type Log struct { //nolint:govet // To be refactored.
	Message  string
	Level    levels.LevelInt
	Metadata map[string]string
}

type Formatter interface {
	Format(log *Log) ([]byte, error)
}
