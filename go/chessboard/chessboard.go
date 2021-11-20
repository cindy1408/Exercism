package chessboard

import "fmt"


// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool
// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard map[string] Rank

func CountInRank(board Chessboard, rank string) int {
    fmt.Println("this is rank", rank)
    fmt.Println("this is chessboard", board)
	_, rankExists := board[rank]
	count := 0 
	if !rankExists {
		count = 0 
	}
	fmt.Println(board[rank])
	for _, squareOccupied := range(board[rank]) {
		// fmt.Println(board[rank])
		if squareOccupied {
			count ++ 
		}
	}
	return count 
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	count := 0
    if file < 1 || file > 9 {
        count = 0 
        return count 
    } else {
    	for _, row := range(cb) {
        indexExist := row[file-1]
		if indexExist {
			count ++
		}
		}	
    }
	return count
}
// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	count := 0
    fmt.Println(cb)
	for _, row := range(cb) {
		for _,_ = range(row) {
				count++
		}
	}
	return count
}


// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
		count := 0
	for _, row := range(cb) {
		for _, index := range(row) {
			if index {
				count++
			}
		}
	}
	return count
}
