package aoc

import (
	"data"
	"fmt"
	"strings"
)

func Day3() {
	lines, _ := data.Load(3)

	priority1, priority2 := 0, 0
	for _, lines := range lines {
		middle := len(lines) / 2
		first, second := lines[:middle], lines[middle:]

		common := FindCommon(first, second)
		priority1 += GetPriority(common[0])
	}

	for i := 0; i < len(lines); i = i + 3 {
		l1, l2, l3 := lines[i], lines[i+1], lines[i+2]
		common1 := FindCommon(l1, l2)
		common2 := FindCommon(l1, l3)
		intersect := FindIntersection(common1, common2)
		priority2 += GetPriority(intersect)
	}

	fmt.Printf("Priority1 = %d\n", priority1)
	fmt.Printf("Priority2 = %d\n", priority2)
}

func GetPriority(common int) int {
	if int('A') <= common && common <= int('Z') {
		return common - int('A') + 1 + 26
	} else {
		return common - int('a') + 1
	}
}

// return a set of number that appears both in s1 and s2
func FindCommon(s1, s2 string) []int {
	var set = []int{}

	for _, c := range s1 {
		add := true
		if strings.Contains(s2, string(c)) {
			for _, v := range set {
				if v == int(c) {
					add = false
				}
			}
			if add {
				set = append(set, int(c))
			}
		}
	}
	return set
}

// n^2 to simulate finding intersection of 2 sets
func FindIntersection(a1, a2 []int) int {
	for _, v := range a1 {
		for _, u := range a2 {
			if v == u {
				return v
			}
		}
	}
	return -1
}
