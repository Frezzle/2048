package main

import (
	"fmt"
	"math/rand"
	"os"
	"text/tabwriter"
)

const empty = -1

var grid = [4][4]int{
	{empty, empty, empty, empty},
	{empty, empty, empty, empty},
	{empty, empty, empty, empty},
	{empty, empty, empty, empty},
}

var w = tabwriter.NewWriter(os.Stdout, 2, 1, 1, ' ', 0)

func main() {
	spawnNumberSomewhere()
	spawnNumberSomewhere()
	spawnNumberSomewhere()
	for {
		printGrid()
		gridBefore := grid
		move()
		if !gridsAreEqual(grid, gridBefore) {
			spawnNumberSomewhere()
		}

		if numberExists(2048) {
			printGrid()
			fmt.Println("You win!")
			return
		}

		if gameIsOver() {
			printGrid()
			fmt.Println("Game over.")
			return
		}
	}
}

func printGrid() {
	fmt.Fprintln(w)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			value := grid[i][j]
			if value == empty {
				fmt.Fprint(w, "_\t ")
			} else {
				fmt.Fprintf(w, "%v\t ", value)
			}
		}
		fmt.Fprintln(w)
	}
	fmt.Fprintln(w)
	w.Flush()
}

func move() {
	var direction string
	fmt.Scanln(&direction)
	switch direction {
	case "w":
		moveUp(&grid)
	case "s":
		moveDown(&grid)
	case "a":
		moveLeft(&grid)
	case "d":
		moveRight(&grid)
	}
}

func moveUp(grid *[4][4]int) {
	for i := 0; i < 4; i++ {
		moveLine(&grid[0][i], &grid[1][i], &grid[2][i], &grid[3][i])
	}
}

func moveDown(grid *[4][4]int) {
	for i := 0; i < 4; i++ {
		moveLine(&grid[3][i], &grid[2][i], &grid[1][i], &grid[0][i])
	}
}

func moveLeft(grid *[4][4]int) {
	for i := 0; i < 4; i++ {
		moveLine(&grid[i][0], &grid[i][1], &grid[i][2], &grid[i][3])
	}
}

func moveRight(grid *[4][4]int) {
	for i := 0; i < 4; i++ {
		moveLine(&grid[i][3], &grid[i][2], &grid[i][1], &grid[i][0])
	}
}

// a, b, c, and d represent 4 numbers next to each other in a line.
// moveLine moves them towards a.
// They are moved and merged according to the rules of 2048.
func moveLine(a, b, c, d *int) {
	line := [4]*int{a, b, c, d}
	merged := [3]bool{}
	for i := 1; i < 4; i++ {
		if *line[i] == empty {
			continue // no number to shift
		}
		// shift number until blocked or merged
		for j := i - 1; j >= 0; j-- {
			if *line[j] == empty {
				// shift number
				*line[j] = *line[j+1]
				*line[j+1] = empty
			} else if *line[j] != *line[j+1] {
				// blocked
				break
			} else if !merged[j] {
				// merge identical numbers
				*line[j] += *line[j+1]
				*line[j+1] = empty
				merged[j] = true
				break
			}
		}
	}
}

func spawnNumberSomewhere() {
	// get all empty locations
	emptys := [][2]int{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if grid[i][j] == empty {
				emptys = append(emptys, [2]int{i, j})
			}
		}
	}
	if len(emptys) == 0 {
		return
	}

	// spawn at random empty place
	location := emptys[rand.Intn(len(emptys))]
	grid[location[0]][location[1]] = 2
}

func gridsAreEqual(a, b [4][4]int) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func numberExists(x int) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if grid[i][j] == x {
				return true
			}
		}
	}
	return false
}

func gameIsOver() bool {
	clone := grid
	moveUp(&clone)
	if !gridsAreEqual(grid, clone) {
		return false
	}
	moveDown(&clone)
	if !gridsAreEqual(grid, clone) {
		return false
	}
	moveLeft(&clone)
	if !gridsAreEqual(grid, clone) {
		return false
	}
	moveRight(&clone)
	return gridsAreEqual(grid, clone)
}
