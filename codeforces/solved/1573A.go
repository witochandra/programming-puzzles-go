package codeforces

import (
	"bufio"
	"fmt"
	"io"
)

func Problem1573A(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	t := 0
	fmt.Fscanln(in, &t)
	for tcNo := 1; tcNo <= t; tcNo++ {
		n := 0
		fmt.Fscanln(in, &n)

		s := ""
		fmt.Fscanln(in, &s)

		ops := 0
		for i, r := range s {
			if r == '0' {
				continue
			}
			ops = ops + int(r-'0')
			if i < len(s)-1 {
				ops++
			}
		}
		fmt.Fprintln(out, ops)
	}
}

// func main() {
// 	Problem1573A(os.Stdin, os.Stdout)
// }
