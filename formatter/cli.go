package formatter

import (
	"bytes"

	"github.com/hueristiq/hqgolog/levels"
	"github.com/logrusorgru/aurora/v3"
)

type CLI struct{}

type CLIOptions struct {
	Colorize bool
}

var (
	_  Formatter     = &CLI{}
	au aurora.Aurora = aurora.NewAurora(false)
)

func NewCLI(options *CLIOptions) *CLI {
	au = aurora.NewAurora(options.Colorize)

	return &CLI{}
}

func (cli *CLI) Format(event *Log) (data []byte, err error) {
	cli.colorizeLabel(event)

	buffer := &bytes.Buffer{}
	buffer.Grow(len(event.Message))

	if label, ok := event.Metadata["label"]; ok && label != "" {
		buffer.WriteRune('[')
		buffer.WriteString(label)
		buffer.WriteRune(']')
		buffer.WriteRune(' ')
		// delete(event.Metadata, "label")
	}

	buffer.WriteString(event.Message)

	// for k, v := range event.Metadata {
	// 	buffer.WriteRune(' ')
	// 	buffer.WriteString(cli.colorizeKey(k))
	// 	buffer.WriteRune('=')
	// 	buffer.WriteString(v)
	// }
	data = buffer.Bytes()

	return
}

// func (cli *CLI) colorizeKey(key string) string {
// 	return au.Bold(key).String()
// }

func (cli *CLI) colorizeLabel(event *Log) {
	label := event.Metadata["label"]

	if label == "" {
		return
	}

	switch event.Level {
	case levels.Levels[levels.LevelSilent]:
		return
	case levels.Levels[levels.LevelFatal]:
		event.Metadata["label"] = au.BrightRed(label).Bold().String()
	case levels.Levels[levels.LevelError]:
		event.Metadata["label"] = au.BrightRed(label).Bold().String()
	case levels.Levels[levels.LevelWarn]:
		event.Metadata["label"] = au.BrightYellow(label).Bold().String()
	case levels.Levels[levels.LevelInfo]:
		event.Metadata["label"] = au.BrightBlue(label).Bold().String()
	case levels.Levels[levels.LevelDebug]:
		event.Metadata["label"] = au.BrightMagenta(label).Bold().String()
	}
}
