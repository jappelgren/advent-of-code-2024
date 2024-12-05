package main

import (
	"fmt"
)

func main() {
	f := ParseFileToBytes("input.txt")

	first, err := FirstStar(f)

	if err != nil {
		fmt.Printf("An error occurred: %v", err)
		return
	}
	fmt.Printf("First star results: %v\n", first)

	second, err := SecondStar(f)

	if err != nil {
		fmt.Printf("An error occurred: %v", err)
		return
	}
	fmt.Printf("Second star results: %v", second)
}

func FirstStar(f []byte) (int, error){
	res := 0
	i := 0
	grid := [][]byte{}
	row := []byte{}
	for i < len(f) {
		if f[i] == '\r' {
			i += 2
			grid = append(grid, row)
			row = []byte{}
		} else {
			row = append(row,f[i])
			i++
		}
	}

	for i, v := range grid {
		for j, jv := range v {
			if jv == 'X' {
				wordsAtPosition := getWords(i, j, grid)
				for _, v := range wordsAtPosition {
					if v == "XMAS" {
						res++
					}
				}
			}
		}
	}

	return res, nil
}

func SecondStar(f []byte) (int, error){
	res := 0
	i := 0
	grid := [][]byte{}
	row := []byte{}

	for i < len(f) {
		if f[i] == '\r' {
			i += 2
			grid = append(grid, row)
			row = []byte{}
		} else {
			row = append(row,f[i])
			i++
		}
	}

	for i, v := range grid {
		if i == 0 || i == len(grid)-1 {
			continue
		}
		for j, jv := range v {
			if j == 0 || j == len(grid[i])-1 {
				continue
			}
			if jv == 'A' {
				mascount := 0
				wordsAtPosition := map[string]string{
					"ur": fmt.Sprintf("%vA%v", string(grid[i+1][j-1]), string(grid[i-1][j+1])),
					"ul": fmt.Sprintf("%vA%v", string(grid[i+1][j+1]), string(grid[i-1][j-1])),
				}

				for _, v := range wordsAtPosition {
					if v == "MAS" || v == "SAM" {
						mascount++
					}
					if mascount == 2 {
						res++
						mascount = 0
					}
				}
			}
		}
	}

	return res, nil
}

func getWords(i int, j int, grid [][]byte) map[string]string{
	words := map[string]string{
		"u": "",
		"ur": "",
		"r": "",
		"dr": "",
		"d" : "",
		"dl" : "",
		"l" : "",
		"ul" : "",
	}
	
	//up
	if i > 2 {
		words["u"] = fmt.Sprintf("X%v%v%v", string(grid[i-1][j]), string(grid[i-2][j]), string(grid[i-3][j]))
	}
	//up right
	if i > 2 && j < len(grid[i]) - 3 {
		words["ur"] = fmt.Sprintf("X%v%v%v", string(grid[i-1][j+1]), string(grid[i-2][j+2]), string(grid[i-3][j+3]))
	}
	//right
	if j < len(grid[i]) - 3 {
		words["r"] = fmt.Sprintf("X%v%v%v", string(grid[i][j+1]), string(grid[i][j+2]), string(grid[i][j+3]))
	}
	//down right
	if j < len(grid[i]) - 3 && i < len(grid) - 3 {
		words["dr"] = fmt.Sprintf("X%v%v%v", string(grid[i+1][j+1]), string(grid[i+2][j+2]), string(grid[i+3][j+3]))
	}
	//down
	if i < len(grid) - 3 {
		words["d"] = fmt.Sprintf("X%v%v%v", string(grid[i+1][j]), string(grid[i+2][j]), string(grid[i+3][j]))
	}
	//down left
	if i < len(grid) - 3 && j > 2 {
		words["dl"] = fmt.Sprintf("X%v%v%v", string(grid[i+1][j-1]), string(grid[i+2][j-2]), string(grid[i+3][j-3]))
	}
	//left
	if j > 2 {
		words["l"] = fmt.Sprintf("X%v%v%v", string(grid[i][j-1]), string(grid[i][j-2]), string(grid[i][j-3]))
	}
	//up left
	if j > 2 && i > 2 {
		words["ul"] = fmt.Sprintf("X%v%v%v", string(grid[i-1][j-1]), string(grid[i-2][j-2]), string(grid[i-3][j-3]))
	}

	return words
}