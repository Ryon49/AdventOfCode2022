package aoc

import (
	"data"
	"fmt"
	"strconv"
	"strings"
)

func Day10() {
	lines, _ := data.Load(10)

	const PrefixLength = len("addx ")

	var (
		cycle   int  = 1
		x       int  = 1
		addWait bool = false
	)

	// PART 1
	sum := 0
	for lineN := 0; lineN < len(lines); {
		// fmt.Printf("cycle %d, \"%s\", x = %d\n", cycle, lines[lineN], x)
		if (cycle-20)%40 == 0 {
			sum += x * cycle
		}
		cycle += 1
		if strings.HasPrefix(lines[lineN], "addx") {
			if !addWait {
				addWait = true
				continue
			} else {
				number, _ := strconv.Atoi(lines[lineN][PrefixLength:])
				x += number
				addWait = false
			}
		}
		lineN += 1
	}
	fmt.Printf("sum of strength = %d\n\n", sum)

	// PART 2
	cycle = 1
	x = 1
	addWait = false
	crt := ""
	for lineN := 0; lineN < len(lines); {
		if InSprite(cycle-1, x) {
			crt += "#"
		} else {
			crt += "."
		}
		if cycle%40 == 0 {
			fmt.Println(crt)
			crt = ""
		}

		cycle += 1
		if strings.HasPrefix(lines[lineN], "addx") {
			if !addWait {
				addWait = true
				continue
			} else {
				number, _ := strconv.Atoi(lines[lineN][PrefixLength:])
				x += number
				addWait = false
			}
		}
		lineN += 1
	}
}

func InSprite(cycle, x int) bool {
	position := cycle % 40

	diff := position - x
	if diff >= -1 && diff <= 1 {

		return true
	}
	return false
}
