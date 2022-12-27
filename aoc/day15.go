package aoc

import (
	"data"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Detection struct {
	sensor, beacon Position
}

func Day15() {
	lines, _ := data.Load(15)

	detections := []Detection{}
	// we use "reserved" to prevent counting on beacon or sensor
	reserved := make(map[Position]bool)

	minX, maxX, minY, maxY := math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	for _, line := range lines {
		sensor, beacon := ParseLine15(line)
		detections = append(detections, Detection{sensor, beacon})
		reserved[sensor] = true
		reserved[beacon] = true

		distance := ManhattanDistance(sensor, beacon)
		minX = int(math.Min(float64(minX), float64(sensor.x-distance)))
		maxX = int(math.Max(float64(maxX), float64(sensor.x+distance)))

		minY = int(math.Min(float64(minY), float64(sensor.y-distance)))
		maxY = int(math.Max(float64(maxY), float64(sensor.y+distance)))
	}

	// fmt.Println(minX, maxX, minY, maxY)
	// targetRow := 2000000
	// marked := 0
	// for i := minY; i <= maxY; i += 1 {
	// 	origin := Position{targetRow, i}
	// 	if _, found := reserved[origin]; found {
	// 		continue
	// 	}

	// 	for _, detection := range detections {
	// 		if ManhattanDistance(origin, detection.sensor) <= ManhattanDistance(detection.beacon, detection.sensor) {
	// 			marked += 1
	// 			break
	// 		}
	// 	}
	// }
	// fmt.Printf("marked = %d\n", marked)

	grid := make([][]byte, 21)
	for i := 0; i <= 20; i += 1 {
		for j := 0; j <= 20; j += 1 {
			grid[i] = append(grid[i], '.')
		}
	}

	maxRange := 4000000
	beacon := LocateBeacon(detections, maxRange)
	fmt.Printf("tuning frequency = %d\n", beacon.y*maxRange+beacon.x)
}

func LocateBeacon(detections []Detection, maxRange int) Position {
	var origin Position
	intersection := make(map[Position]bool)
	for _, detection := range detections {
		distance := ManhattanDistance(detection.sensor, detection.beacon)

		minX, maxX := detection.sensor.x-distance-1, detection.sensor.x+distance+1
		minY, maxY := detection.sensor.y-distance-1, detection.sensor.y+distance+1

		// push the edge into intersections
		top, bottom := Position{minX, detection.sensor.y}, Position{maxX, detection.sensor.y}
		left, right := Position{detection.sensor.x, minY}, Position{detection.sensor.x, maxY}

		origin = top
		for ManhattanDistance(origin, detection.sensor) == distance+1 {
			// 	   top
			//    -
			// left
			if origin.x >= 0 && origin.x <= maxRange && origin.y >= 0 && origin.y <= maxRange {
				// bounding
				intersection[origin] = true
			}
			origin.x += 1
			origin.y -= 1
		}

		origin = Position{top.x + 1, top.y + 1}
		for ManhattanDistance(origin, detection.sensor) == distance+1 {
			// 	   top
			//    	  -
			// 			right
			if origin.x >= 0 && origin.x <= maxRange && origin.y >= 0 && origin.y <= maxRange {
				// bounding
				intersection[origin] = true
			}
			origin.x += 1
			origin.y += 1
		}

		origin = bottom
		for ManhattanDistance(origin, detection.sensor) == distance+1 {
			if origin.x >= 0 && origin.x <= maxRange && origin.y >= 0 && origin.y <= maxRange {
				// bounding
				intersection[origin] = true
			}
			origin.x -= 1
			origin.y -= 1
			if origin == left {
				break
			}
		}

		origin = Position{bottom.x - 1, bottom.y + 1}
		for ManhattanDistance(origin, detection.sensor) == distance+1 {
			if origin.x >= 0 && origin.x <= maxRange && origin.y >= 0 && origin.y <= maxRange {
				// bounding
				intersection[origin] = true
			}
			origin.x -= 1
			origin.y += 1
			if origin == right {
				break
			}
		}
	}

	for origin, _ := range intersection {
		found := true
		for _, detection := range detections {
			if ManhattanDistance(origin, detection.sensor) <= ManhattanDistance(detection.sensor, detection.beacon) {
				found = false
			}
		}
		if found {
			return origin
		}
	}
	return Position{0, 0}
}

func ManhattanDistance(origin, target Position) int {
	diffX := origin.x - target.x
	diffY := origin.y - target.y
	if diffX < 0 {
		diffX = -diffX
	}
	if diffY < 0 {
		diffY = -diffY
	}
	return diffX + diffY
}

func Projected(x, minX int) int {
	return x - minX
}

func ParseLine15(line string) (Position, Position) {
	pattern := regexp.MustCompile("(-?\\d+)")
	matches := pattern.FindAllString(line, -1)

	sx, _ := strconv.Atoi(matches[0])
	sy, _ := strconv.Atoi(matches[1])

	bx, _ := strconv.Atoi(matches[2])
	by, _ := strconv.Atoi(matches[3])

	return Position{sy, sx}, Position{by, bx}
}
