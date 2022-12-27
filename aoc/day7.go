package aoc

import (
	"data"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var (
	DirPattern, _  = regexp.Compile("dir (\\w+)")
	FilePattern, _ = regexp.Compile("(\\d+) (.+)")
)

type File struct {
	name string
	size int
}

type Node struct {
	name string
	size int

	parent      *Node
	files       []File
	directories []*Node
}

func Day7() {
	lines, _ := data.Load(7)

	root := BuildTree(lines)
	sum := root.DFS()

	// part 1
	sumLessThan := root.FindSum(100000)
	fmt.Printf("total size < 100000 = %d\n", sumLessThan)

	// part 2
	spaceNeeded := 30000000 - (70000000 - sum) // space needed - avaiable space
	cloest := root.FindClosest(spaceNeeded)
	fmt.Printf("smallest directory to delete has size = %d\n", cloest)
}

func BuildTree(lines []string) *Node {
	root := new(Node)
	root.name = "/"

	var cur *Node = root

	for i := 0; i < len(lines); {
		_, commands := ParseCommand(lines[i])

		if commands[0] == "cd" {
			if commands[1] == "/" {
				cur = root
			} else if commands[1] == ".." {
				cur = cur.parent
			} else {
				for _, directory := range cur.directories {
					if directory.name == commands[1] {
						cur = directory
					}
				}
			}
			i += 1
		} else if commands[0] == "ls" {
			i = cur.LS(lines, i+1)
		}

	}
	return root
}

func (node *Node) LS(lines []string, i int) int {
	for ; i < len(lines); i += 1 {
		if isDir, dirname := ParseDir(lines[i]); isDir {
			// create a new Node pointer and add to current Node
			newNode := new(Node)
			newNode.name = dirname
			newNode.parent = node
			node.directories = append(node.directories, newNode)

		} else if isFile, size, filename := ParseFile(lines[i]); isFile {
			newFile := File{name: filename, size: size}
			node.files = append(node.files, newFile)
		} else {
			// reaches next command
			break
		}
	}
	return i
}

func (node *Node) DFS() int {
	if node.size > 0 {
		return node.size
	}
	totalSize := 0
	for _, f := range node.files {
		totalSize += f.size
	}
	for _, dir := range node.directories {
		totalSize += dir.DFS()
	}
	node.size = totalSize
	return totalSize
}

func (node *Node) FindSum(atMost int) int {
	sum := 0
	if node.size <= atMost {
		sum += node.size
	}
	for _, dir := range node.directories {
		sum += dir.FindSum(atMost)
	}
	return sum
}

// return a minimum directory size that > minimum
func (node *Node) FindClosest(minimum int) int {
	if node.size < minimum {
		return math.MaxInt32
	}
	cloest := node.size
	for _, dir := range node.directories {
		cur := dir.FindClosest(minimum)
		if (cloest - minimum) > (cur - minimum) {
			cloest = cur
		}
	}
	return cloest
}

func ParseCommand(line string) (bool, []string) {
	if strings.HasPrefix(line, "$") {
		return true, strings.Split(line[2:], " ")
	}
	return false, nil
}

func ParseDir(line string) (bool, string) {
	if DirPattern.MatchString(line) {
		matches := DirPattern.FindStringSubmatch(line)
		return true, matches[1]
	}
	return false, ""
}
func ParseFile(line string) (bool, int, string) {
	if FilePattern.MatchString(line) {
		matches := FilePattern.FindStringSubmatch(line)
		size, _ := strconv.Atoi(matches[1])
		return true, size, matches[2]
	}
	return false, 0, ""
}
