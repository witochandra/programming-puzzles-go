package adventofcode

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Problem2022Day14Part1(r io.Reader, w io.Writer) {
	type Point struct {
		X, Y int
	}

	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	var (
		minPoint = Point{500, 0}
		maxPoint = Point{500, 0}
		lines    = make([][]Point, 0)
		terrain  = make([][]int, 0)
	)
	for {
		input, _ := in.ReadString('\n')
		input = strings.TrimSpace(input)
		if len(input) == 0 {
			break
		}
		line := make([]Point, 0)
		lineStrings := strings.Split(input, " -> ")
		for _, lineString := range lineStrings {
			pointStrings := strings.Split(lineString, ",")
			x, _ := strconv.Atoi(pointStrings[0])
			y, _ := strconv.Atoi(pointStrings[1])
			point := Point{x, y}
			line = append(line, point)
			if point.X < minPoint.X {
				minPoint.X = point.X
			} else if point.X > maxPoint.X {
				maxPoint.X = point.X
			}
			if point.Y < minPoint.Y {
				minPoint.Y = point.Y
			} else if point.Y > maxPoint.Y {
				maxPoint.Y = point.Y
			}
		}
		lines = append(lines, line)
	}

	printTerrain := func() {
		result := ""
		for _, row := range terrain {
			for _, cell := range row {
				if cell == -1 {
					result += "#"
				} else if cell == 1 {
					result += "o"
				} else if cell == 2 {
					result += "~"
				} else {
					result += " "
				}
			}
			result += "\n"
		}
		// fmt.Println(result)
	}

	// create terrain
	for y := minPoint.Y; y <= maxPoint.Y; y++ {
		terrain = append(terrain, make([]int, maxPoint.X-minPoint.X+1))
	}
	for _, line := range lines {
		var previousPoint *Point
		for _, point := range line {
			if previousPoint == nil {
				previousPoint = &Point{
					X: point.X,
					Y: point.Y,
				}
				continue
			}
			if previousPoint.X == point.X {
				y := previousPoint.Y
				for {
					terrain[y-minPoint.Y][point.X-minPoint.X] = -1
					if y == point.Y {
						break
					} else if y > point.Y {
						y--
					} else if y < point.Y {
						y++
					}
				}
			} else if previousPoint.Y == point.Y {
				x := previousPoint.X
				for {
					terrain[point.Y-minPoint.Y][x-minPoint.X] = -1
					if x == point.X {
						break
					} else if x > point.X {
						x--
					} else if x < point.X {
						x++
					}
				}
			}
			previousPoint = &Point{
				X: point.X,
				Y: point.Y,
			}
		}
	}

	// drop sands
	done := false
	numberOfSands := 0
	for {
		numberOfSands++
		sandPosition := Point{500, 0}
		for {
			if sandPosition.Y+1 > maxPoint.Y || sandPosition.X-1 < minPoint.X || sandPosition.X+1 > maxPoint.X {
				done = true
				break
			}
			if terrain[sandPosition.Y+1][sandPosition.X-minPoint.X] == 0 {
				sandPosition.Y++
				continue
			}
			if terrain[sandPosition.Y+1][sandPosition.X-minPoint.X-1] == 0 {
				sandPosition.X--
				sandPosition.Y++
				continue
			}
			if terrain[sandPosition.Y+1][sandPosition.X-minPoint.X+1] == 0 {
				sandPosition.X++
				sandPosition.Y++
				continue
			}
			break
		}
		if done {
			terrain[sandPosition.Y][sandPosition.X-minPoint.X] = 2
			break
		}
		terrain[sandPosition.Y][sandPosition.X-minPoint.X] = 1
	}
	printTerrain()
	fmt.Fprint(out, numberOfSands-1)
}

