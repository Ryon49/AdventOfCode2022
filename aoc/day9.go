package aoc

import (
	"data"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Position struct {
	x, y int
}

func (p Position) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func Day9() {
	lines, _ := data.Load(9)

	// part 1
	count1 := SimulateTail(lines, 2)
	fmt.Println(count1)

	// part 2
	count2 := SimulateTail(lines, 10)
	fmt.Println(count2)
}

func SimulateTail(lines []string, maxLen int) int {
	positions := make([]Position, 0)
	// H = queue[0], T = queue[1]
	positions = append(positions, Position{0, 0})
	positions = append(positions, Position{0, 0})

	visited := make(map[Position]bool)
	// part 1
	for _, line := range lines {
		dir, steps := ParseDirection(line)
		for ; steps > 0; steps -= 1 {
			switch dir {
			case "R":
				positions[0].y += 1
			case "L":
				positions[0].y -= 1
			case "U":
				positions[0].x += 1
			case "D":
				positions[0].x -= 1
			}

			last := positions[len(positions)-1]
			for i := 1; i < len(positions); i += 1 {
				if !IsTouching(positions[i-1], positions[i]) {
					// fmt.Printf("Before i = %d, %s\n", i, positions)
					positions[i].x, positions[i].y = NextPosition(positions[i-1], positions[i])
				}
			}
			if last != positions[len(positions)-1] {
				// fmt.Printf("last %s = %s\n", last, positions[len(positions)-1])
				if len(positions) == maxLen {
					visited[last] = true
					// fmt.Printf("added %s\n", last)
				} else {
					positions = append(positions, last)
				}
			}
			// fmt.Println(positions)
		}
	}
	if maxLen == len(positions) {
		visited[positions[len(positions)-1]] = true
	}
	return len(visited)
}

func ParseDirection(line string) (string, int) {
	split := strings.Split(line, " ")

	steps, _ := strconv.Atoi(split[1])
	return split[0], steps
}

func NextPosition(H, T Position) (int, int) {
	// there are total of 12 possibilities
	switch {
	case H.x == T.x && H.y > T.y:
		// T - H
		return H.x, H.y - 1
	case H.x == T.x && H.y < T.y:
		// H - T
		return H.x, H.y + 1
	case H.y == T.y && H.x > T.x:
		// H
		// -
		// T
		return H.x - 1, H.y
	case H.y == T.y && H.x < T.x:
		// T
		// -
		// H
		return H.x + 1, H.y
	case H.x-T.x == 2 && (H.y-T.y == 1 || H.y-T.y == -1):
		// - H -
		// - - -
		// T - T1
		return H.x - 1, H.y
	case H.x-T.x == -2 && (H.y-T.y == 1 || H.y-T.y == -1):
		// T - T1
		// - - -
		// - H -
		return H.x + 1, H.y
	case H.y-T.y == 2 && (H.x-T.x == 1 || H.x-T.x == -1):
		// T  - -
		// -  - H
		// T1 - -
		return H.x, H.y - 1
	case H.y-T.y == -2 && (H.x-T.x == 1 || H.x-T.x == -1):
		// - - T
		// H - -
		// - - T1
		return H.x, H.y + 1
	case H.x-T.x == 2 && H.y-T.y == 2:
		// - - H
		// - - -
		// T - -
		return H.x - 1, H.y - 1
	case H.x-T.x == 2 && H.y-T.y == -2:
		// H - -
		// - - -
		// - - T
		return H.x - 1, H.y + 1
	case H.x-T.x == -2 && H.y-T.y == 2:
		// T - -
		// - - -
		// - - H
		return H.x + 1, H.y - 1
	case H.x-T.x == -2 && H.y-T.y == -2:
		// - - T
		// - - -
		// H - -
		return H.x + 1, H.y + 1

	default:
		log.Fatal(fmt.Sprintf("Miss a case: %s, %s\n", H, T))

	}
	return 0, 0
}

func IsTouching(H, T Position) bool {
	x := H.x - T.x
	y := H.y - T.y
	if x == 2 || x == -2 || y == 2 || y == -2 {
		return false
	}
	return true
}
