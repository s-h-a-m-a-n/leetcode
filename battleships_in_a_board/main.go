package main

import "fmt"

func countBattleships(board [][]byte) int {
	var count int
	bc := map[int]struct{}{}

	for _, r := range board {
		for j, c := range r {
			if c == byte('.') {
				if _, ok := bc[j]; ok {
					delete(bc, j)
					count++
				}
				continue
			}

			if _, ok := bc[j-1]; ok {
				delete(bc, j-1)
			}
			bc[j] = struct{}{}
		}
	}

	count = count + len(bc)

	return count
}

func main() {
	board := [][]byte{
		{'X', '.', '.', 'X', '.', '.', '.', '.'},
		{'.', '.', '.', 'X', '.', 'X', 'X', 'X'},
		{'.', 'X', '.', 'X', '.', '.', '.', '.'},
	}

	fmt.Println(countBattleships(board))
}
