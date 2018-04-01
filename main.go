package main

import (
	"fmt"
)

var xAxis = "ABCDEFGH"
var yAxis = [...]int{1, 2, 3, 4, 5, 6, 7, 8}

type Square struct {
	coloured bool
	x        byte
	y        int
}

type Board struct {
	rows [8][8]Square
}

func (b *Board) print() {
	fmt.Printf(" ")
	for _, c := range xAxis {
		fmt.Printf("  %c ", c)
	}
	fmt.Println()

	for i := range b.rows {
		fmt.Printf("  ")
		for _, s := range b.rows[i] {
			if s.coloured {
				fmt.Printf("■■■■")
			} else {
				fmt.Printf("□□□□")
			}
		}
		fmt.Printf("\n%d ", b.rows[i][0].y)
		for _, s := range b.rows[i] {
			if s.coloured {
				fmt.Printf("■♔ ■")
			} else {
				fmt.Printf("□□□□")
			}
		}
		fmt.Printf("\n  ")
		for _, s := range b.rows[i] {
			if s.coloured {
				fmt.Printf("■■■■")
			} else {
				fmt.Printf("□□□□")
			}
		}

		fmt.Println()
	}
}

func newBoard() (board Board) {
	coloured := true

	for i := range board.rows {
		coloured = !coloured
		for j := range board.rows[i] {
			board.rows[i][j] = Square{coloured, xAxis[j], 9 - yAxis[i]}
			coloured = !coloured
		}
	}
	return
}

func main() {
	board := newBoard()
	board.print()
}
