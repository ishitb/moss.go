package utils

import (
	"errors"
	"os"
)

func SaveFile(path string, data []byte) {
	err := os.WriteFile(path, data, 0644)

	if err != nil {
		panic(err)
	}
}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func GetSize(name string) (int, error) {
	fileInfo, err := os.Stat(name)
	if err == nil {
		return int(fileInfo.Size()), nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return 0, nil
	}
	return 0, err
}
