package main

import (
	"fmt"
)

func main() {
	board := newBoard()
	for i := 0; i < 8; i++ {
		board.addPiece(i, 1, newPawn(true))
		board.addPiece(i, 6, newPawn(false))
	}

	fmt.Println(board.rows[0][1].piece.validMove(move{0, 1}, move{0, 2}))
	// type placement struct {
	// 	y int
	// 	coloured bool
	// }
	// for _, p := range []placement{placement{0, true}, placement{7, false}} {
	// 	board.addPiece(0, p.y, newRook(p.coloured))
	// 	board.addPiece(1, p.y, newKnight(p.coloured))
	// 	board.addPiece(2, p.y, newBishop(p.coloured))
	// 	board.addPiece(3, p.y, newQueen(p.coloured))
	// 	board.addPiece(4, p.y, newKing(p.coloured))
	// 	board.addPiece(5, p.y, newBishop(p.coloured))
	// 	board.addPiece(6, p.y, newKnight(p.coloured))
	// 	board.addPiece(7, p.y, newRook(p.coloured))
	// }
	board.print()
}
