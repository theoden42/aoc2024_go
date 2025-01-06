package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("garden_groups.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	input_grid := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		input_grid = append(input_grid, line)
	}
	num_rows := len(input_grid)
	num_cols := len(input_grid[0])
	visited := make([][]bool, num_rows)
	for i := 0; i < num_rows; i++ {
		visited[i] = make([]bool, num_cols)

	}
	can_travel := func(i, j, x, y int) bool {
		return (i >= 0 && i < num_rows && j >= 0 && j < num_cols && !visited[i][j] && input_grid[i][j] == input_grid[x][y])
	}
	is_boundary := func(i, j, x, y int) bool {
		return (i < 0 || j < 0 || i >= num_rows || j >= num_cols || input_grid[i][j] != input_grid[x][y])
	}
	var dfs func(i, j int) []int
	dfs = func(i, j int) []int {
		visited[i][j] = true
		count_area := 1
		count_peri := 0
		if is_boundary(i+1, j, i, j) {
			count_peri += 1
		}
		if is_boundary(i, j+1, i, j) {
			count_peri += 1
		}
		if is_boundary(i-1, j, i, j) {
			count_peri += 1
		}
		if is_boundary(i, j-1, i, j) {
			count_peri += 1
		}
		if can_travel(i+1, j, i, j) {
			val := dfs(i+1, j)
			count_area += val[0]
			count_peri += val[1]
		}
		if can_travel(i, j+1, i, j) {
			val := dfs(i, j+1)
			count_area += val[0]
			count_peri += val[1]
		}
		if can_travel(i, j-1, i, j) {
			val := dfs(i, j-1)
			count_area += val[0]
			count_peri += val[1]
		}
		if can_travel(i-1, j, i, j) {
			val := dfs(i-1, j)
			count_area += val[0]
			count_peri += val[1]
		}
		return []int{count_area, count_peri}
	}
	final_ans := 0
	for i := 0; i < num_rows; i++ {
		for j := 0; j < num_cols; j++ {
			if !visited[i][j] {
				val := dfs(i, j)
				temp_ans := val[0] * val[1]
				final_ans += temp_ans
			}
		}
	}
	fmt.Println(final_ans)
}
