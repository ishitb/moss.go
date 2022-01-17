package utils

import (
	"errors"
	"os"
	"path/filepath"
)

func ReadFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		ErrorP(err.Error())
	}

	return string(data)
}

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

func GetFilesByWildcard(wildcard string) []string {
	matches, err := filepath.Glob(wildcard)

	for _, file := range matches {
		size, existError := GetSize(file)

		if size <= 0 || existError != nil {
			ErrorP(existError.Error())
		}
	}

	if err != nil {
		ErrorP(err.Error())
	}

	return matches
}
