package adventofcode

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2022Day09_Example(t *testing.T) {
	input := strings.TrimSpace(`
R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`)
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day09Part1(strings.NewReader(input), writer)

		assert.Equal(t, "13", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day09Part2(strings.NewReader(input), writer)

		assert.Equal(t, "1", writer.String())
	})
}

func TestProblem2022Day09_Example2(t *testing.T) {
	input := strings.TrimSpace(`
R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`)
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day09Part1(strings.NewReader(input), writer)

		assert.Equal(t, "88", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day09Part2(strings.NewReader(input), writer)

		assert.Equal(t, "36", writer.String())
	})
}

func TestProblem2022Day09_Puzzle(t *testing.T) {
	fileContent, err := os.ReadFile("./day09_input.txt")
	if err != nil {
		t.Error(err)
		return
	}
	input := strings.TrimSpace(string(fileContent))
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day09Part1(strings.NewReader(input), writer)

		assert.Equal(t, "6011", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day09Part2(strings.NewReader(input), writer)

		assert.Equal(t, "2419", writer.String())
	})
}
