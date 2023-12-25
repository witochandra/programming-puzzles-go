package adventofcode

import (
	"os"
)

func ReadFile(path string) string {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(fileContent)
}
