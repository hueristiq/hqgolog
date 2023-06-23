# hqgolog

[![license](https://img.shields.io/badge/license-MIT-gray.svg?color=0040FF)](https://github.com/hueristiq/hqgolog/blob/master/LICENSE) ![maintenance](https://img.shields.io/badge/maintained%3F-yes-0040ff.svg) [![open issues](https://img.shields.io/github/issues-raw/hueristiq/hqgolog.svg?style=flat&color=0040ff)](https://github.com/hueristiq/hqgolog/issues?q=is:issue+is:open) [![closed issues](https://img.shields.io/github/issues-closed-raw/hueristiq/hqgolog.svg?style=flat&color=0040ff)](https://github.com/hueristiq/hqgolog/issues?q=is:issue+is:closed) [![contribution](https://img.shields.io/badge/contributions-welcome-0040ff.svg)](https://github.com/hueristiq/hqgolog/blob/master/CONTRIBUTING.md)

A [Go(Golang)](https://golang.org/) package to do structured logging.

## Installation

```bash
go get -v -u github.com/hueristiq/hqgolog
```

## Usage

```go
package main

import (
	"github.com/hueristiq/hqgolog"
	"github.com/hueristiq/hqgolog/formatter"
	"github.com/hueristiq/hqgolog/levels"
)

func main() {
	hqgolog.DefaultLogger.SetMaxLevel(levels.LevelDebug)
	hqgolog.DefaultLogger.SetFormatter(formatter.NewCLI(&formatter.CLIOptions{
		Colorize: true,
	}))


	hqgolog.Print().Msg("Print message")
	hqgolog.Info().Msg("Info message")
	hqgolog.Warn().Msg("Warn message")
	hqgolog.Error().Msg("Error message")
	hqgolog.Fatal().Msg("Fatal message")
}
```

## Contributing

[Issues](https://github.com/hueristiq/hqgolog/issues) and [Pull Requests](https://github.com/hueristiq/hqgolog/pulls) are welcome! **Check out the [contribution guidelines](./CONTRIBUTING.md)**.

## Licensing

This package is distributed under the [MIT license](https://github.com/hueristiq/hqgolog/blob/master/LICENSE).