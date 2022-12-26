package adventofcode

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2022Day12_Example(t *testing.T) {
	input := strings.TrimSpace(`
Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
`)
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day12Part1(strings.NewReader(input), writer)

		assert.Equal(t, "31", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day12Part2(strings.NewReader(input), writer)

		assert.Equal(t, "29", writer.String())
	})
}

func TestProblem2022Day12_Puzzle(t *testing.T) {
	fileContent, err := os.ReadFile("./day12_input.txt")
	if err != nil {
		t.Error(err)
		return
	}
	input := strings.TrimSpace(string(fileContent))
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day12Part1(strings.NewReader(input), writer)

		assert.Equal(t, "497", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day12Part2(strings.NewReader(input), writer)

		assert.Equal(t, "492", writer.String())
	})
}
