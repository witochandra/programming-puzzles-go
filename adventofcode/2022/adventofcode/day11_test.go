package adventofcode

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2022Day11_Example(t *testing.T) {
	input := strings.TrimSpace(`
  Monkey 0:
	Starting items: 79, 98
	Operation: new = old * 19
	Test: divisible by 23
	  If true: throw to monkey 2
	  If false: throw to monkey 3
  
  Monkey 1:
	Starting items: 54, 65, 75, 74
	Operation: new = old + 6
	Test: divisible by 19
	  If true: throw to monkey 2
	  If false: throw to monkey 0
  
  Monkey 2:
	Starting items: 79, 60, 97
	Operation: new = old * old
	Test: divisible by 13
	  If true: throw to monkey 1
	  If false: throw to monkey 3
  
  Monkey 3:
	Starting items: 74
	Operation: new = old + 3
	Test: divisible by 17
	  If true: throw to monkey 0
	  If false: throw to monkey 1
`)
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day11Part1(strings.NewReader(input), writer)

		assert.Equal(t, "10605", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day11Part2(strings.NewReader(input), writer)

		assert.Equal(t, "2713310158", writer.String())
	})
}

func TestProblem2022Day11_Puzzle(t *testing.T) {
	fileContent, err := os.ReadFile("./day11_input.txt")
	if err != nil {
		t.Error(err)
		return
	}
	input := strings.TrimSpace(string(fileContent))
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day11Part1(strings.NewReader(input), writer)

		assert.Equal(t, "117640", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day11Part2(strings.NewReader(input), writer)

		assert.Equal(t, "30616425600", writer.String())
	})
}
