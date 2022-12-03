package reader

import (
	"io"
	"os"
	"strings"
)

func ReadStings(path string) []string {
	file, _ := os.Open(path)
	defer file.Close()

	data, _ := io.ReadAll(file)

	return strings.Split(string(data), "\n")
}
