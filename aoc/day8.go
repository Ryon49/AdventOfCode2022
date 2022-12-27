package aoc

import (
	"data"
	"fmt"
	"strconv"
	"strings"
)

func Day8() {
	lines, _ := data.Load(8)

	grid := ParseGrid(lines)

	// part 1
	visible := 2*len(grid[0]) + 2*(len(grid)-2) // edge
	for i := 1; i < len(grid)-1; i += 1 {
		for j := 1; j < len(grid[0])-1; j += 1 {
			if IsVisible(grid, i, j) {
				visible += 1
			}
		}
	}
	fmt.Printf("View at most %d trees\n", visible)

	highestScore := 0
	for i := 1; i < len(grid)-1; i += 1 {
		for j := 1; j < len(grid[0])-1; j += 1 {
			score := FindScore(grid, i, j)
			if score > highestScore {
				highestScore = score
			}
		}
	}
	fmt.Printf("Highest score = %d\n", highestScore)

}

func IsVisible(grid [][]int, x, y int) bool {

	var (
		visiableN = true
		visiableS = true
		visiableE = true
		visiableW = true
	)
	// check north
	for i := x - 1; i >= 0; i -= 1 {
		if grid[x][y] <= grid[i][y] {
			visiableN = false
			break
		}
	}

	// check south
	for i := x + 1; i < len(grid); i += 1 {
		if grid[x][y] <= grid[i][y] {
			visiableS = false
			break
		}
	}

	// check east
	for j := y - 1; j >= 0; j -= 1 {
		if grid[x][y] <= grid[x][j] {
			visiableE = false
			break
		}
	}

	// check west
	for j := y + 1; j < len(grid[0]); j += 1 {
		if grid[x][y] <= grid[x][j] {
			visiableW = false
			break
		}
	}
	return visiableN || visiableS || visiableE || visiableW
}

func FindScore(grid [][]int, x, y int) int {
	var score int = 1
	var visible int = 0

	// check north
	for i := x - 1; i >= 0; i -= 1 {
		visible += 1
		if grid[x][y] <= grid[i][y] {
			break
		}
	}
	score, visible = score*visible, 0

	// check south
	for i := x + 1; i < len(grid); i += 1 {
		visible += 1
		if grid[x][y] <= grid[i][y] {
			break
		}
	}
	score, visible = score*visible, 0

	// check east
	for j := y - 1; j >= 0; j -= 1 {
		visible += 1
		if grid[x][y] <= grid[x][j] {
			break
		}
	}
	score, visible = score*visible, 0

	// check west
	for j := y + 1; j < len(grid[0]); j += 1 {
		visible += 1
		if grid[x][y] <= grid[x][j] {
			break
		}
	}
	score, visible = score*visible, 0
	if x == 1 && y == 2 {
		fmt.Println(score)
	}
	return score
}

func ParseGrid(lines []string) [][]int {
	row, col := len(lines), len(lines[0])

	var grid [][]int = make([][]int, row)
	for i, _ := range grid {
		grid[i] = make([]int, col)
	}

	for i := 0; i < row; i += 1 {
		line := strings.Split(lines[i], "")
		for j := 0; j < col; j += 1 {
			grid[i][j], _ = strconv.Atoi(line[j])
		}
	}
	return grid
}
