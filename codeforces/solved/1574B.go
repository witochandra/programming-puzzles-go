package codeforces

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

func Problem1574B(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	t := 0
	fmt.Fscanln(in, &t)
	for tcNo := 1; tcNo <= t; tcNo++ {
		a, b, c, m := 0, 0, 0, 0
		fmt.Fscanln(in, &a, &b, &c, &m)

		minM := 0
		lettersCount := []int{a, b, c}
		sort.Ints(lettersCount)

		minM = lettersCount[2] - lettersCount[1] - lettersCount[0] - 1
		if minM < 0 {
			minM = 0
		}

		maxM := 0

		if a > 1 {
			maxM += a - 1
		}
		if b > 1 {
			maxM += b - 1
		}
		if c > 1 {
			maxM += c - 1
		}

		if m >= minM && m <= maxM {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

// func main() {
// 	Problem1574B(os.Stdin, os.Stdout)
// }
