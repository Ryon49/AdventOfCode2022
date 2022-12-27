package aoc

// type ShapeFunc interface {
// 	MoveLeft(grid [][]byte)
// 	MoveRight(grid [][]byte)
// 	MoveBottom(grid [][]byte) bool
// 	FormRock(grid [][]byte)
// 	Top() int
// }

// type Shape struct {
// 	// this indicates the left and bottom position of shape in start
// 	left, bottom Position
// }

// type HorizonalBar Shape
// type Cross Shape
// type ReflectedL Shape
// type VerticalBar Shape
// type Square Shape

// const RightEdge int = 6

// func Day17Old() {
// 	lines, _ := data.Load(17)

// 	shapes := []func(int) ShapeFunc{
// 		func(height int) ShapeFunc {
// 			// horizontal bar line, left == bottom
// 			return &HorizonalBar{left: Position{height + 3, 2}, bottom: Position{height + 3, 2}}
// 		},
// 		func(height int) ShapeFunc {
// 			// cross, left == bottom
// 			return &Cross{left: Position{height + 4, 2}, bottom: Position{height + 3, 3}}
// 		},
// 		func(height int) ShapeFunc {
// 			// reflected L, left == bottom
// 			return &ReflectedL{left: Position{height + 3, 2}, bottom: Position{height + 3, 2}}
// 		},
// 		func(height int) ShapeFunc {
// 			// vertical bar line, left == bottom
// 			return &VerticalBar{left: Position{height + 3, 2}, bottom: Position{height + 3, 2}}
// 		},
// 		func(height int) ShapeFunc {
// 			// horizontal bar line, left == bottom
// 			return &Square{left: Position{height + 3, 2}, bottom: Position{height + 3, 2}}
// 		},
// 	}

// 	height, rocks := 0, 0
// 	patternIndex, shapeIndex := 0, 0
// 	shape := shapes[shapeIndex](height)

// 	grid := CreateGrid17()
// 	fmt.Println(height)
// 	trace := []byte{}
// 	for rocks < 2022 {
// 		// fmt.Printf("%c\n", c)
// 		trace = append(trace, lines[0][patternIndex])
// 		switch lines[0][patternIndex] {
// 		case '<':
// 			shape.MoveLeft(grid)
// 		case '>':
// 			shape.MoveRight(grid)

// 		}
// 		if !shape.MoveBottom(grid) {
// 			shape.FormRock(grid)
// 			if shape.Top() > height {
// 				height = shape.Top()
// 			}
// 			// fmt.Println(height)

// 			shapeIndex = (shapeIndex + 1) % 5
// 			// fmt.Printf("height = %d, next = %d\n", height, shapeIndex)
// 			// fmt.Printf("trace = %s\n", trace)
// 			// PrintGridReverse(grid, height+7)
// 			// trace = []byte{}

// 			shape = shapes[shapeIndex](height)
// 			rocks += 1

// 		}
// 		patternIndex = (patternIndex + 1) % len(lines[0])
// 	}
// 	fmt.Println(height)
// }

// func CreateGrid17() (grid [][]byte) {
// 	for i := 0; i <= 2023*4; i += 1 {
// 		grid = append(grid, []byte{'.', '.', '.', '.', '.', '.', '.'})
// 	}
// 	return
// }

// func PrintGridReverse(grid [][]byte, upto int) {
// 	for i := upto; i >= 0; i -= 1 {
// 		fmt.Printf("%d\t ", i)
// 		for j := 0; j < len(grid[0]); j += 1 {
// 			fmt.Printf("%c", grid[i][j])
// 		}
// 		fmt.Println()
// 	}
// }

// // region HorizontalBar
// func (shape *HorizonalBar) MoveLeft(grid [][]byte) {
// 	if shape.left.y == 0 || grid[shape.left.x][shape.left.y-1] == '#' {
// 		return
// 	}
// 	shape.left.y -= 1
// 	shape.bottom.y -= 1

// }

// func (shape *HorizonalBar) MoveRight(grid [][]byte) {
// 	if shape.left.y+3 == 6 || grid[shape.left.x][shape.left.y+2+1] == '#' {
// 		return
// 	}
// 	shape.left.y += 1
// 	shape.bottom.y += 1

