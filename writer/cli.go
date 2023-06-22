package writer

import (
	"os"
	"sync"

	"github.com/hueristiq/hqgolog/levels"
)

type CLI struct{}

var (
	_     Writer = &CLI{}
	mutex *sync.Mutex
)

func NewCLI() *CLI {
	mutex = &sync.Mutex{}

	return &CLI{}
}

func (w *CLI) Write(data []byte, level levels.LevelInt) {
	mutex.Lock()
	defer mutex.Unlock()

	switch level {
	case levels.Levels[levels.LevelError], levels.Levels[levels.LevelFatal]:
		os.Stderr.Write(data)
		os.Stderr.Write([]byte("\n"))
	default:
		os.Stdout.Write(data)
		os.Stdout.Write([]byte("\n"))
	}
}
