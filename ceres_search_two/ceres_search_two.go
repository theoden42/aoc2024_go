package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("ceres_search_two.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var grid []string
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	final_ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'A' && i < len(grid)-1 && i > 0 && j < len(grid[i])-1 && j > 0 {
				first_diag := make(map[byte]int)
				second_diag := make(map[byte]int)
				first_diag[grid[i-1][j-1]] += 1
				first_diag[grid[i+1][j+1]] += 1
				second_diag[grid[i-1][j+1]] += 1
				second_diag[grid[i+1][j-1]] += 1
				if first_diag['M'] == 1 && first_diag['S'] == 1 && second_diag['M'] == 1 && second_diag['S'] == 1 {
					final_ans += 1
				}
			}
		}
	}
	fmt.Println(final_ans)
}
