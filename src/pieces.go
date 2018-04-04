package main

type move struct {
	y int
	x int
}

type ChessPiece interface {
	getColoured() bool
	getSymbol() string 
	validMove(from, to move) bool
	// validMoves(m move) []move
}

var symbols = [...]string{"♙", "♟", "♘", "♞", "♗", "♝", "♖", "♜", "♕", "♛", "♔", "♚"}
func getSymbol(idx int, coloured bool) string {
	if coloured {
		return symbols[idx]
	} else {
		return symbols[idx + 1]
	}
}

type Pawn struct {
	coloured bool
}

func (p *Pawn) validMove(from, to move) bool {
	if p.coloured {
		return to.y == from.y + 1 || to.y == from.y + 2
	}
	return to.y == from.y - 1 || to.y == from.y - 2
}

func (p *Pawn) getColoured() bool {
	return p.coloured
}

func (p *Pawn) getSymbol() string {
	return getSymbol(0, p.coloured)
}


// func newKnight(coloured bool) (p Piece) {
// 	return Piece{coloured, getSymbol(2, coloured)}
// }

// func newBishop(coloured bool) (p Piece) {
// 	return Piece{coloured, getSymbol(4, coloured)}
// }

// func newRook(coloured bool) (p Piece) {
// 	return Piece{coloured, getSymbol(6, coloured)}
// }

// func newQueen(coloured bool) (p Piece) {
// 	return Piece{coloured, getSymbol(8, coloured)}
// }

// func newKing(coloured bool) (p Piece) {
// 	return Piece{coloured, getSymbol(10, coloured)}
// }

