package adventofcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2025Day01_Part1(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		expected := "3"

		reader := strings.NewReader(ReadFile("./inputs/day01_example.txt"))
		writer := &strings.Builder{}

		Problem2025Day01Part1(reader, writer)

		assert.Equal(t, expected, writer.String())
	})
	t.Run("puzzle", func(t *testing.T) {
		expected := "1023"

		reader := strings.NewReader(ReadFile("./inputs/day01_puzzle.txt"))
		writer := &strings.Builder{}

		Problem2025Day01Part1(reader, writer)

		assert.Equal(t, expected, writer.String())
	})
}

func TestProblem2025Day01_Part2(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		expected := "6"

		reader := strings.NewReader(ReadFile("./inputs/day01_example.txt"))
		writer := &strings.Builder{}

		Problem2025Day01Part2(reader, writer)

		assert.Equal(t, expected, writer.String())
	})
	t.Run("puzzle", func(t *testing.T) {
		expected := "1023"

		reader := strings.NewReader(ReadFile("./inputs/day01_puzzle.txt"))
		writer := &strings.Builder{}

		Problem2025Day01Part2(reader, writer)

		assert.Equal(t, expected, writer.String())
	})
}
