package aoc

import (
	"data"
	"fmt"
	"regexp"
	"strconv"
)

var ActionPattern, _ = regexp.Compile("move (\\d+) from (\\d) to (\\d)")

func Day5() {
	lines, _ := data.Load(5)

	var value string
	var values []string
	stack1 := BuildStack(lines[:9])
	stack2 := BuildStack(lines[:8])
	for _, line := range lines[10:] {
		n_crate, from, to := ParseActions(line)

		// problem 1
		for n := n_crate; n > 0; n -= 1 {
			stack1[from], value = Pop(stack1[from])
			stack1[to] = Push(stack1[to], value)
		}
		// problem 2
		stack2[from], values = PopM(stack2[from], n_crate)
		stack2[to] = PushM(stack2[to], values)
	}

	for col := 1; col < 10; col += 1 {
		if len(stack1[col]) > 0 {
			fmt.Printf("%s", stack1[col][len(stack1[col])-1])
		}
	}
	fmt.Println()
	for col := 1; col < 10; col += 1 {
		if len(stack2[col]) > 0 {
			fmt.Printf("%s", stack2[col][len(stack2[col])-1])
		}
	}
	fmt.Println()
}

func BuildStack(rows []string) [][]string {
	// pad by 1
	stack := make([][]string, 10)

	indexes := []int{1, 5, 9, 13, 17, 21, 25, 29, 33}
	for rowId := len(rows) - 1; rowId >= 0; rowId -= 1 {
		row := rows[rowId]

		for i, index := range indexes {
			if row[index] != 32 {
				stack[i+1] = append(stack[i+1], string(row[index]))
			}
		}
	}
	return stack
}

// simulate Push of single crate in Queue
func Push(q []string, value string) []string {
	return append(q, value)
}

// simulate Pop of single crate in Queue
func Pop(q []string) ([]string, string) {
	last := len(q) - 1
	return q[:last], q[last]
}

// simulate Push of multiple crates in Queue
func PushM(q []string, values []string) []string {
	return append(q, values...)
}

// simulate Pop of multiple crates in Queue
func PopM(q []string, n int) ([]string, []string) {
	last := len(q) - n
	return q[:last], q[last:]
}

func ParseActions(line string) (int, int, int) {
	res := ActionPattern.FindStringSubmatch(line)
	n, _ := strconv.Atoi(res[1])
	from, _ := strconv.Atoi(res[2])
	to, _ := strconv.Atoi(res[3])
	return n, from, to
}
