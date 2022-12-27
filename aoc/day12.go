package aoc

import (
	"data"
	"fmt"
)

type Path struct {
	pos  Position
	step int
}

var Directions []Position = []Position{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func Day12() {
	lines, _ := data.Load(12)

	S, E := FindSE(lines)

	steps := SearchGrid(lines, E, S)
	fmt.Println(steps)
}

func SearchGrid(grid []string, S, E Position) int {
	visited := make(map[Position]bool)

	visited[S] = true

	queue := []Path{{S, 0}}
	for {
		for len(queue) > 0 {
			curPath := queue[0]
			queue = queue[1:]

			curValue := GridValue(grid, curPath.pos, S, E)
			// Part 1
			// if curPath.pos == E {
			// 	return curPath.step
			// }

			// Part 2
			if GridValue(grid, curPath.pos, S, E) == 'a' {
				return curPath.step
			}
			for _, dir := range Directions {
				nextPos := Position{curPath.pos.x + dir.x, curPath.pos.y + dir.y}
				if nextPos.x < 0 || nextPos.x >= len(grid) || nextPos.y < 0 || nextPos.y >= len(grid[0]) {
					continue
				}

				if _, found := visited[nextPos]; found {
					continue
				}
				nextValue := GridValue(grid, nextPos, S, E)
				if curValue-nextValue == 1 {
					// fmt.Printf("Add %s (%c)\n", nextPos, nextValue)
					queue = append([]Path{{nextPos, curPath.step + 1}}, queue...)
					visited[nextPos] = true
				} else if curValue <= nextValue {
					// fmt.Printf("Add %s (%c)\n", nextPos, nextValue)
					queue = append(queue, Path{nextPos, curPath.step + 1})
					visited[nextPos] = true
				}
			}
		}
	}
}

func GridValue(grid []string, pos, S, E Position) byte {
	if pos == S {
		return 'z'
	} else if pos == E {
		return 'a'
	} else {
		return grid[pos.x][pos.y]
	}
}

func FindSE(grid []string) (S, E Position) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'S' {
				S = Position{i, j}
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'E' {
				E = Position{i, j}
			}
		}
	}
	return
}
