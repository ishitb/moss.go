package utils

import (
	"bufio"
	"fmt"
	"strings"
)

func GetInput(prompt string, reader *bufio.Reader) string {
	fmt.Printf("%v: ", prompt)

	uniqueId, _ := reader.ReadString('\n')
	return strings.TrimSpace(uniqueId)
}
