package adventofcode

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Problem2021Day04(r io.Reader, w io.Writer, findLastWinning bool) {
	type Board struct {
		Numbers      []int
		MarksByIndex map[int]bool
	}

	createBoard := func(numbers []int) *Board {
		board := &Board{
			Numbers:      numbers,
			MarksByIndex: make(map[int]bool),
		}
		return board
	}

	mark := func(board *Board, number int) bool {
		for i, n := range board.Numbers {
			if n == number {
				board.MarksByIndex[i] = true
				return true
			}
		}
		return false
	}

	isBoardWinning := func(board *Board) bool {
		if len(board.MarksByIndex) < 5 {
			return false
		}
		// check whether there is a completed row
		for _, i := range []int{0, 5, 10, 15, 20} {
			count := 0
			for j := 0; j < 5; j++ {
				if !board.MarksByIndex[i+j] {
					break
				}
				count++
			}
			if count == 5 {
				return true
			}
		}
		// check whether there is a completed column
		for i := 0; i < 5; i++ {
			count := 0
			for j := 0; j < 5; j++ {
				if !board.MarksByIndex[(j*5)+i] {
					break
				}
				count++
			}
			if count == 5 {
				return true
			}
		}
		return false
	}

	calculateScore := func(board *Board, multiplier int) int {
		sum := 0
		for i, n := range board.Numbers {
			if board.MarksByIndex[i] {
				continue
			}
			sum += n
		}
		return sum * multiplier
	}

	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	incomingNumbers := func() []int {
		numberString := ""
		fmt.Fscanln(in, &numberString)
		numbers := make([]int, 0)
		for _, v := range strings.Split(numberString, ",") {
			n, _ := strconv.Atoi(v)
			numbers = append(numbers, n)
		}
		return numbers
	}()

	boards := make([]*Board, 0)
	for {
		_, err := in.Peek(1)
		if err != nil {
			break
		}
		fmt.Fscanln(in)
		numbersInBoard := make([]int, 0)
		for i := 0; i < 5; i++ {
			numbers := make([]int, 5)
			fmt.Fscanln(in, &numbers[0], &numbers[1], &numbers[2], &numbers[3], &numbers[4])
			numbersInBoard = append(numbersInBoard, numbers...)
		}
		boards = append(boards, createBoard(numbersInBoard))
	}

	score := 0
	findFirstWinningBoardScore := func() int {
		for _, currentNumber := range incomingNumbers {
			for _, board := range boards {
				mark(board, currentNumber)
				if isBoardWinning(board) {
					return calculateScore(board, currentNumber)
				}
			}
		}
		return -1
	}
	findLastWinningBoardScore := func() int {
		winningBoardByIndex := make(map[int]bool)
		for _, currentNumber := range incomingNumbers {
			for i, board := range boards {
				if winningBoardByIndex[i] {
					continue
				}
				mark(board, currentNumber)
				if isBoardWinning(board) {
					winningBoardByIndex[i] = true
					if len(winningBoardByIndex) == len(boards) {
						return calculateScore(board, currentNumber)
					}
				}
			}
		}
		return -1
	}
	if findLastWinning {
		score = findLastWinningBoardScore()
	} else {
		score = findFirstWinningBoardScore()
	}
	fmt.Fprintln(out, score)
}

func Problem2021Day04Part1(r io.Reader, w io.Writer) {
	Problem2021Day04(r, w, false)
}

func Problem2021Day04Part2(r io.Reader, w io.Writer) {
	Problem2021Day04(r, w, true)
}
