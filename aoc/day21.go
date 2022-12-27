package aoc

import (
	"data"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Operation21 struct {
	ready                  bool
	name                   string
	value                  int
	op, operand1, operand2 string

	str string
}

type Equal struct {
	value    int
	variable string
}

func Day21() {

	lines, _ := data.Load(21)

	fmt.Println(Day21Part1(lines, make(map[string]int), "root"))
	fmt.Println(Day21Part2(lines))
}

func Day21Part2(lines []string) int {
	values := make(map[string]int)
	pending := make(map[string]Operation21)
	for _, line := range lines {
		if strings.HasPrefix(line, "humn") {
			continue
		}
		operation := ParseLine21(line)
		if operation.ready {
			values[operation.name] = operation.value
		} else {
			pending[operation.name] = operation
		}
	}

	for len(pending) > 0 {
		size := len(pending)
		for name, operation := range pending {
			v1, f1 := values[operation.operand1]
			v2, f2 := values[operation.operand2]
			if f1 && f2 {
				if operation.op == "+" {
					values[name] = v1 + v2
				} else if operation.op == "*" {
					values[name] = v1 * v2
				} else if operation.op == "-" {
					values[name] = v1 - v2
				} else if operation.op == "/" {
					values[name] = v1 / v2
				}
				delete(pending, name)
			}
		}
		if size == len(pending) {
			break
		}
	}

	// revert the operations
	newLines := []string{}
	for _, operation := range pending {
		// fmt.Println(operation.str)
		knownV, value, unknownV := DecomposeOperation(values, operation)

		if operation.name == "root" {
			newLines = append(newLines, fmt.Sprintf("%s: %d", unknownV, value))
		} else {
			var line string
			if operation.op == "+" {
				// a = b + c
				// b = a - c
				// c = a - b
				line = fmt.Sprintf("%s: %s - %s", unknownV, operation.name, knownV)
			} else if operation.op == "*" {
				// a = b * c
				// b = a / c
				// c = a / b
				line = fmt.Sprintf("%s: %s / %s", unknownV, operation.name, knownV)
			} else if operation.op == "-" {
				// a = b - c
				// b = a + c
				// c = b - a
				if unknownV == operation.operand1 {
					line = fmt.Sprintf("%s: %s + %s", unknownV, operation.name, knownV)
				} else {
					line = fmt.Sprintf("%s: %s - %s", unknownV, knownV, operation.name)
				}
			} else if operation.op == "/" {
				// a = b / c
				// b = a * c
				// c = b / a
				if unknownV == operation.operand1 {
					line = fmt.Sprintf("%s: %s * %s", unknownV, operation.name, knownV)
				} else {
					line = fmt.Sprintf("%s: %s / %s", unknownV, knownV, operation.name)
				}
			}
			// fmt.Println(line)
			newLines = append(newLines, line)
		}
	}

	return Day21Part1(newLines, values, "humn")
}

func Day21Part1(lines []string, values map[string]int, target string) int {
	pending := make(map[string]Operation21)
	for _, line := range lines {
		operation := ParseLine21(line)
		if operation.ready {
			values[operation.name] = operation.value
		} else {
			pending[operation.name] = operation
		}
	}

	for len(pending) > 0 {
		for name, operation := range pending {
			v1, f1 := values[operation.operand1]
			v2, f2 := values[operation.operand2]
			if f1 && f2 {
				if operation.op == "+" {
					values[name] = v1 + v2
				} else if operation.op == "*" {
					values[name] = v1 * v2
				} else if operation.op == "-" {
					values[name] = v1 - v2
				} else if operation.op == "/" {
					values[name] = v1 / v2
				}
				delete(pending, name)
			}
		}
	}
	return values[target]
}

// return name, known value from values, and unknown variable
func DecomposeOperation(values map[string]int, op Operation21) (string, int, string) {
	if v, found := values[op.operand1]; found {
		return op.operand1, v, op.operand2
	}
	return op.operand2, values[op.operand2], op.operand1
}

func ParseLine21(line string) Operation21 {
	ValuePattern := regexp.MustCompile("(\\w{4}): (\\d+)")
	EqPattern := regexp.MustCompile("(\\w{4}): (\\w{4}) (.) (\\w{4})")

	if ValuePattern.MatchString(line) {
		matches := ValuePattern.FindStringSubmatch(line)

		value, _ := strconv.Atoi(matches[2])
		return Operation21{ready: true, name: matches[1], value: value, str: line}
	} else {
		matches := EqPattern.FindStringSubmatch(line)
		return Operation21{ready: false, name: matches[1], op: matches[3],
			operand1: matches[2], operand2: matches[4], str: line}
	}
}
