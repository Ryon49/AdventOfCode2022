package aoc

import (
	"data"
	"fmt"
)

type Shape struct {
	coordinates []Position
}

type CacheKey17 struct {
	shapeIndex, jetIndex int
}

type CacheValue17 struct {
	rock, height int
}

func Day17() {
	lines, _ := data.Load(17)
	input := lines[0]

	shapes := []func(int) *Shape{
		func(height int) *Shape {
			// horizontal bar line, left == bottom
			return &Shape{[]Position{{height + 3, 2}, {height + 3, 3}, {height + 3, 4}, {height + 3, 5}}}
		},
		func(height int) *Shape {
			// cross, left == bottom
			return &Shape{[]Position{{height + 4, 2}, {height + 3, 3}, {height + 4, 3}, {height + 5, 3}, {height + 4, 4}}}
		},
		func(height int) *Shape {
			// reflected L, left == bottom
			return &Shape{[]Position{{height + 3, 2}, {height + 3, 3}, {height + 3, 4}, {height + 4, 4}, {height + 5, 4}}}
		},
		func(height int) *Shape {
			// vertical bar line, left == bottom
			return &Shape{[]Position{{height + 3, 2}, {height + 4, 2}, {height + 5, 2}, {height + 6, 2}}}
		},
		func(height int) *Shape {
			// horizontal bar line, left == bottom
			return &Shape{[]Position{{height + 3, 2}, {height + 4, 2}, {height + 3, 3}, {height + 4, 3}}}
		},
	}

	shapeIndex, jetIndex := 0, 0
	height, numRocks := 0, 0

	rocks := make(map[Position]bool)
	cache := make(map[CacheKey17]CacheValue17)

	shape := shapes[shapeIndex](height)

	totalRocks := 1000000000000

	for numRocks < totalRocks {
		c := input[jetIndex]
		// fmt.Printf("%c\n", c)
		if c == '<' {
			shape.ShapeMoveLeft(rocks)
		} else {
			shape.ShapeMoveRight(rocks)
		}

		if !shape.ShapeMoveDown(rocks) {
			for _, pos := range shape.coordinates {
				rocks[pos] = true
			}

			if shape.Top() > height {
				height = shape.Top()
			}
			numRocks += 1

			key := CacheKey17{shapeIndex, jetIndex}
			if value, found := cache[key]; found {
				previousRock, previousHeight := value.rock, value.height

				cycleLength := numRocks - previousRock
				cycleHeight := height - previousHeight

				if numRocks%cycleLength == totalRocks%cycleLength {
					fmt.Println("found a cycle")
					rocksRemaining := totalRocks - numRocks
					cyclesRemaining := (rocksRemaining / cycleLength)

					last := cyclesRemaining*cycleHeight + height
					fmt.Printf("After %d rocks, heigt = %d", totalRocks, last)
					// 1572093023267
					return
				}
			} else {
				cache[key] = CacheValue17{numRocks, height}
			}

			shapeIndex = (shapeIndex + 1) % 5
			shape = shapes[shapeIndex](height)
			// PrintGridReverse(rocks, height+3)
		}

		jetIndex = (jetIndex + 1) % len(input)
	}
	fmt.Println(height)
}

func PrintGridReverse(rocks map[Position]bool, upto int) {
	for i := upto; i >= 0; i -= 1 {
		fmt.Printf("%d\t ", i)
		for j := 0; j < 7; j += 1 {
			pos := Position{i, j}
			if _, found := rocks[pos]; found {
				fmt.Printf("%c", '#')
			} else {
				fmt.Printf("%c", '.')
			}
		}
		fmt.Println()
	}
}

func (shape *Shape) ShapeMoveLeft(grid map[Position]bool) {
	for _, pos := range shape.coordinates {
		nextPos := Position{pos.x, pos.y - 1}

		if _, found := grid[nextPos]; pos.y == 0 || found {
			return
		}
	}

	for i, _ := range shape.coordinates {
		shape.coordinates[i].y -= 1
	}
}

func (shape *Shape) ShapeMoveRight(grid map[Position]bool) {
	for _, pos := range shape.coordinates {
		nextPos := Position{pos.x, pos.y + 1}
		if _, found := grid[nextPos]; pos.y == 6 || found {
			return
		}
	}

	for i, _ := range shape.coordinates {
		shape.coordinates[i].y += 1
	}
}

func (shape *Shape) ShapeMoveDown(grid map[Position]bool) bool {
	for _, pos := range shape.coordinates {
		nextPos := Position{pos.x - 1, pos.y}
		if _, found := grid[nextPos]; pos.x == 0 || found {
			return false
		}
	}

	for i, _ := range shape.coordinates {
		shape.coordinates[i].x -= 1
	}
	return true
}

func (shape *Shape) Top() int {
	top := 0
	for _, pos := range shape.coordinates {
		if pos.x > top {
			top = pos.x
		}
	}
	return top + 1
}
