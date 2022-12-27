package aoc

import (
	"data"
	"fmt"
	"strconv"
)

func Day22() {
	lines, _ := data.Load(22)

	grid, current := CreateGrid22(lines[:len(lines)-2])
	fmt.Println(current, len(grid), len(grid[0]))
	// PrintGrid(grid)

	directions := []Position{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	instructions := lines[len(lines)-1]

	var (
		steps     int
		faceIndex int
		nextFace  int
	)

	for i := 0; i < len(instructions); {
		steps, i = NextInt(instructions, i)
		face := directions[faceIndex]
		for ; steps > 0; steps -= 1 {
			nextPosition := NextPosition22(grid, current, face)
			if current == nextPosition {
				break
			}
			current = nextPosition
		}

		nextFace, i = NextFace(instructions, i)
		if nextFace == 0 {
			break
		}
		faceIndex += nextFace
		if faceIndex == 4 {
			faceIndex = 0
		} else if faceIndex == -1 {
			faceIndex = 3
		}

		// if faceIndex == 0 {
		// 	grid[current.x][current.y] = '>'
		// } else if faceIndex == 1 {
		// 	grid[current.x][current.y] = 'v'
		// } else if faceIndex == 2 {
		// 	grid[current.x][current.y] = '<'
		// } else if faceIndex == 3 {
		// 	grid[current.x][current.y] = '^'
		// }
		// PrintGrid(grid)
		// grid[at.x][at.y] = '.'
		// fmt.Println()
	}
	fmt.Printf("final password = %d\n", 1000*(current.x+1)+4*(current.y+1)+faceIndex)
}

func NextPosition22(grid [][]byte, current, face Position) Position {
	height, width := len(grid), len(grid[0])

	if face.x != 0 {
		nextX := current.x + face.x
		for {
			if nextX == height {
				nextX = 0
			} else if nextX == -1 {
				nextX = height - 1
			} else if grid[nextX][current.y] == ' ' {
				nextX += face.x
				if nextX == current.x {
					fmt.Println("Something is wrong in x + face.x")
				}
			} else {
				break
			}
		}
		if grid[nextX][current.y] == '#' {
			return current
		} else {
			return Position{nextX, current.y}
		}
	} else {
		nextY := current.y + face.y
		for {
			if nextY == width {
				nextY = 0
			} else if nextY == -1 {
				nextY = width - 1
			} else if grid[current.x][nextY] == ' ' {
				nextY += face.y
				if nextY == current.y {
					fmt.Println("Something is wrong in y + face.y")
				}
			} else {
				break
			}
		}

		if grid[current.x][nextY] == '#' {
			return current
		}
		return Position{current.x, nextY}
	}

}

func NextInt(line string, i int) (int, int) {
	end := i
	for end < len(line) && line[end] >= '0' && line[end] <= '9' {
		end += 1
	}
	value, _ := strconv.Atoi(line[i:end])
	return value, end
}

func NextFace(line string, i int) (int, int) {
	if i == len(line) {
		return 0, 0
	}
	if line[i] == 'R' {
		return 1, i + 1
	} else {
		return -1, i + 1
	}
}

func CreateGrid22(lines []string) ([][]byte, Position) {
	// Calculate grid's height and width
	height, width := len(lines), 0
	for _, line := range lines {
		if len(line) > width {
			width = len(line)
		}
	}

	// create an empty grid with walls
	row := []byte{}
	for j := 0; j < width; j += 1 {
		row = append(row, ' ')
	}
	grid := [][]byte{}
	for i := 0; i < height; i += 1 {
		grid = append(grid, []byte{})
		grid[i] = append(grid[i], row...)
	}

	for i, line := range lines {
		for j, c := range line {
			if c == '.' {
				// open path
				grid[i][j] = '.'
			} else if c == '#' {
				grid[i][j] = '#'
			}
		}
	}

	// Calculate beginning position
	var y int
	for j := 0; j < width; j += 1 {
		if grid[0][j] == '.' {
			y = j
			break
		}
	}

	return grid, Position{0, y}
}
