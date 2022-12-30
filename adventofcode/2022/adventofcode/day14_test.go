package adventofcode

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2022Day14_Example(t *testing.T) {
	input := strings.TrimSpace(`
498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
`)
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day14Part1(strings.NewReader(input), writer)

		assert.Equal(t, "24", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day14Part2(strings.NewReader(input), writer)

		assert.Equal(t, "93", writer.String())
	})
}

func TestProblem2022Day14_Puzzle(t *testing.T) {
	fileContent, err := os.ReadFile("./day14_input.txt")
	if err != nil {
		t.Error(err)
		return
	}
	input := strings.TrimSpace(string(fileContent))
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day14Part1(strings.NewReader(input), writer)

		assert.Equal(t, "897", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day14Part2(strings.NewReader(input), writer)

		assert.Equal(t, "26683", writer.String())
	})
}
