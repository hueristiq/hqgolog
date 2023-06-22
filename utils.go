package hqgolog

import (
	"strings"

	"golang.org/x/term"
)

func appendRest(data []byte, character string) []byte {
	dataStr := string(data)
	dataLen := len(dataStr)

	width, _, _ := term.GetSize(0)

	dataStr = dataStr + strings.Repeat(character[0:1], width-dataLen)

	return []byte(dataStr)
}
