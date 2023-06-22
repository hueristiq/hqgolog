package writer

import "github.com/hueristiq/hqgolog/levels"

type Writer interface {
	Write(data []byte, level levels.LevelInt)
}
