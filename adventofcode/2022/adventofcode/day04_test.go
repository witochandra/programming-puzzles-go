package adventofcode

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2022Day04_Example(t *testing.T) {
	input := strings.TrimSpace(`
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`)
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day04Part1(strings.NewReader(input), writer)

		assert.Equal(t, "2", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day04Part2(strings.NewReader(input), writer)

		assert.Equal(t, "4", writer.String())
	})
}

func TestProblem2022Day04_Puzzle(t *testing.T) {
	fileContent, err := os.ReadFile("./day04_input.txt")
	if err != nil {
		t.Error(err)
		return
	}
	input := strings.TrimSpace(string(fileContent))
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day04Part1(strings.NewReader(input), writer)

		assert.Equal(t, "536", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day04Part2(strings.NewReader(input), writer)

		assert.Equal(t, "845", writer.String())
	})
}
