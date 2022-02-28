package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	//panic("Please implement CountInRank()")
	rankArr, rankExists := cb[rank]
	occupiedSquares := 0

	if !rankExists {
		return 0
	}

	for _, val := range rankArr {
		if val {
			occupiedSquares++
		}
	}
	return occupiedSquares

}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	//panic("Please implement CountInFile()")
	occupiedSquares := 0
	if file < 1 || file > 8 {
		return 0
	}

	for _, val := range cb {
		if val[file-1] {
			occupiedSquares++
		}
	}
	return occupiedSquares

}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	//panic("Please implement CountAll()")
	return len(cb) * len(cb)
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	panic("Please implement CountOccupied()")
}
