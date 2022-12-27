package aoc

import (
	"data"
	"fmt"
)

func Day6() {
	lines, _ := data.Load(6)

	input := lines[0]
	for i := 4; i < len(input); i += 1 {
		head := string(input[i-4 : i])
		if Valid(head) {
			fmt.Println(i)
			break
		}
	}
	for i := 14; i < len(input); i += 1 {
		head := string(input[i-14 : i])
		if Valid(head) {
			fmt.Println(i)
			break
		}
	}
}

func Valid(s string) bool {
	var set [26]int
	for _, c := range s {
		at := c - 'a'
		if set[at] > 0 {
			return false
		}
		set[at] += 1
	}
	return true
}
