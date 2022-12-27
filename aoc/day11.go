package aoc

import (
	"fmt"
	"sort"
)

func Day11() {
	nextMonkey := []func(int) int{
		func(worry int) int {
			if worry%3 == 0 {
				return 2
			}
			return 1
		},
		func(worry int) int {
			if worry%13 == 0 {
				return 7
			}
			return 2
		},
		func(worry int) int {
			if worry%19 == 0 {
				return 4
			}
			return 7
		},
		func(worry int) int {
			if worry%17 == 0 {
				return 6
			}
			return 5
		},
		func(worry int) int {
			if worry%5 == 0 {
				return 6
			}
			return 3
		},
		func(worry int) int {
			if worry%7 == 0 {
				return 1
			}
			return 0
		},
		func(worry int) int {
			if worry%11 == 0 {
				return 5
			}
			return 0
		},
		func(worry int) int {
			if worry%2 == 0 {
				return 4
			}
			return 3
		},
	}
	allItems := [][]int{
		{54, 98, 50, 94, 69, 62, 53, 85},
		{71, 55, 82},
		{77, 73, 86, 72, 87},
		{97, 91},
		{78, 97, 51, 85, 66, 63, 62},
		{88},
		{87, 57, 63, 86, 87, 53},
		{73, 59, 82, 65},
	}
	ops := GenerateOps(3)
	inspection := []int{0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < 20; i += 1 {
		for monkey, _ := range allItems {
			for len(allItems[monkey]) > 0 {
				inspection[monkey] += 1

				item := allItems[monkey][0]
				allItems[monkey] = allItems[monkey][1:]

				worry := ops[monkey](item)
				next := nextMonkey[monkey](worry)
				allItems[next] = append(allItems[next], worry)
			}
		}
	}
	sort.Ints(inspection)
	fmt.Printf("Part 1 = %d\n", inspection[6]*inspection[7])

	allItems = [][]int{
		{54, 98, 50, 94, 69, 62, 53, 85},
		{71, 55, 82},
		{77, 73, 86, 72, 87},
		{97, 91},
		{78, 97, 51, 85, 66, 63, 62},
		{88},
		{87, 57, 63, 86, 87, 53},
		{73, 59, 82, 65},
	}
	ops = GenerateOps(1)
	inspection = []int{0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < 10000; i += 1 {
		for monkey, _ := range allItems {
			for len(allItems[monkey]) > 0 {
				inspection[monkey] += 1

				item := allItems[monkey][0]
				allItems[monkey] = allItems[monkey][1:]

				worry := ops[monkey](item)
				next := nextMonkey[monkey](worry)
				allItems[next] = append(allItems[next], worry)
			}
		}
	}
	sort.Ints(inspection)
	fmt.Printf("Part 2 = %d\n", inspection[6]*inspection[7])
}

func GenerateOps(relief int) []func(int) int {
	mod := 3 * 13 * 19 * 17 * 5 * 7 * 11 * 2
	return []func(int) int{
		func(old int) int { return ((old * 13) / relief) % mod },
		func(old int) int { return ((old + 2) / relief) % mod },
		func(old int) int { return ((old + 8) / relief) % mod },
		func(old int) int { return ((old + 1) / relief) % mod },
		func(old int) int { return ((old * 17) / relief) % mod },
		func(old int) int { return ((old + 3) / relief) % mod },
		func(old int) int { return ((old * old) / relief) % mod },
		func(old int) int { return ((old + 6) / relief) % mod },
	}
}

// 0 -> 2 -> 1 -> 2
// 0 -> 2 -> 1 -> 0
// 0 -> 2 -> 3 -> 0
// 0 -> 2 -> 3 -> 1 -> 2
// 0 -> 2 -> 3 -> 1 -> 0

// 1 -> 2 -> 1
// 1 -> 2 -> 3 -> 0
// 1 -> 2 -> 3 -> 1
