package codeforces

import (
	"bufio"
	"fmt"
	"io"
)

func Problem1569A(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	t := 0
	fmt.Fscanln(in, &t)

	for tcNo := 0; tcNo < t; tcNo++ {
		n := 0
		s := ""
		fmt.Fscanln(in, &n)
		fmt.Fscanln(in, &s)

		countA := 0
		countB := 0
		l := -1
		r := -1
		for i, c := range s {
			switch c {
			case 'a':
				countA++
			case 'b':
				countB++
			}
			if countA == 0 || countB == 0 {
				continue
			}
			if countA > 0 && countB > 0 {
				l = i
				r = i + 1
				break
			}
		}
		fmt.Fprintf(out, "%d %d\n", l, r)
	}
}

// func main() {
// 	Problem1569A(os.Stdin, os.Stdout)
// }
