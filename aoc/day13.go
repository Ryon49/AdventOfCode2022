package aoc

import (
	"data"
	"fmt"
	"sort"
)

type IntList struct {
	value   int
	lists   []IntList
	isValue bool
}

func Day13() {
	lines, _ := data.Load(13)

	// part 1
	sumOfRightOrder := 0
	for i, pairIndex := 0, 1; i < len(lines); i, pairIndex = i+3, pairIndex+1 {
		p0 := ParseLine(lines[i])
		p1 := ParseLine(lines[i+1])
		lessThan, _ := CompareIntList(p0, p1)

		if lessThan {
			sumOfRightOrder += pairIndex
		}
	}
	fmt.Printf("Sum of right orders = %d\n", sumOfRightOrder)

	// part 2

	array := []IntList{}
	for i, pairIndex := 0, 1; i < len(lines); i, pairIndex = i+3, pairIndex+1 {
		p0 := ParseLine(lines[i])
		array = append(array, p0)

		p1 := ParseLine(lines[i+1])
		array = append(array, p1)
	}

	decoder2 := IntList{value: 2, isValue: true}
	decoder6 := IntList{value: 6, isValue: true}
	array = append(array, IntList{isValue: false, lists: []IntList{decoder2}})
	array = append(array, IntList{isValue: false, lists: []IntList{decoder6}})

	sort.Slice(array, func(i, j int) bool {
		lessThan, _ := CompareIntList(array[i], array[j])
		return lessThan
	})

	decoder := 1
	for i, intList := range array {
		if _, equal := CompareIntList(decoder2, intList); equal {
			fmt.Printf("decoder 2 = %d\n", i)
			decoder *= (i + 1)
		}
		if _, equal := CompareIntList(decoder6, intList); equal {
			fmt.Printf("decoder 6 = %d\n", i)
			decoder *= (i + 1)
		}
	}
	fmt.Printf("decoder key = %d\n", decoder)
}

func ParseLine(line string) IntList {
	stack := make([]IntList, 0)

	value, i := 0, 0
	for i < len(line) {
		if line[i] == '[' {
			stack = append([]IntList{{}}, stack...)
			i += 1
		} else if line[i] == ']' {
			if i != len(line)-1 {
				pop := stack[0]
				stack = stack[1:]
				stack[0].lists = append(stack[0].lists, pop)
			}
			i += 1
		} else if IsDigit(line[i]) {
			value, i = ParseInt(line, i)
			stack[0].lists = append(stack[0].lists, IntList{value: value, isValue: true})
		} else {
			i += 1
		}
	}

	return stack[0]
}

func ParseInt(line string, i int) (int, int) {
	ret := 0
	for IsDigit(line[i]) {
		d := int(line[i] - '0')
		ret = ret*10 + d
		i += 1
	}
	return ret, i
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func CompareIntList(int1, int2 IntList) (bool, bool) {
	if int1.isValue && int2.isValue {
		return int1.value < int2.value, int1.value == int2.value
	}
	if int1.isValue && !int2.isValue {
		wrap := IntList{isValue: false, lists: []IntList{int1}}
		return CompareIntList(wrap, int2)
	} else if !int1.isValue && int2.isValue {
		wrap := IntList{isValue: false, lists: []IntList{int2}}
		return CompareIntList(int1, wrap)
	}

	if len(int1.lists) > len(int2.lists) {
		greaterEqual, notEqual := CompareIntList(int2, int1)
		return !greaterEqual, notEqual
	}

	for i := 0; i < len(int1.lists); i += 1 {
		lessThan, equal := CompareIntList(int1.lists[i], int2.lists[i])
		if lessThan {
			return true, false
		} else if !equal {
			return false, false
		}
	}

	// if reaches here, both IntList are identiy up to int1, so if int2 still have elements left
	// int1 is less than int2
	if len(int1.lists) < len(int2.lists) {
		return true, false
	}
	return false, true
}
