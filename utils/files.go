package utils

import "os"

func SaveFile(path string, data []byte) {
	err := os.WriteFile(path, data, 0644)

	if err != nil {
		panic(err)
	}
}
