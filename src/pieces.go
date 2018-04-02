package main

type Pawn struct {
	coloured bool
	symbol string
}

func newPawn(coloured bool) (p Pawn) {
	p.coloured = coloured
	if coloured {
		p.symbol = "♙"
	} else {
		p.symbol = "♟"
	}
	return
}
