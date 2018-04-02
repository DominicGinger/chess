package main

func main() {
	board := newBoard()
	for i := 0; i < 8; i++ {
		board.addPiece(i, 1, newPawn(true))
		board.addPiece(i, 6, newPawn(false))
	}
	board.print()
}
