package main

import "fmt"

type Rank []bool

type Chessboard map[string] Rank

func CountInRank(board Chessboard, rank string) int {
	fmt.Println("this is rank", rank)

	_, rankExists := board[rank]
	count := 0 
	if !rankExists {
		count = 0 
	}
	fmt.Println(board[rank])
	for _, squareOccupied := range(board[rank]) {
		fmt.Println(board[rank])
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
    fmt.Println(file)
    if file < 1 || file > 9 {
        count = 0 
        fmt.Println(count)
        return count 
    } else {
    	for _, row := range(cb) {
        fmt.Println(row)
        indexExist := row[file-1]
		if indexExist {
			count ++
		}
		}	
    }
	fmt.Println("count", count)
	return count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	count := 0
	for _, row := range(cb) {
		for range(row) {
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


func main() {
	// board := Chestboard [A:[true false true false false false false true] B:[false false false false true false false false] C:[false false true false false false false false] D:[false false false false false false false false] E:[false false false false false true false true] F:[false false false false false false false false] G:[false false false true false false false false] H:[true true true true true true false true]]
	// CountInRank(board, "A")
}