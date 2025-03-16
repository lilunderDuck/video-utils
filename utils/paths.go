package utils

import (
	"os"
	"path/filepath"
)

func GetCurrentDir() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}

func GetFileName(fromPath string) string {
	return filepath.Base(fromPath)
}

func MergeDir(paths ...string) string {
	return filepath.Join(paths...)
}
