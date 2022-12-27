package aoc

import (
	"data"
	"fmt"
	"log"
)

const Rock = 1
const Paper = 2
const Scissor = 3

const Loss = 0
const Draw = 3
const Win = 6

func Day2() {
	lines, err := data.Load(2)
	if err != nil {
		log.Fatal("WTF")
	}

	strat1, strat2 := 0, 0
	for _, round := range lines {
		strat1 += Strategy1(string(round[0]), string(round[2]))
		strat2 += Strategy2(string(round[0]), string(round[2]))
	}
	fmt.Printf("Part 1 score = %d\n", strat1)
	fmt.Printf("Part 2 score = %d\n", strat2)
}

func Strategy1(other, me string) (score int) {
	if other == "A" && me == "X" {
		// rock vs rock
		score += Draw + Rock
	} else if other == "A" && me == "Y" {
		// rock vs paper
		score += Win + Paper
	} else if other == "A" && me == "Z" {
		// rock vs scissor
		score += Loss + Scissor
	} else if other == "B" && me == "X" {
		// paper vs rock
		score += Loss + Rock
	} else if other == "B" && me == "Y" {
		// paper vs paper
		score += Draw + Paper
	} else if other == "B" && me == "Z" {
		// paper vs scissor
		score += Win + Scissor
	} else if other == "C" && me == "X" {
		// scissor vs rock
		score += Win + Rock
	} else if other == "C" && me == "Y" {
		// scissor vs paper
		score += Loss + Paper
	} else if other == "C" && me == "Z" {
		// scissor vs scissor
		score += Draw + Scissor
	}
	return
}

func Strategy2(other, me string) (score int) {
	if other == "A" && me == "X" {
		// rock vs Lose
		score += Loss + Scissor
	} else if other == "A" && me == "Y" {
		// rock vs Draw
		score += Draw + Rock
	} else if other == "A" && me == "Z" {
		// rock vs Win
		score += Win + Paper
	} else if other == "B" && me == "X" {
		// paper vs Lose
		score += Loss + Rock
	} else if other == "B" && me == "Y" {
		// paper vs Draw
		score += Draw + Paper
	} else if other == "B" && me == "Z" {
		// paper vs Win
		score += Win + Scissor
	} else if other == "C" && me == "X" {
		// scissor vs Lose
		score += Loss + Paper
	} else if other == "C" && me == "Y" {
		// scissor vs Draw
		score += Draw + Scissor
	} else if other == "C" && me == "Z" {
		// scissor vs Win
		score += Win + Rock
	}
	return
}