// }

// func (shape *HorizonalBar) MoveBottom(grid [][]byte) bool {
// 	for i := 0; i < 4; i += 1 {
// 		if shape.bottom.x == 0 || grid[shape.bottom.x-1][shape.bottom.y+i] == '#' {
// 			// if reached bottom or hit a rock
// 			return false
// 		}
// 	}
// 	shape.left.x -= 1
// 	shape.bottom.x -= 1

// 	return true
// }

// func (shape HorizonalBar) Top() int {
// 	return shape.left.x + 1
// }

// func (shape HorizonalBar) FormRock(grid [][]byte) {
// 	for i := 0; i < 4; i += 1 {
// 		grid[shape.left.x][shape.left.y+i] = '#'
// 	}
// }

// // endregion

// // region Cross
// func (shape *Cross) MoveLeft(grid [][]byte) {
// 	if shape.left.y == 0 || grid[shape.left.x][shape.left.y-1] == '#' {
// 		// check left moving left
// 		return
// 	}
// 	if grid[shape.bottom.x][shape.left.y-1] == '#' || grid[shape.bottom.x+2][shape.left.y-1] == '#' {
// 		// check bottom moving left
// 		return
// 	}
// 	shape.left.y -= 1
// 	shape.bottom.y -= 1
// 	// fmt.Printf("Left = %s\n", shape.bottom)
// }

// func (shape *Cross) MoveRight(grid [][]byte) {
// 	if shape.left.y+2 == 6 || grid[shape.left.x][shape.left.y+2+1] == '#' {
// 		// check left moving right
// 		return
// 	}
// 	if grid[shape.bottom.x][shape.left.y+1+1] == '#' || grid[shape.bottom.x+2][shape.left.y+1+1] == '#' {
// 		// check bottom moving right
// 		return
// 	}
// 	shape.left.y += 1
// 	shape.bottom.y += 1
// 	// fmt.Printf("Right = %s\n", shape.bottom)
// }

// func (shape *Cross) MoveBottom(grid [][]byte) bool {
// 	if shape.bottom.x == 0 || grid[shape.bottom.x-1][shape.bottom.y] == '#' {
// 		// check bottom moving down
// 		return false
// 	}
// 	if grid[shape.left.x-1][shape.left.y] == '#' || grid[shape.left.x-1][shape.left.y+2] == '#' {
// 		// check left and right moving down
// 		return false
// 	}

// 	shape.left.x -= 1
// 	shape.bottom.x -= 1
// 	// fmt.Printf("Bottom = %s\n", shape.bottom)
// 	return true
// }

// func (shape Cross) Top() int {
// 	return shape.left.x + 1 + 1
// }

// func (shape Cross) FormRock(grid [][]byte) {
// 	// top
// 	grid[shape.bottom.x+2][shape.bottom.y] = '#'
// 	// middle
// 	for i := 0; i < 3; i += 1 {
// 		grid[shape.left.x][shape.left.y+i] = '#'
// 	}
// 	// bottom
// 	grid[shape.bottom.x][shape.bottom.y] = '#'
// }

// // endregion

// // region ReflectedL
// func (shape *ReflectedL) MoveLeft(grid [][]byte) {
// 	if shape.left.y == 0 || grid[shape.left.x][shape.left.y-1] == '#' {
// 		// check moving left
// 		return
// 	}
// 	if grid[shape.left.x+1][shape.left.y+2-1] == '#' || grid[shape.left.x+2][shape.left.y+2-1] == '#' {
// 		// check moving left
// 		return
// 	}
// 	shape.left.y -= 1
// 	shape.bottom.y -= 1
// 	// fmt.Printf("Left = %s\n", shape.bottom)
// }

// func (shape *ReflectedL) MoveRight(grid [][]byte) {
// 	if shape.left.y+2 == 6 {
// 		return
// 	}
// 	for i := 0; i < 3; i += 1 {
// 		if grid[shape.left.x+i][shape.left.y+2+1] == '#' {
// 			return
// 		}
// 	}
// 	shape.left.y += 1
// 	shape.bottom.y += 1
// }

