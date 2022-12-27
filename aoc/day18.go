package aoc

import (
	"data"
	"fmt"
	"strconv"
	"strings"
)

type Coordinate struct {
	x, y, z int
}

func (c Coordinate) NeighborX() (Coordinate, Coordinate) {
	return Coordinate{c.x - 1, c.y, c.z}, Coordinate{c.x + 1, c.y, c.z}
}

func (c Coordinate) NeighborY() (Coordinate, Coordinate) {
	return Coordinate{c.x, c.y - 1, c.z}, Coordinate{c.x, c.y + 1, c.z}
}

func (c Coordinate) NeighborZ() (Coordinate, Coordinate) {
	return Coordinate{c.x, c.y, c.z - 1}, Coordinate{c.x, c.y, c.z + 1}
}

func Day18() {
	lines, _ := data.Load(18)

	cubes := make(map[Coordinate]bool)

	faces := 0

	minX, maxX, minY, maxY, minZ, maxZ := 100, 0, 100, 0, 100, 0

	for _, line := range lines {
		coord := ParseLine18(line)

		if coord.x < minX {
			minX = coord.x
		}
		if coord.x > maxX {
			maxX = coord.x
		}

		if coord.y < minY {
			minY = coord.y
		}
		if coord.y > maxY {
			maxY = coord.y
		}

		if coord.z < minZ {
			minZ = coord.z
		}
		if coord.z > maxZ {
			maxZ = coord.z
		}

		faces += 6
		cubes[coord] = true

		// check x coordinate
		leftX, rightX := coord.NeighborX()
		if _, found := cubes[leftX]; found {
			faces -= 2
		}
		if _, found := cubes[rightX]; found {
			faces -= 2
		}

		// check ycoordinate
		leftY, rightY := coord.NeighborY()
		if _, found := cubes[leftY]; found {
			faces -= 2
		}
		if _, found := cubes[rightY]; found {
			faces -= 2
		}

		// check ycoordinate
		leftZ, rightZ := coord.NeighborZ()
		if _, found := cubes[leftZ]; found {
			faces -= 2
		}
		if _, found := cubes[rightZ]; found {
			faces -= 2
		}
	}

	fmt.Println(minX, maxX, minY, maxY, minZ, maxZ)

	// bound := func(c Coordinate) bool {
	// 	return c.IsBounded(minX, maxX, minY, maxY, minZ, maxZ)
	// }

	// trapped := make(map[Coordinate]bool)
	// air := 0
	// // count the cubes that are trapped
	// for i := minX + 1; i <= maxX-1; i += 1 {
	// 	for j := minY + 1; j <= maxY-1; j += 1 {
	// 		for k := minZ + 1; k <= maxZ-1; k += 1 {
	// 			coord := Coordinate{i, j, k}
	// 			if _, found := cubes[coord]; found {
	// 				continue
	// 			}

	// 			// check if coord is trapped within cubes

	// 			// check x coordinate
	// 			// check ycoordinate
	// 			// check ycoordinate

	// 			fmt.Println("Air = ", coord)
	// 			trapped[coord] = true
	// 			air += 6
	// 			// check x coordinate
	// 			leftX, rightX := coord.NeighborX()
	// 			if _, found := trapped[leftX]; bound(leftX) && found {
	// 				air -= 2
	// 			}
	// 			if _, found := trapped[rightX]; bound(rightX) && found {
	// 				air -= 2
	// 			}

	// 			// check ycoordinate
	// 			leftY, rightY := coord.NeighborY()
	// 			if _, found := trapped[leftY]; bound(leftY) && found {
	// 				air -= 2
	// 			}
	// 			if _, found := trapped[rightY]; bound(rightY) && found {
	// 				air -= 2
	// 			}

	// 			// check ycoordinate
	// 			leftZ, rightZ := coord.NeighborZ()
	// 			if _, found := trapped[leftZ]; bound(leftZ) && found {
	// 				air -= 2
	// 			}
	// 			if _, found := trapped[rightZ]; bound(rightZ) && found {
	// 				air -= 2
	// 			}
	// 		}
	// 	}
	// }

	fmt.Printf("Part1 = %d\n", faces)
	// fmt.Printf("Part2 = %d\n", air)
	// fmt.Printf("Part2 = %d\n", faces-air)
}

func (c Coordinate) IsBounded(minX, maxX, minY, maxY, minZ, maxZ int) bool {
	if c.x == minX || c.x == maxX || c.y == minY || c.y == maxY || c.z == minZ || c.z == maxZ {
		return false
	}
	return true
}

func ParseLine18(line string) Coordinate {
	split := strings.Split(line, ",")

	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])
	z, _ := strconv.Atoi(split[2])

	return Coordinate{x, y, z}
}
