package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("ceres_search.txt")
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
			if grid[i][j] == 'X' {
				if j < len(grid[i])-3 && grid[i][j:j+4] == "XMAS" {
					final_ans += 1
				}
				if j >= 3 && grid[i][j-3:j+1] == "SAMX" {
					final_ans += 1
				}
				if i >= 3 && grid[i-1][j] == 'M' && grid[i-2][j] == 'A' && grid[i-3][j] == 'S' {
					final_ans += 1
				}
				if i < len(grid)-3 && grid[i+1][j] == 'M' && grid[i+2][j] == 'A' && grid[i+3][j] == 'S' {
					final_ans += 1
				}
				if i >= 3 && j >= 3 && grid[i-1][j-1] == 'M' && grid[i-2][j-2] == 'A' && grid[i-3][j-3] == 'S' {
					final_ans += 1
				}
				if i < len(grid)-3 && j < len(grid[i])-3 && grid[i+1][j+1] == 'M' && grid[i+2][j+2] == 'A' && grid[i+3][j+3] == 'S' {
					final_ans += 1
				}
				if i >= 3 && j < len(grid[i])-3 && grid[i-1][j+1] == 'M' && grid[i-2][j+2] == 'A' && grid[i-3][j+3] == 'S' {
					final_ans += 1
				}
				if i < len(grid)-3 && j >= 3 && grid[i+1][j-1] == 'M' && grid[i+2][j-2] == 'A' && grid[i+3][j-3] == 'S' {
					final_ans += 1
				}
			}
		}
	}

	fmt.Println(final_ans)
}
