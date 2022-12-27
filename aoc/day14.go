package aoc

import (
	"data"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day14() {
	lines, _ := data.Load(14)

	// part 1
	grid, minY := CreateGrid14(lines, true)

	// mark the source of the sand
	source := Position{0, ProjectColumn(500, minY, true)}
	grid[source.x][source.y] = '+'

	numRested := 0
	for {
		x, y, valid := DropLocation(grid, source)
		if !valid {
			break
		}
		numRested += 1
		grid[x][y] = 'o'
	}

	fmt.Printf("Part 1, umRested = %d\n", numRested)

	// part 1
	grid, minY = CreateGrid14(lines, false)

	// mark the source of the sand
	source = Position{0, ProjectColumn(500, minY, false)}
	grid[source.x][source.y] = '+'

	numRested = 0
	for {
		x, y, _ := DropLocation(grid, source)
		numRested += 1
		if x == source.x && y == source.y {
			break
		}
		grid[x][y] = 'o'
	}

	fmt.Printf("Part 2, numRested = %d\n", numRested)

	// PrintGrid(grid)

}

func DropLocation(grid [][]byte, source Position) (int, int, bool) {
	x, y := source.x, source.y
	// fall straight down
	for grid[x+1][y] == '-' {
		x += 1
	}

	if grid[x+1][y-1] == '-' {
		// check down left
		return DropLocation(grid, Position{x + 1, y - 1})
	} else if grid[x+1][y+1] == '-' {
		// check down right
		return DropLocation(grid, Position{x + 1, y + 1})
	} else if grid[x+1][y-1] == '~' || grid[x+1][y+1] == '~' {
		return x, y, false
	} else {
		// stay still then
		return x, y, true
	}
}

func CreateGrid14(lines []string, endless bool) ([][]byte, int) {
	maxX, minY, maxY := 0, 1000, 0

	// preprocess, find range of grid
	for _, line := range lines {
		_, highX, lowY, highY := ParseLine14(line)
		if highX > maxX {
			maxX = highX
		}
		if lowY < minY {
			minY = lowY
		}
		if highY > maxY {
			maxY = highY
		}
	}

	// fmt.Println(0, maxX, minY, maxY)

	// create an empty grid
	grid := [][]byte{}

	if endless {
		// a single row, padding with void '~'
		row := []byte{}
		row = append(row, '~')
		for col := 0; col < maxY-minY+1; col++ {
			row = append(row, '-')
		}
		row = append(row, '~')
		// add body
		for i := 0; i <= maxX; i += 1 {
			grid = append(grid, []byte{})
			grid[i] = append(grid[i], row...)
		}

		// as an illustration, the bottom row is marked as an endles abyss
		abyss := []byte{}
		for col := 0; col < maxY-minY+3; col++ {
			abyss = append(abyss, '~')
		}
		grid = append(grid, abyss)
	} else {
		// We add two additional row below,
		maxX += 2
		minY = 0
		maxY += maxY

		row := []byte{}
		for col := 0; col < maxY-minY+1; col++ {
			row = append(row, '-')
		}
		for i := 0; i <= maxX; i += 1 {
			grid = append(grid, []byte{})
			grid[i] = append(grid[i], row...)
		}
		lines = append(lines, fmt.Sprintf("%d,%d -> %d,%d", minY, maxX, maxY, maxX))
	}

	for _, line := range lines {
		positions, _, _, _ := ParseLine14(line)
		for i := 1; i < len(positions); i++ {
			from, to := positions[i-1], positions[i]

			switch {
			case from.x == to.x && from.y < to.y:
				// same row, from --- to
				for y := from.y; y <= to.y; y += 1 {
					grid[from.x][ProjectColumn(y, minY, endless)] = '#'
				}
			case from.x == to.x && from.y > to.y:
				// same row, to --- fromt
				for y := to.y; y <= from.y; y += 1 {
					grid[from.x][ProjectColumn(y, minY, endless)] = '#'
				}
			case from.x < to.x && from.y == to.y:
				// from  same column
				// to
				y := ProjectColumn(from.y, minY, endless)
				for x := from.x; x <= to.x; x += 1 {
					grid[x][y] = '#'
				}
			case from.x > to.x && from.y == to.y:
				// to  same column
				// from
				y := ProjectColumn(from.y, minY, endless)
				for x := to.x; x <= from.x; x += 1 {
					grid[x][y] = '#'
				}
			}
		}
	}
	return grid, minY
}

func ProjectColumn(y, minY int, endless bool) int {
	if endless {
		return y - minY + 1
	} else {
		return y - minY
	}
}

// return
// 1. list of coordinate forming the rock
// 2. min x
// 3. max x
// 4. max y
func ParseLine14(line string) ([]Position, int, int, int) {
	pattern, _ := regexp.Compile("(\\d+,\\d+)")

	pos := []Position{}
	maxX, minY, maxY := 0, 1000, 0

	matches := pattern.FindAllString(line, -1)
	for _, m := range matches {
		split := strings.Split(m, ",")
		col, _ := strconv.Atoi(split[0])
		row, _ := strconv.Atoi(split[1])
		pos = append(pos, Position{row, col})
		if row > maxX {
			maxX = row
		}
		if col < minY {
			minY = col
		}
		if col > maxY {
			maxY = col
		}
	}
	return pos, maxX, minY, maxY
}

func PrintGrid(grid [][]byte) {
	for i := 0; i < len(grid); i += 1 {
		fmt.Printf("%d\t ", i)
		for j := 0; j < len(grid[0]); j += 1 {
			fmt.Printf("%c", grid[i][j])
		}
		fmt.Println()
	}
}
