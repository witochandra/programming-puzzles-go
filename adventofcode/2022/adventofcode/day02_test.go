package adventofcode

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2022Day02_Example(t *testing.T) {
	input := strings.TrimSpace(`
A Y
B X
C Z
`)
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day02Part1(strings.NewReader(input), writer)

		assert.Equal(t, "15", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day02Part2(strings.NewReader(input), writer)

		assert.Equal(t, "12", writer.String())
	})
}

func TestProblem2022Day02_Puzzle(t *testing.T) {
	fileContent, err := os.ReadFile("./day02_input.txt")
	if err != nil {
		t.Error(err)
		return
	}
	input := strings.TrimSpace(string(fileContent))
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day02Part1(strings.NewReader(input), writer)

		assert.Equal(t, "13484", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day02Part2(strings.NewReader(input), writer)

		assert.Equal(t, "13433", writer.String())
	})
}
