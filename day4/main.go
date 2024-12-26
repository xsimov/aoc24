package main

import (
	"fmt"
	"strings"
)

const (
	NORTH        = "N"
	NORTH_EAST   = "NE"
	EAST         = "E"
	SOUTH_EAST   = "SE"
	SOUTH        = "S"
	SOUTH_WEST   = "SW"
	WEST         = "W"
	NORTH_WEST   = "NW"
	WORD_TO_FIND = "XMAS"
)

var OutOfBoundsCell = Cell{letter: ".", x: -1, y: -1}

func main() {
	fileContent := parseFile("data.txt")
	rows := strings.Split(fileContent, "\n")
	rows = rows[:len(rows)-1]
	board := Board{grid: make([][]Cell, len(rows)), xmax: len(rows[0]) - 1, ymax: len(rows) - 1}

	for i := 0; i < len(rows); i++ {
	}

	for y, row := range rows {
		letters := strings.Split(row, "")
		board.grid[y] = make([]Cell, len(rows[y]))

		for x, letter := range letters {
			board.grid[y][x] = Cell{letter: letter, x: x, y: y, board: &board}
		}
	}
	fmt.Println(board)

	found := 0
	resultFound := make(chan int, 8)
	done := make(chan bool, 8)
	go wordsFoundInDirection(&board, NORTH, resultFound, done)
	go wordsFoundInDirection(&board, NORTH_EAST, resultFound, done)
	go wordsFoundInDirection(&board, EAST, resultFound, done)
	go wordsFoundInDirection(&board, SOUTH_EAST, resultFound, done)
	go wordsFoundInDirection(&board, SOUTH, resultFound, done)
	go wordsFoundInDirection(&board, SOUTH_WEST, resultFound, done)
	go wordsFoundInDirection(&board, WEST, resultFound, done)
	go wordsFoundInDirection(&board, NORTH_WEST, resultFound, done)

	for i := 0; i < 8; i++ {
		<-done
	}

	close(resultFound)

	for elem := range resultFound {
		fmt.Println("found:", elem)
		found += elem
	}

	fmt.Println("Part 1: Total XMAS found:", found)

	foundXMas := make(chan bool)

	startedJobs := 0
	for x := 0; x <= board.xmax; x++ {
		for y := 0; y <= board.ymax; y++ {
			current := board.grid[y][x]
			if current.letter == string('A') {
				go checkXFormedMAS(current, foundXMas)
				startedJobs += 1
			}
		}
	}

	xFormedMAS := 0

	for i := 0; i < startedJobs; i++ {
		found := <-foundXMas
		if found {
			xFormedMAS += 1
		}
	}

	fmt.Println("Part 2: Total XMAS found:", xFormedMAS)
	fmt.Println("Part 2: Total As found:", startedJobs)
}

func checkXFormedMAS(c Cell, found chan<- bool) {
	ne := c.Navigate(NORTH_EAST)
	sw := c.Navigate(SOUTH_WEST)
	nw := c.Navigate(NORTH_WEST)
	se := c.Navigate(SOUTH_EAST)

	var nwse, nesw bool
	if (ne.letter == "M" && sw.letter == "S") || (ne.letter == "S" && sw.letter == "M") {
		nwse = true
	}
	if (nw.letter == "M" && se.letter == "S") || (nw.letter == "S" && se.letter == "M") {
		nesw = true
	}
	found <- (nwse && nesw)
}

func wordsFoundInDirection(b *Board, d string, resultFound chan<- int, done chan<- bool) {
	counter := 0
	for x := 0; x <= b.xmax; x++ {
		for y := 0; y <= b.ymax; y++ {
			current := b.grid[y][x]

			if current.letter == string(WORD_TO_FIND[0]) {
				searching := true
				newCell := current

				for i := 1; (i < len(WORD_TO_FIND)) && searching; i++ {
					newCell = newCell.Navigate(d)
					if newCell.letter != string(WORD_TO_FIND[i]) {
						searching = false
					}
				}
				if searching {
					counter += 1
				}
			}
		}
	}
	resultFound <- counter
	done <- true
}

type Board struct {
	grid       [][]Cell
	xmax, ymax int
}

func (b Board) String() string {
	printedBoard := make([]string, b.xmax*b.ymax)

	printedBoard = append(printedBoard, "Board:\n")
	for _, row := range b.grid {
		for _, cell := range row {
			printedBoard = append(printedBoard, cell.letter)
		}
		printedBoard = append(printedBoard, "\n")
	}
	maxes := fmt.Sprintf("X Max: %d\nY Max: %d\n", b.xmax, b.ymax)
	printedBoard = append(printedBoard, maxes)

	return strings.Join(printedBoard, "")
}

func (b Board) OutOfBounds(x, y int) bool {
	return x < 0 || y < 0 || x > b.xmax || y > b.ymax
}

type Cell struct {
	board  *Board
	letter string
	x, y   int
}

func (c Cell) String() string {
	return fmt.Sprintf("Cell {\n letter: %s,\n x: %d,\n y: %d\n}", c.letter, c.x, c.y)
}

func (c Cell) OutOfBounds() bool {
	return c.x == -1
}

func (c Cell) Navigate(d string) Cell {
	switch d {
	case NORTH:
		if !c.board.OutOfBounds(c.x, c.y-1) {
			return c.board.grid[c.y-1][c.x]
		}
	case NORTH_EAST:
		if !c.board.OutOfBounds(c.x+1, c.y-1) {
			return c.board.grid[c.y-1][c.x+1]
		}
	case EAST:
		if !c.board.OutOfBounds(c.x+1, c.y) {
			return c.board.grid[c.y][c.x+1]
		}
	case SOUTH_EAST:
		if !c.board.OutOfBounds(c.x+1, c.y+1) {
			return c.board.grid[c.y+1][c.x+1]
		}
	case SOUTH:
		if !c.board.OutOfBounds(c.x, c.y+1) {
			return c.board.grid[c.y+1][c.x]
		}
	case SOUTH_WEST:
		if !c.board.OutOfBounds(c.x-1, c.y+1) {
			return c.board.grid[c.y+1][c.x-1]
		}
	case WEST:
		if !c.board.OutOfBounds(c.x-1, c.y) {
			return c.board.grid[c.y][c.x-1]
		}
	case NORTH_WEST:
		if !c.board.OutOfBounds(c.x-1, c.y-1) {
			return c.board.grid[c.y-1][c.x-1]
		}
	}
	return OutOfBoundsCell
}
