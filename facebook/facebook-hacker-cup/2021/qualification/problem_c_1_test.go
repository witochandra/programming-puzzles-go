package fhc_qualification_2021

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblemC1(t *testing.T) {
	input := `
6
2
10 20
1 2
4
1 1 1 1
2 1
4 2
1 3
4
1 1 1 1
2 1
4 1
1 3
6
5 4 1 3 2 4
5 1
5 4
5 3
5 2
6 3
9
2 14 7 6 11 3 6 1 8
4 5
6 7
8 9
1 3
6 8
2 4
4 1
1 8
1
10
`
	expected := `
Case #1: 30
Case #2: 4
Case #3: 3
Case #4: 12
Case #5: 32
Case #6: 10
`

	reader := strings.NewReader(strings.TrimSpace(input))
	writer := &strings.Builder{}

	fmt.Println(strings.TrimSpace(input))
	ProblemC1(reader, writer)

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(writer.String()))
}
