package adventofcode

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2022Day03_Example(t *testing.T) {
	input := strings.TrimSpace(`
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`)
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day03Part1(strings.NewReader(input), writer)

		assert.Equal(t, "157", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day03Part2(strings.NewReader(input), writer)

		assert.Equal(t, "70", writer.String())
	})
}

func TestProblem2022Day03_Puzzle(t *testing.T) {
	fileContent, err := os.ReadFile("./day03_input.txt")
	if err != nil {
		t.Error(err)
		return
	}
	input := strings.TrimSpace(string(fileContent))
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day03Part1(strings.NewReader(input), writer)

		assert.Equal(t, "8123", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day03Part2(strings.NewReader(input), writer)

		assert.Equal(t, "2620", writer.String())
	})
}
