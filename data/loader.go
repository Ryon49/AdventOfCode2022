package data

import (
	"bufio"
	"fmt"
	"os"
)

func Load(day int) (lines []string, err error) {
	path := fmt.Sprintf("./data/day%d.txt", day)
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}