func Problem2022Day14Part2(r io.Reader, w io.Writer) {
	type Point struct {
		X, Y int
	}

	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	var (
		minPoint = Point{500, 0}
		maxPoint = Point{500, 0}
		lines    = make([][]Point, 0)
		terrain  = make([][]int, 0)
	)
	for {
		input, _ := in.ReadString('\n')
		input = strings.TrimSpace(input)
		if len(input) == 0 {
			break
		}
		line := make([]Point, 0)
		lineStrings := strings.Split(input, " -> ")
		for _, lineString := range lineStrings {
			pointStrings := strings.Split(lineString, ",")
			x, _ := strconv.Atoi(pointStrings[0])
			y, _ := strconv.Atoi(pointStrings[1])
			point := Point{x, y}
			line = append(line, point)
			if point.X < minPoint.X {
				minPoint.X = point.X
			} else if point.X > maxPoint.X {
				maxPoint.X = point.X
			}
			if point.Y < minPoint.Y {
				minPoint.Y = point.Y
			} else if point.Y > maxPoint.Y {
				maxPoint.Y = point.Y
			}
		}
		lines = append(lines, line)
	}
	// Add 2 more levels to the terrain
	maxPoint.Y += 2
	lines = append(lines, []Point{
		{minPoint.X, maxPoint.Y},
		{maxPoint.X, maxPoint.Y},
	})

	printTerrain := func() {
		result := ""
		for _, row := range terrain {
			for _, cell := range row {
				if cell == -1 {
					result += "#"
				} else if cell == 1 {
					result += "o"
				} else if cell == 2 {
					result += "~"
				} else {
					result += " "
				}
			}
			result += "\n"
		}
		// fmt.Println(result)
	}

	// create terrain
	for y := minPoint.Y; y <= maxPoint.Y; y++ {
		terrain = append(terrain, make([]int, maxPoint.X-minPoint.X+1))
	}
	for _, line := range lines {
		var previousPoint *Point
		for _, point := range line {
			if previousPoint == nil {
				previousPoint = &Point{
					X: point.X,
					Y: point.Y,
				}
				continue
			}
			if previousPoint.X == point.X {
				y := previousPoint.Y
				for {
					terrain[y-minPoint.Y][point.X-minPoint.X] = -1
					if y == point.Y {
						break
					} else if y > point.Y {
						y--
					} else if y < point.Y {
						y++
					}
				}
			} else if previousPoint.Y == point.Y {
				x := previousPoint.X
				for {
					terrain[point.Y-minPoint.Y][x-minPoint.X] = -1
					if x == point.X {
						break
					} else if x > point.X {
						x--
					} else if x < point.X {
						x++
					}
				}
			}
			previousPoint = &Point{
				X: point.X,
				Y: point.Y,
			}
		}
	}

	// drop sands
	done := false
	numberOfSands := 0
	for {
		numberOfSands++
		sandPosition := Point{500, 0}
		for {
			if sandPosition.X-1 < minPoint.X {
				minPoint.X--
				newTerrain := make([][]int, 0)
				for y, row := range terrain {
					cell := 0
					if y == maxPoint.Y {
						cell = -1
					}
					newTerrain = append(newTerrain, append([]int{cell}, row...))
				}
				terrain = newTerrain
				continue
			}
			if sandPosition.X+1 > maxPoint.X {
				maxPoint.X++
				newTerrain := make([][]int, 0)
				for y, row := range terrain {
					cell := 0
					if y == maxPoint.Y {
						cell = -1
					}
					newTerrain = append(newTerrain, append(row, cell))
				}
				terrain = newTerrain
				continue
			}
			if terrain[sandPosition.Y+1][sandPosition.X-minPoint.X] == 0 {
				sandPosition.Y++
				continue
			}
			if terrain[sandPosition.Y+1][sandPosition.X-minPoint.X-1] == 0 {
				sandPosition.X--
				sandPosition.Y++
				continue
			}
			if terrain[sandPosition.Y+1][sandPosition.X-minPoint.X+1] == 0 {
				sandPosition.X++
				sandPosition.Y++
				continue
			}
			if sandPosition.X == 500 && sandPosition.Y == 0 {
				done = true
			}
			break
		}
		if done {
			terrain[sandPosition.Y][sandPosition.X-minPoint.X] = 2
			break
		}
		terrain[sandPosition.Y][sandPosition.X-minPoint.X] = 1
	}
	printTerrain()
	fmt.Fprint(out, numberOfSands)
}
