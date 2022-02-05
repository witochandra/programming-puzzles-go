package adventofcode

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2021Day03_Example(t *testing.T) {
	input := strings.TrimSpace(`
00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
`)
	t.Run("part 1", func(t *testing.T) {
		expected := "198"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day03Part1(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
	t.Run("part 2", func(t *testing.T) {
		expected := "230"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day03Part2(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
}

func TestProblem2021Day03_Puzzle(t *testing.T) {
	fileContent, err := os.ReadFile("./day03_input.txt")
	if err != nil {
		t.Error(err)
		return
	}
	input := strings.TrimSpace(string(fileContent))
	t.Run("part 1", func(t *testing.T) {
		expected := "2003336"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day03Part1(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
	t.Run("part 2", func(t *testing.T) {
		expected := "1877139"

		reader := strings.NewReader(input)
		writer := &strings.Builder{}

		Problem2021Day03Part2(reader, writer)

		assert.Equal(t, expected, strings.TrimSpace(writer.String()))
	})
}
