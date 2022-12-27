package aoc

// TODO

// type Valve struct {
// 	name     string
// 	flowRate int
// 	edges    []string

// 	distance map[string]int
// }

// func Day16() {
// 	lines, _ := data.Load(16)

// 	// load graph
// 	G := make(map[string]Valve)
// 	for _, line := range lines {
// 		valve := Parse16(line)
// 		G[valve.name] = valve
// 	}
// 	// remove valves that has a flow rate == 0
// 	for _, valve := range G {
// 		if valve.name != "AA" && valve.flowRate == 0 {
// 			delete(G, valve.name)
// 		}
// 	}
// 	// for each remaining valve, calculate the distance between this to others
// 	for _, valveA := range G {
// 		for _, valveB := range G {
// 			if _, found := valveA.distance[valveB.name]; !found && valveA.name != valveB.name {
// 				distance := Distance16(G, valveA, valveB, make(map[string]bool))
// 				valveA.distance[valveB.name] = distance
// 				valveB.distance[valveA.name] = distance
// 			}
// 		}
// 	}

// 	for k, v := range G {
// 		fmt.Println(k, v.distance)
// 	}

// }

// func Distance16(G map[string]Valve, a, b Valve, visited map[string]bool) int {
// 	if a.name == b.name {
// 		return 0
// 	}
// 	for _, edge := range a.edges {
// 		if edge == b.name {
// 			return 1
// 		}
// 	}

// 	distance := 10000
// 	visited[b.name] = true
// 	for _, edge := range a.edges {
// 		if _, found := visited[b.name]; found {
// 			continue
// 		}
// 		d := Distance16(G, G[edge], b, visited)
// 		if d < distance {
// 			distance = d
// 		}
// 	}
// 	return distance
// }

// func Parse16(line string) (v Valve) {
// 	pattern := regexp.MustCompile("Valve (\\w+) has flow rate=(\\d+); tunnels? leads? to valves? (.+)")
// 	matches := pattern.FindStringSubmatch(line)

// 	v.name = matches[1]
// 	flowRate, _ := strconv.Atoi(matches[2])
// 	v.flowRate = flowRate
// 	v.edges = strings.Split(matches[3], ", ")

// 	v.distance = make(map[string]int)
// 	return
// }
