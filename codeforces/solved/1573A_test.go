package codeforces

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	input := `
7
3
007
4
1000
5
00000
3
103
4
2020
9
123456789
30
001678294039710047203946100020		
`
	expected := `
7
2
0
5
6
53
115	
`

	reader := strings.NewReader(strings.TrimSpace(input))
	writer := &strings.Builder{}

	Problem1573A(reader, writer)

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(writer.String()))
}