// func (shape *ReflectedL) MoveBottom(grid [][]byte) bool {
// 	for i := 0; i < 3; i += 1 {
// 		if shape.bottom.x == 0 || grid[shape.bottom.x-1][shape.bottom.y+i] == '#' {
// 			// if reached bottom or hit a rock
// 			return false
// 		}
// 	}
// 	shape.left.x -= 1
// 	shape.bottom.x -= 1

// 	// fmt.Printf("Bottom = %s\n", shape.bottom)
// 	return true
// }

// func (shape ReflectedL) Top() int {
// 	return shape.left.x + 2 + 1
// }

// func (shape ReflectedL) FormRock(grid [][]byte) {
// 	// middle
// 	for i := 0; i < 3; i += 1 {
// 		grid[shape.left.x][shape.left.y+i] = '#'
// 	}
// 	grid[shape.left.x+1][shape.left.y+2] = '#'
// 	grid[shape.left.x+2][shape.left.y+2] = '#'
// }

// // endregion

// // region VerticalBar
// func (shape *VerticalBar) MoveLeft(grid [][]byte) {

// 	if shape.left.y == 0 {
// 		// check moving left
// 		return
// 	}
// 	for i := 0; i < 4; i += 1 {
// 		if grid[shape.left.x+i][shape.left.y-1] == '#' {
// 			return
// 		}
// 	}
// 	shape.left.y -= 1
// 	shape.bottom.y -= 1
// }

// func (shape *VerticalBar) MoveRight(grid [][]byte) {
// 	if shape.left.y == 6 {
// 		return
// 	}
// 	for i := 0; i < 4; i += 1 {
// 		// check moving right (3 blocks)
// 		if grid[shape.left.x+i][shape.left.y+1] == '#' {
// 			return
// 		}
// 	}
// 	shape.left.y += 1
// 	shape.bottom.y += 1
// }

// func (shape *VerticalBar) MoveBottom(grid [][]byte) bool {
// 	if shape.bottom.x == 0 || grid[shape.bottom.x-1][shape.bottom.y] == '#' {
// 		// if reached bottom or hit a rock
// 		return false
// 	}
// 	shape.left.x -= 1
// 	shape.bottom.x -= 1

// 	// fmt.Printf("Bottom = %s\n", shape.bottom)
// 	return true
// }

// func (shape VerticalBar) Top() int {
// 	return shape.left.x + 3 + 1
// }

// func (shape VerticalBar) FormRock(grid [][]byte) {
// 	// middle
// 	for i := 0; i < 4; i += 1 {
// 		grid[shape.left.x+i][shape.left.y] = '#'
// 	}
// }

// // endregion

// // region Square
// func (shape *Square) MoveLeft(grid [][]byte) {
// 	if shape.left.y == 0 || grid[shape.left.x][shape.left.y-1] == '#' || grid[shape.left.x+1][shape.left.y-1] == '#' {
// 		// check moving left
// 		return
// 	}
// 	shape.left.y -= 1
// 	shape.bottom.y -= 1
// 	// fmt.Printf("Left = %s\n", shape.bottom)
// }

// func (shape *Square) MoveRight(grid [][]byte) {
// 	if shape.left.y+1 == 6 || grid[shape.left.x][shape.left.y+1+1] == '#' || grid[shape.left.x+1][shape.left.y+1+1] == '#' {
// 		return
// 	}
// 	shape.left.y += 1
// 	shape.bottom.y += 1
// 	// fmt.Printf("Right = %s\n", shape.bottom)
// }

// func (shape *Square) MoveBottom(grid [][]byte) bool {
// 	if shape.bottom.x == 0 || grid[shape.bottom.x-1][shape.bottom.y] == '#' || grid[shape.bottom.x-1][shape.bottom.y+1] == '#' {
// 		// if reached bottom or hit a rock
// 		return false
// 	}
// 	shape.left.x -= 1
// 	shape.bottom.x -= 1
// 	return true
// }

// func (shape Square) Top() int {
// 	return shape.left.x + 1 + 1
// }

// func (shape Square) FormRock(grid [][]byte) {
// 	// middle
// 	for i := 0; i < 2; i += 1 {
// 		for j := 0; j < 2; j += 1 {
// 			grid[shape.left.x+i][shape.left.y+j] = '#'
// 		}
// 	}
// }

// // endregion
