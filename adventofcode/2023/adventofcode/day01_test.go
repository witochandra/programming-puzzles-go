package adventofcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2023Day01_Part1(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		expected := "142"

		reader := strings.NewReader(ReadFile("./day01_part1_example.txt"))
		writer := &strings.Builder{}

		Problem2023Day01Part1(reader, writer)

		assert.Equal(t, expected, writer.String())
	})
	t.Run("puzzle", func(t *testing.T) {
		expected := "57346"

		reader := strings.NewReader(ReadFile("./day01_puzzle.txt"))
		writer := &strings.Builder{}

		Problem2023Day01Part1(reader, writer)

		assert.Equal(t, expected, writer.String())
	})
}

func TestProblem2023Day01_Part2(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		expected := "281"

		reader := strings.NewReader(ReadFile("./day01_part2_example.txt"))
		writer := &strings.Builder{}

		Problem2023Day01Part2(reader, writer)

		assert.Equal(t, expected, writer.String())
	})
	t.Run("puzzle", func(t *testing.T) {
		expected := "57345"

		reader := strings.NewReader(ReadFile("./day01_puzzle.txt"))
		writer := &strings.Builder{}

		Problem2023Day01Part2(reader, writer)

		assert.Equal(t, expected, writer.String())
	})
}
