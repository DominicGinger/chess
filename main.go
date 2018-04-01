package main

import (
	"fmt"
)

type Square struct {
	coloured bool
	x        int
	y        int
}

type Board struct {
	rows [8][8]Square
}

var xAxis = [...]int{67, 68, 69, 70, 71, 72, 73, 74}
var yAxis = [...]int{1, 2, 3, 4, 5, 6, 7, 8}

func main() {
	board := Board{}

	coloured := true
	for i := range board.rows {
		coloured = !coloured
		for j := range board.rows[i] {
			board.rows[i][j] = Square{coloured, xAxis[i], yAxis[j]}
			if coloured {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
				fmt.Printf(" %d %d", board.rows[i][j].x, board.rows[i][j].y)
			coloured = !coloured
		}
		fmt.Println()
	}
}
