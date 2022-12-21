package adventofcode

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem2022Day07_Example(t *testing.T) {
	input := strings.TrimSpace(`
$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`)
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day07Part1(strings.NewReader(input), writer)

		assert.Equal(t, "95437", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day07Part2(strings.NewReader(input), writer)

		assert.Equal(t, "24933642", writer.String())
	})
}

func TestProblem2022Day07_Puzzle(t *testing.T) {
	fileContent, err := os.ReadFile("./day07_input.txt")
	if err != nil {
		t.Error(err)
		return
	}
	input := strings.TrimSpace(string(fileContent))
	t.Run("part 1", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day07Part1(strings.NewReader(input), writer)

		assert.Equal(t, "1427048", writer.String())
	})
	t.Run("part 2", func(t *testing.T) {
		writer := &strings.Builder{}

		Problem2022Day07Part2(strings.NewReader(input), writer)

		assert.Equal(t, "TZLTLWRNF", writer.String())
	})
}
