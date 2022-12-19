package adventofcode

import (
	"bufio"
	"fmt"
	"io"
)

func Problem2022Day05Part1(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	stacks := make([]string, 0)
	{
		for {
			raw, _, _ := in.ReadLine()
			if len(raw) == 0 {
				break
			}
			for i, r := range raw {
				if i%4 != 1 {
					continue
				}
				stackNumber := i / 4
				if len(stacks) <= stackNumber {
					stacks = append(stacks, "")
				}
				if raw[i-1] != '[' || raw[i+1] != ']' {
					continue
				}
				stacks[stackNumber] = string(r) + stacks[stackNumber]
			}
		}
	}
	for {
		number, from, to := 0, 0, 0
		count, err := fmt.Fscanf(in, "move %d from %d to %d\n", &number, &from, &to)
		if count == 0 || err != nil {
			break
		}
		for i := 0; i < number; i++ {
			fromStack := stacks[from-1]
			crateToMove := string(fromStack[len(fromStack)-1])
			stacks[from-1] = fromStack[:len(fromStack)-1]
			stacks[to-1] = stacks[to-1] + crateToMove
		}
	}
	result := ""
	for _, stack := range stacks {
		if len(stack) == 0 {
			result += " "
		} else {
			result += string(stack[len(stack)-1])
		}
	}
	out.WriteString(result)
}

func Problem2022Day05Part2(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	stacks := make([]string, 0)
	// Scan stacks
	{
		for {
			raw, _, _ := in.ReadLine()
			if len(raw) == 0 {
				break
			}
			for i, r := range raw {
				if i%4 != 1 {
					continue
				}
				stackNumber := i / 4
				if len(stacks) <= stackNumber {
					stacks = append(stacks, "")
				}
				if raw[i-1] != '[' || raw[i+1] != ']' {
					continue
				}
				stacks[stackNumber] = string(r) + stacks[stackNumber]
			}
		}
	}

	// Scan & perform actions
	{
		for {
			number, from, to := 0, 0, 0
			count, err := fmt.Fscanf(in, "move %d from %d to %d\n", &number, &from, &to)
			if count == 0 || err != nil {
				break
			}
			fromStack := stacks[from-1]
			stacks[from-1] = fromStack[:len(fromStack)-number]
			stacks[to-1] = stacks[to-1] + fromStack[len(fromStack)-number:]
		}
	}

	result := ""
	for _, stack := range stacks {
		if len(stack) == 0 {
			result += " "
		} else {
			result += string(stack[len(stack)-1])
		}
	}
	out.WriteString(result)
}
