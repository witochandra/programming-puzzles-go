package adventofcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2023Day03_Part1(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		reader := strings.NewReader(ReadFile("./day03_example.txt"))
		writer := &strings.Builder{}

		Problem2023Day03Part1(reader, writer)

		assert.Equal(t, "4361", writer.String())
	})
	t.Run("puzzle", func(t *testing.T) {
		reader := strings.NewReader(ReadFile("./day03_puzzle.txt"))
		writer := &strings.Builder{}

		Problem2023Day03Part1(reader, writer)

		assert.Equal(t, "532331", writer.String())
	})
}

func TestProblem2023Day03_Part2(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		reader := strings.NewReader(ReadFile("./day03_example.txt"))
		writer := &strings.Builder{}

		Problem2023Day03Part2(reader, writer)

		assert.Equal(t, "467835", writer.String())
	})
	t.Run("puzzle", func(t *testing.T) {
		reader := strings.NewReader(ReadFile("./day03_puzzle.txt"))
		writer := &strings.Builder{}

		Problem2023Day03Part2(reader, writer)

		assert.Equal(t, "82301120", writer.String())
	})
}
