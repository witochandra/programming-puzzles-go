package codeforces

import (
	"bufio"
	"fmt"
	"io"
)

func Problem1569B(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	t := 0
	fmt.Fscanln(in, &t)
	for tcNo := 1; tcNo <= t; tcNo++ {
		numberOfPlayers := 0
		fmt.Fscanln(in, &numberOfPlayers)

		playerTypes := ""
		fmt.Fscanln(in, &playerTypes)

		impossible := false
		tournament := make([][]rune, numberOfPlayers)
		for i := 0; i < numberOfPlayers; i++ {
			tournament[i] = make([]rune, numberOfPlayers)
			tournament[i][i] = 'X'
		}
		queue := make([][]int, 0)
		queue = append(queue, []int{0, 0})
		for len(queue) > 0 {
			reset := false
			start := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			for i := start[0]; i < numberOfPlayers; i++ {
				winCount := 0
				for j := i + 1; j < numberOfPlayers; j++ {
					if i == start[0] && j < start[1] {
						if tournament[i][j] == '+' {
							winCount++
						}
						continue
					}
					if playerTypes[i] == '1' || playerTypes[j] == '1' {
						tournament[i][j] = '='
						tournament[j][i] = '='
					} else if playerTypes[i] == '2' {
						if winCount > 0 || (i == start[0] && j == start[1]) {
							tournament[i][j] = '-'
							tournament[j][i] = '+'
						} else {
							queue = append(queue, []int{i, j})
							tournament[i][j] = '+'
							tournament[j][i] = '-'
							winCount++
						}
					}
				}
				winCount = 0
				loseCount := 0
				for j := 0; j < numberOfPlayers; j++ {
					switch tournament[i][j] {
					case '+':
						winCount++
					case '-':
						loseCount++
					}
				}
				acceptable := (playerTypes[i] == '1' && loseCount == 0) ||
					(playerTypes[i] != '1' && winCount > 0)
				if !acceptable {
					reset = true
					break
				}
			}
			if !reset {
				break
			} else if reset && len(queue) == 0 {
				impossible = true
				break
			}
		}
		if impossible {
			fmt.Fprintln(out, "NO")
		} else {
			fmt.Fprintln(out, "YES")
			for i := 0; i < numberOfPlayers; i++ {
				fmt.Fprintln(out, string(tournament[i]))
			}
		}
	}
}

// func main() {
// 	Problem1569B(os.Stdin, os.Stdout)
// }
