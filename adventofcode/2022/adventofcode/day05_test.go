package adventofcode

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2022Day05_Example(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day05Part1(strings.NewReader(input), writer)

		assert.Equal(t, "CMZ", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day05Part2(strings.NewReader(input), writer)

		assert.Equal(t, "MCD", writer.String())
	})
}

func TestProblem2022Day05_Puzzle(t *testing.T) {
	fileContent, err := os.ReadFile("./day05_input.txt")
	if err != nil {
		t.Error(err)
		return
	}
	input := strings.TrimSpace(string(fileContent))
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day05Part1(strings.NewReader(input), writer)

		assert.Equal(t, "TGWSMRBPN", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day05Part2(strings.NewReader(input), writer)

		assert.Equal(t, "TZLTLWRNF", writer.String())
	})
}
