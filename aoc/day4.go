package aoc

import (
	"data"
	"fmt"
	"strconv"
	"strings"
)

func Day4() {
	lines, _ := data.Load(4)

	count1, count2 := 0, 0
	for _, pair := range lines {
		p1, p2 := CommaSplit(pair)

		p1L, p1R := IntSplit(p1)
		p2L, p2R := IntSplit(p2)
		if p1L <= p2L && p1R >= p2R {
			count1 += 1
		} else if p2L <= p1L && p2R >= p1R {
			count1 += 1
		}

		if p1L <= p2L && p2L <= p1R {
			count2 += 1
		} else if p2L <= p1L && p1L <= p2R {
			count2 += 1
		}
	}

	fmt.Printf("count1 = %d\n", count1)
	fmt.Printf("count2 = %d\n", count2)
}

func CommaSplit(s string) (string, string) {
	split := strings.Split(s, ",")
	return split[0], split[1]
}

func IntSplit(s string) (int, int) {
	split := strings.Split(s, "-")
	v1, _ := strconv.Atoi(split[0])
	v2, _ := strconv.Atoi(split[1])
	return v1, v2
}
