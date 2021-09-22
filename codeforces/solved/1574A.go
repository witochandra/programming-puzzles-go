package codeforces

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Problem1574A(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	t := 0
	fmt.Fscanln(in, &t)

	for tcNo := 1; tcNo <= t; tcNo++ {
		n := 0
		fmt.Fscanln(in, &n)

		result := make([]string, 0)
		for i := 0; i < n; i++ {
			temp := strings.Repeat("()", n-i-1) + strings.Repeat("(", i+1) + strings.Repeat(")", i+1)
			result = append(result, temp)
		}
		fmt.Fprintln(out, strings.Join(result, "\n"))
	}
}

// func main() {
// 	Problem1574A(os.Stdin, os.Stdout)
// }
