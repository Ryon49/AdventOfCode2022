package aoc

import (
	"data"
	"fmt"
	"regexp"
	"strconv"
)

type Blueprint struct {
	id       int
	ore      int
	clay     int
	obsidian []int
	geode    []int
}

type TimeNode struct {
	blueprint Blueprint
	time      int

	ore, clay, obsidian, geode             int
	oreBot, clayBot, obsidianBot, geodeBot int

	next []TimeNode
}

func (node *TimeNode) DFS() int {
	// checl if has enough resource to build Ore Bot

	bp := node.blueprint
	if node.ore >= node.blueprint.ore {
		t := TimeNode{blueprint: bp, time: node.time + 1,
			ore: node.ore - bp.ore, clay: node.clay, obsidian: node.obsidian, geode: node.geode,
			oreBot: node.oreBot + 1, clayBot: node.clayBot, obsidianBot: node.obsidianBot, geodeBot: node.geodeBot}
		t.DFS()
	}

	return 0
}

func Day19() {
	lines, _ := data.Load(19)

	for _, line := range lines[:1] {
		blueprint := ParseLine19(line)
		fmt.Println(blueprint)
		// root := TimeNode{blueprint: bleuprint, time: 0,
		// 	ore: 0, clay: 0, obsidian: 0, geode: 0,
		// 	oreBot: 1, clayBot: 0, obsidianBot: 0, geodeBot: 0}
	}
}

func ParseLine19(line string) Blueprint {
	pattern := regexp.MustCompile("(\\d+)")
	matches := pattern.FindAllString(line, -1)

	id, _ := strconv.Atoi(matches[0])

	oreBot, _ := strconv.Atoi(matches[1])
	clayBot, _ := strconv.Atoi(matches[2])

	obsidianOre, _ := strconv.Atoi(matches[3])
	obsidianclay, _ := strconv.Atoi(matches[4])

	geodeOre, _ := strconv.Atoi(matches[5])
	geodeObsidian, _ := strconv.Atoi(matches[6])

	return Blueprint{
		id:       id,
		ore:      oreBot,
		clay:     clayBot,
		obsidian: []int{obsidianOre, obsidianclay},
		geode:    []int{geodeOre, geodeObsidian}}
}
