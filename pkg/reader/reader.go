package reader

import (
	"io"
	"os"
	"strings"
)

func ReadStings(path string) []string {
	data := Read(path)

	return strings.Split(data, "\n")
}

func SpltBy(path, sep string) []string {
	data := Read(path)

	return strings.Split(data, sep)
}

func Read(path string) string {
	file, _ := os.Open(path)
	defer file.Close()

	data, _ := io.ReadAll(file)

	return string(data)
}
