package fhc_qualification_2020

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblemA_Small(t *testing.T) {
	input := `
5
2
YY
YY
2
NY
YY
2
NN
YY
5
YNNYY
YYYNY
10
NYYYNNYYYY
YYNYYNYYNY
`

	expected := `
Case #1:
YY
YY
Case #2:
YY
NY
Case #3:
YN
NY
Case #4:
YNNNN
YYNNN
NNYYN
NNNYN
NNNYY
Case #5:
YYYNNNNNNN
NYYNNNNNNN
NNYNNNNNNN
NNYYNNNNNN
NNYYYNNNNN
NNNNNYNNNN
NNNNNNYYYN
NNNNNNYYYN
NNNNNNNNYN
NNNNNNNNYY
`

	inputReader := strings.NewReader(strings.TrimSpace(input))
	output := &strings.Builder{}

	ProblemA(inputReader, output)
	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(output.String()))
}
