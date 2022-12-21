package adventofcode

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2022Day06_Example1(t *testing.T) {
	inputs := []string{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		"bvwbjplbgvbhsrlpgdmjqwftvncz",
		"nppdvjthqldpwncqszvftbrmjlhg",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
	}
	t.Run("part 1", func(t *testing.T) {
		expected := []string{
			"7",
			"5",
			"6",
			"10",
			"11",
		}
		for i, input := range inputs {
			writer := &strings.Builder{}

			Problem2022Day06Part1(strings.NewReader(input), writer)

			assert.Equal(t, expected[i], writer.String())
		}
	})
	t.Run("part 2", func(t *testing.T) {
		expected := []string{
			"19",
			"23",
			"23",
			"29",
			"26",
		}
		for i, input := range inputs {
			writer := &strings.Builder{}

			Problem2022Day06Part2(strings.NewReader(input), writer)

			assert.Equal(t, expected[i], writer.String())
		}
	})
}

func TestProblem2022Day06_Puzzle(t *testing.T) {
	fileContent, err := os.ReadFile("./day06_input.txt")
	if err != nil {
		t.Error(err)
		return
	}
	input := strings.TrimSpace(string(fileContent))
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day06Part1(strings.NewReader(input), writer)

		assert.Equal(t, "1623", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day06Part2(strings.NewReader(input), writer)

		assert.Equal(t, "3774", writer.String())
	})
}
