package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("hoof_it.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	grid := [][]byte{}
	for scanner.Scan() {
		line := []byte(scanner.Text())
		grid = append(grid, line)
	}
	num_rows := len(grid)
	num_cols := len(grid[0])
	visited := make([][]int, num_rows)
	for i := 0; i < num_rows; i++ {
		visited[i] = make([]int, num_cols)
		for j := 0; j < num_cols; j++ {
			visited[i][j] = -1
		}
	}
	// for part 1 do a single dfs traversal without dp
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if visited[i][j] != -1 {
			return visited[i][j]
		}
		if grid[i][j] == '9' {
			visited[i][j] = 1
			return 1
		}
		value := 0
		if i > 0 && grid[i-1][j] == grid[i][j]+1 {
			value += dfs(i-1, j)
		}
		if j > 0 && grid[i][j-1] == grid[i][j]+1 {
			value += dfs(i, j-1)
		}
		if i < num_rows-1 && grid[i+1][j] == grid[i][j]+1 {
			value += dfs(i+1, j)
		}
		if j < num_cols-1 && grid[i][j+1] == grid[i][j]+1 {
			value += dfs(i, j+1)
		}
		visited[i][j] = value
		return value
	}
	final_ans := 0
	for i := 0; i < num_rows; i++ {
		for j := 0; j < num_cols; j++ {
			if grid[i][j] == '0' {
				for i := 0; i < num_rows; i++ {
					for j := 0; j < num_cols; j++ {
						visited[i][j] = -1
					}
				}
				final_ans += dfs(i, j)
			}
		}
	}
	fmt.Println(final_ans)
}
