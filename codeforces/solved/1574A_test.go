package codeforces

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem1574A(t *testing.T) {
	input := `
4
1
2
3
4
`
	expected := `
()
()()
(())
()()()
()(())
((()))
()()()()
()()(())
()((()))
(((())))
`

	reader := strings.NewReader(strings.TrimSpace(input))
	writer := &strings.Builder{}

	Problem1574A(reader, writer)

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(writer.String()))
}
