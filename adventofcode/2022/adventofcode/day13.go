package adventofcode

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func Problem2022Day13Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()
	pairIndex, sumOfIndices := 1, 0
	for {
		raw1, err := in.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if len(raw1) > 0 && raw1[len(raw1)-1] == '\n' {
			raw1 = raw1[:len(raw1)-1]
		}
		if len(raw1) == 0 {
			continue
		}
		raw2, err := in.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if len(raw2) > 0 && raw2[len(raw2)-1] == '\n' {
			raw2 = raw2[:len(raw2)-1]
		}
		if len(raw2) == 0 {
			continue
		}
		var parse func(s []byte) ([]interface{}, int)
		parse = func(s []byte) ([]interface{}, int) {
			if s[0] != '[' {
				panic("expected [ at the start")
			}
			idx := 1
			result := make([]interface{}, 0)
			for {
				if s[idx] == ']' {
					idx += 1
					break
				}
				if s[idx] == '[' {
					childResult, childEndIdx := parse(s[idx:])
					result = append(result, childResult)
					idx += childEndIdx
				} else if s[idx] >= '0' && s[idx] <= '9' {
					end := idx + 1
					for {
						if s[end] < '0' || s[end] > '9' {
							break
						}
						end++
					}
					v, _ := strconv.Atoi(string(s[idx:end]))
					result = append(result, v)
					idx = end
				} else if s[idx] == ',' {
					idx += 1
				} else {
					panic("unexpected character")
				}
			}
			return result, idx
		}
		input1, _ := parse(raw1)
		input2, _ := parse(raw2)
		var areInTheRightOrder func(lhs, rhs []interface{}) *bool
		areInTheRightOrder = func(lhs, rhs []interface{}) *bool {
			for i := 0; i < len(lhs) || i < len(rhs); i++ {
				if i >= len(lhs) {
					result := true
					return &result
				}
				if i >= len(rhs) {
					result := false
					return &result
				}
				lhsVal, lhsOK := lhs[i].(int)
				rhsVal, rhsOK := rhs[i].(int)
				if lhsOK && rhsOK {
					if lhsVal < rhsVal {
						result := true
						return &result
					}
					if lhsVal > rhsVal {
						result := false
						return &result
					}
				} else if lhsOK {
					result := areInTheRightOrder([]interface{}{lhsVal}, rhs[i].([]interface{}))
					if result != nil {
						return result
					}
				} else if rhsOK {
					result := areInTheRightOrder(lhs[i].([]interface{}), []interface{}{rhsVal})
					if result != nil {
						return result
					}
				} else {
					result := areInTheRightOrder(lhs[i].([]interface{}), rhs[i].([]interface{}))
					if result != nil {
						return result
					}
				}
			}
			return nil
		}
		result := areInTheRightOrder(input1, input2)
		if result != nil && *result {
			sumOfIndices += pairIndex
		}
		pairIndex += 1
	}
	fmt.Fprint(out, sumOfIndices)
}

func Problem2022Day13Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()
	packets := make([][]interface{}, 0)
	for {
		raw1, err := in.ReadBytes('\n')
		if len(raw1) > 0 && raw1[len(raw1)-1] == '\n' {
			raw1 = raw1[:len(raw1)-1]
		}
		if len(raw1) == 0 {
			if err == io.EOF {
				break
			}
			continue
		}
		var parse func(s []byte) ([]interface{}, int)
		parse = func(s []byte) ([]interface{}, int) {
			if s[0] != '[' {
				panic("expected [ at the start")
			}
			idx := 1
			result := make([]interface{}, 0)
			for {
				if s[idx] == ']' {
					idx += 1
					break
				}
				if s[idx] == '[' {
					childResult, childEndIdx := parse(s[idx:])
					result = append(result, childResult)
					idx += childEndIdx
				} else if s[idx] >= '0' && s[idx] <= '9' {
					end := idx + 1
					for {
						if s[end] < '0' || s[end] > '9' {
							break
						}
						end++
					}
					v, _ := strconv.Atoi(string(s[idx:end]))
					result = append(result, v)
					idx = end
				} else if s[idx] == ',' {
					idx += 1
				} else {
					panic("unexpected character")
				}
			}
			return result, idx
		}
		packet, _ := parse(raw1)
		packets = append(packets, packet)
	}

	var areInTheRightOrder func(lhs, rhs []interface{}) *bool
	areInTheRightOrder = func(lhs, rhs []interface{}) *bool {
		for i := 0; i < len(lhs) || i < len(rhs); i++ {
			if i >= len(lhs) {
				result := true
				return &result
			}
			if i >= len(rhs) {
				result := false
				return &result
			}
			lhsVal, lhsOK := lhs[i].(int)
			rhsVal, rhsOK := rhs[i].(int)
			if lhsOK && rhsOK {
				if lhsVal < rhsVal {
					result := true
					return &result
				}
				if lhsVal > rhsVal {
					result := false
					return &result
				}
			} else if lhsOK {
				result := areInTheRightOrder([]interface{}{lhsVal}, rhs[i].([]interface{}))
				if result != nil {
					return result
				}
			} else if rhsOK {
				result := areInTheRightOrder(lhs[i].([]interface{}), []interface{}{rhsVal})
				if result != nil {
					return result
				}
			} else {
				result := areInTheRightOrder(lhs[i].([]interface{}), rhs[i].([]interface{}))
				if result != nil {
					return result
				}
			}
		}
		return nil
	}
	packets = append(packets, []interface{}{[]interface{}{2}})
	packets = append(packets, []interface{}{[]interface{}{6}})
	for i := len(packets) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			lhs, rhs := packets[j], packets[j+1]
			result := areInTheRightOrder(lhs, rhs)
			if result == nil || !*result {
				packets[j] = rhs
				packets[j+1] = lhs
			}
		}
	}
	result := 1
	for i, packet := range packets {
		str := fmt.Sprintf("%v", packet)
		if str == "[[2]]" || str == "[[6]]" {
			result *= (i + 1)
		}
	}
	fmt.Fprint(out, result)
}
