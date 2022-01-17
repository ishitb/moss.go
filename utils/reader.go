package utils

import (
	"bufio"
	"strings"
)

func GetInput(prompt string, reader *bufio.Reader) string {
	Info("%v: ", prompt)

	uniqueId, _ := reader.ReadString('\n')
	return strings.TrimSpace(uniqueId)
}
