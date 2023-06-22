package levels

type LevelInt int
type LevelStr string

const (
	LevelSilent LevelStr = "silent"
	LevelFatal  LevelStr = "fatal"
	LevelError  LevelStr = "error"
	LevelWarn   LevelStr = "Warn"
	LevelInfo   LevelStr = "info"
	LevelDebug  LevelStr = "debug"
)

var Levels = map[LevelStr]LevelInt{
	LevelSilent: LevelInt(0),
	LevelFatal:  LevelInt(1),
	LevelError:  LevelInt(2),
	LevelWarn:   LevelInt(3),
	LevelInfo:   LevelInt(4),
	LevelDebug:  LevelInt(5),
}
