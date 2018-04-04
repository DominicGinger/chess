package main

type move struct {
	x int
	y int
}

type ChessPiece interface {
	getColoured() bool
	getSymbol() string 
	validMove(from, to move) bool
	validMoves(m move) []move
}

var symbols = [...]string{"♙", "♟", "♘", "♞", "♗", "♝", "♖", "♜", "♕", "♛", "♔", "♚"}
func getSymbol(idx int, coloured bool) string {
	if coloured {
		return symbols[idx]
	} else {
		return symbols[idx + 1]
	}
}

type Piece struct {
	coloured bool
}

func (p *Piece) getColoured() bool {
	return p.coloured
}

type Pawn struct {
	Piece
}

func (p *Pawn) getSymbol() string {
	return getSymbol(0, p.coloured)
}

func (p *Pawn) validMove(from, to move) bool {
	if p.coloured {
		return to.y == from.y + 1 || to.y == from.y + 2
	}
	return to.y == from.y - 1 || to.y == from.y - 2
}

func (p *Pawn) validMoves(from move) []move {
	return []move{
		move{from.x, from.y + 1},
		move{from.x, from.y + 2},
	}
}

type Knight struct {
	Piece
}

func (p *Knight) getSymbol() string {
	return getSymbol(2, p.coloured)
}

func (_ *Knight) validMove(from, to move) bool {
	return ((to.y == from.y + 2 || to.y == from.y - 2) && (to.x == from.x + 1 || to.x == from.x - 1)) ||
		((to.x == from.x + 2 || to.x == from.x - 2) && (to.y == from.y + 1 || to.y == from.y - 1))
}

func (_ *Knight) validMoves(from move) []move {
	return []move{
		move{from.x, from.y + 1},
		move{from.x, from.y + 2},
	}
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

