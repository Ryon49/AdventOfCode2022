package aoc

import (
	"data"
	"fmt"
	"strconv"
)

type NumberWrapper struct {
	n     int
	order int
}

func Day20() {
	lines, _ := data.Load(20)

	sequence := []NumberWrapper{}
	for order, line := range lines {
		n, _ := strconv.Atoi(line)
		sequence = append(sequence, NumberWrapper{n, order})
	}

	length := len(sequence)
	fmt.Println(length)

	for order := 0; order < length; order += 1 {
		number, index := FindOrder(sequence, order)

		tmp := append(sequence[:index], sequence[index+1:]...)

		endIndex := index + number.n
		if endIndex <= 0 {
			endIndex = GoModulo(endIndex-1, length)
		} else if endIndex >= length {
			endIndex = GoModulo(endIndex+1, length)
		}

		sequence = append(tmp[:endIndex+1], tmp[endIndex:]...)
		sequence[endIndex] = number
	}

	// fmt.Println(sequence)
	zero := FindNum(sequence, 0)

	n1000 := (zero + 1000) % length
	n2000 := (zero + 2000) % length
	n3000 := (zero + 3000) % length

	fmt.Println(sequence[n1000].n, sequence[n2000].n, sequence[n3000].n)
	fmt.Println(sequence[n1000].n + sequence[n2000].n + sequence[n3000].n)
}

func GoModulo(d, m int) int {
	// https://stackoverflow.com/questions/43018206/modulo-of-negative-integers-in-go
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

func FindOrder(sequence []NumberWrapper, order int) (NumberWrapper, int) {
	i := 0
	for {
		if sequence[i].order == order {
			return sequence[i], i
		}
		i += 1
	}
}

func FindNum(sequence []NumberWrapper, n int) int {
	i := 0
	for {
		if sequence[i].n == n {
			return i
		}
		i += 1
	}
}
