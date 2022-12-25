package adventofcode

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

func Problem2022Day11Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	itemsByMonkey := make([][]int, 0)
	operationByMonkey := make([][]string, 0)
	testDivisorByMonkey := make([]int, 0)
	testTrueByMonkey := make([]int, 0)
	testFalseByMonkey := make([]int, 0)
	inspectionCountByMonkey := make([]int, 0)
	for {
		line, err := in.ReadString('\n')
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			if err == io.EOF {
				break
			}
			continue
		}
		if strings.HasPrefix(line, "Starting items: ") {
			items := make([]int, 0)
			for _, itemStr := range strings.Split(line[16:], ", ") {
				item, _ := strconv.Atoi(itemStr)
				items = append(items, item)
			}
			itemsByMonkey = append(itemsByMonkey, items)
			inspectionCountByMonkey = append(inspectionCountByMonkey, 0)
		} else if strings.HasPrefix(line, "Operation: new = ") {
			operators := strings.Split(line[17:], " ")
			operationByMonkey = append(operationByMonkey, operators)
		} else if strings.HasPrefix(line, "Test: divisible by ") {
			divisor, _ := strconv.Atoi(line[19:])
			testDivisorByMonkey = append(testDivisorByMonkey, divisor)
		} else if strings.HasPrefix(line, "If true: throw to monkey ") {
			toMonkey, _ := strconv.Atoi(line[25:])
			testTrueByMonkey = append(testTrueByMonkey, toMonkey)
		} else if strings.HasPrefix(line, "If false: throw to monkey ") {
			toMonkey, _ := strconv.Atoi(line[26:])
			testFalseByMonkey = append(testFalseByMonkey, toMonkey)
		}
	}

	for round := 1; round <= 20; round++ {
		for monkey := 0; monkey < len(itemsByMonkey); monkey++ {
			for _, item := range itemsByMonkey[monkey] {
				lhs, rhs := 0, 0
				if operationByMonkey[monkey][0] == "old" {
					lhs = item
				} else {
					lhs, _ = strconv.Atoi(operationByMonkey[monkey][0])
				}
				if operationByMonkey[monkey][2] == "old" {
					rhs = item
				} else {
					rhs, _ = strconv.Atoi(operationByMonkey[monkey][2])
				}
				newItem := 0
				if operationByMonkey[monkey][1] == "*" {
					newItem = lhs * rhs
				} else {
					newItem = lhs + rhs
				}
				newItem /= 3
				if newItem%testDivisorByMonkey[monkey] == 0 {
					itemsByMonkey[testTrueByMonkey[monkey]] = append(itemsByMonkey[testTrueByMonkey[monkey]], newItem)
				} else {
					itemsByMonkey[testFalseByMonkey[monkey]] = append(itemsByMonkey[testFalseByMonkey[monkey]], newItem)
				}
				inspectionCountByMonkey[monkey]++
			}
			itemsByMonkey[monkey] = make([]int, 0)
		}
	}
	sort.Ints(inspectionCountByMonkey)
	fmt.Fprint(out, inspectionCountByMonkey[len(inspectionCountByMonkey)-1]*inspectionCountByMonkey[len(inspectionCountByMonkey)-2])
}

func Problem2022Day11Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	itemsByMonkey := make([][]int, 0)
	operationByMonkey := make([][]string, 0)
	testDivisorByMonkey := make([]int, 0)
	testTrueByMonkey := make([]int, 0)
	testFalseByMonkey := make([]int, 0)
	inspectionCountByMonkey := make([]int, 0)
	var gcd func(a, b int) int
	gcd = func(a, b int) int {
		if a == 0 {
			return b
		}
		return gcd(b%a, a)
	}
	lcm := func(arr []int) int {
		ans := arr[0]

		for i := 1; i < len(arr); i++ {
			ans = int((int64(ans) * int64(arr[i])) / int64(gcd(ans, arr[i])))
		}

		return ans
	}

	for {
		line, err := in.ReadString('\n')
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			if err == io.EOF {
				break
			}
			continue
		}
		if strings.HasPrefix(line, "Starting items: ") {
			items := make([]int, 0)
			for _, itemStr := range strings.Split(line[16:], ", ") {
				item, _ := strconv.Atoi(itemStr)
				items = append(items, item)
			}
			itemsByMonkey = append(itemsByMonkey, items)
			inspectionCountByMonkey = append(inspectionCountByMonkey, 0)
		} else if strings.HasPrefix(line, "Operation: new = ") {
			operators := strings.Split(line[17:], " ")
			operationByMonkey = append(operationByMonkey, operators)
		} else if strings.HasPrefix(line, "Test: divisible by ") {
			divisor, _ := strconv.Atoi(line[19:])
			testDivisorByMonkey = append(testDivisorByMonkey, divisor)
		} else if strings.HasPrefix(line, "If true: throw to monkey ") {
			toMonkey, _ := strconv.Atoi(line[25:])
			testTrueByMonkey = append(testTrueByMonkey, toMonkey)
		} else if strings.HasPrefix(line, "If false: throw to monkey ") {
			toMonkey, _ := strconv.Atoi(line[26:])
			testFalseByMonkey = append(testFalseByMonkey, toMonkey)
		}
	}
	monkeyLCM := lcm(testDivisorByMonkey)
	for round := 1; round <= 10000; round++ {
		for monkey := 0; monkey < len(itemsByMonkey); monkey++ {
			for _, item := range itemsByMonkey[monkey] {
				var (
					lhs, rhs int
				)
				if operationByMonkey[monkey][0] == "old" {
					lhs = item
				} else {
					lhs, _ = strconv.Atoi(operationByMonkey[monkey][0])
				}
				if operationByMonkey[monkey][2] == "old" {
					rhs = item
				} else {
					rhs, _ = strconv.Atoi(operationByMonkey[monkey][2])
				}
				var newItem int
				if operationByMonkey[monkey][1] == "*" {
					newItem = lhs * rhs
				} else {
					newItem = lhs + rhs
				}
				newItem %= monkeyLCM
				if newItem%testDivisorByMonkey[monkey] == 0 {
					itemsByMonkey[testTrueByMonkey[monkey]] = append(itemsByMonkey[testTrueByMonkey[monkey]], newItem)
				} else {
					itemsByMonkey[testFalseByMonkey[monkey]] = append(itemsByMonkey[testFalseByMonkey[monkey]], newItem)
				}
				inspectionCountByMonkey[monkey]++
			}
			itemsByMonkey[monkey] = make([]int, 0)
		}
	}
	sort.Ints(inspectionCountByMonkey)
	lhs := int64(inspectionCountByMonkey[len(inspectionCountByMonkey)-1])
	rhs := int64(inspectionCountByMonkey[len(inspectionCountByMonkey)-2])
	fmt.Fprint(out, lhs*rhs)
}
