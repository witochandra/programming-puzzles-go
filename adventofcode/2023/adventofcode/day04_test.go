package adventofcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2023Day04_Part1(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		reader := strings.NewReader(ReadFile("./day04_example.txt"))
		writer := &strings.Builder{}

		Problem2023Day04Part1(reader, writer)

		assert.Equal(t, "13", writer.String())
	})
	t.Run("puzzle", func(t *testing.T) {
		reader := strings.NewReader(ReadFile("./day04_puzzle.txt"))
		writer := &strings.Builder{}

		Problem2023Day04Part1(reader, writer)

		assert.Equal(t, "15205", writer.String())
	})
}

func TestProblem2023Day04_Part2(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		reader := strings.NewReader(ReadFile("./day04_example.txt"))
		writer := &strings.Builder{}

		Problem2023Day04Part2(reader, writer)

		assert.Equal(t, "30", writer.String())
	})
	t.Run("puzzle", func(t *testing.T) {
		reader := strings.NewReader(ReadFile("./day04_puzzle.txt"))
		writer := &strings.Builder{}

		Problem2023Day04Part2(reader, writer)

		assert.Equal(t, "108402", writer.String())
	})
}
