package adventofcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2023Day02_Part1(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		expected := "8"

		reader := strings.NewReader(ReadFile("./day02_example.txt"))
		writer := &strings.Builder{}

		Problem2023Day02Part1(reader, writer)

		assert.Equal(t, expected, writer.String())
	})
	t.Run("puzzle", func(t *testing.T) {
		expected := "2204"

		reader := strings.NewReader(ReadFile("./day02_puzzle.txt"))
		writer := &strings.Builder{}

		Problem2023Day02Part1(reader, writer)

		assert.Equal(t, expected, writer.String())
	})
}

func TestProblem2023Day02_Part2(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		expected := "2286"

		reader := strings.NewReader(ReadFile("./day02_example.txt"))
		writer := &strings.Builder{}

		Problem2023Day02Part2(reader, writer)

		assert.Equal(t, expected, writer.String())
	})
	t.Run("puzzle", func(t *testing.T) {
		expected := "71036"

		reader := strings.NewReader(ReadFile("./day02_puzzle.txt"))
		writer := &strings.Builder{}

		Problem2023Day02Part2(reader, writer)

		assert.Equal(t, expected, writer.String())
	})
}
