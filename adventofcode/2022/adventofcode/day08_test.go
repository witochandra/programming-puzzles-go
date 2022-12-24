package adventofcode

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2022Day08_Example(t *testing.T) {
	input := strings.TrimSpace(`
30373
25512
65332
33549
35390
`)
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day08Part1(strings.NewReader(input), writer)

		assert.Equal(t, "21", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day08Part2(strings.NewReader(input), writer)

		assert.Equal(t, "8", writer.String())
	})
}

func TestProblem2022Day08_Puzzle(t *testing.T) {
	fileContent, err := os.ReadFile("./day08_input.txt")
	if err != nil {
		t.Error(err)
		return
	}
	input := strings.TrimSpace(string(fileContent))
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day08Part1(strings.NewReader(input), writer)

		assert.Equal(t, "1849", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day08Part2(strings.NewReader(input), writer)

		assert.Equal(t, "201600", writer.String())
	})
}
