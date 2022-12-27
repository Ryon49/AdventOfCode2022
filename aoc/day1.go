package aoc

import (
	"data"
	"fmt"
	"log"
	"sort"
	"strconv"
)

func Day1() {
	lines, err := data.Load(1)
	if err != nil {
		log.Fatal("WTF?")
	}

	var weight int
	var weights = []int{}
	for i := 0; i < len(lines); {
		weight, i = NextWeight(lines, i)
		weights = append(weights, weight)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(weights)))

	fmt.Printf("Top 1 = %d\n", weights[0])
	fmt.Printf("Top 3 sum = %d\n", weights[0]+weights[1]+weights[2])
}

func NextWeight(lines []string, i int) (int, int) {
	weight := 0
	for ; i < len(lines); i++ {
		if lines[i] == "" {
			break
		}
		value, _ := strconv.Atoi(lines[i])
		weight += value
	}
	return weight, i + 1
}
