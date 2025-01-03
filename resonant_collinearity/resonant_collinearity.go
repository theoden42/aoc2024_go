package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("resonant_collinearity.txt")
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
	character_pos := make(map[byte][][2]int)
	for i := 0; i < num_rows; i++ {
		for j := 0; j < num_cols; j++ {
			if input_grid[i][j] == '.' {
				continue
			}
			character_pos[input_grid[i][j]] = append(character_pos[input_grid[i][j]], [2]int{i, j})
		}
	}
	antinodes := make(map[[2]int]bool)
	for _, positions := range character_pos {
		sort.Slice(positions, func(i, j int) bool {
			return positions[i][0] < positions[j][0]
		})
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				row_dif := positions[j][0] - positions[i][0]
				row_val := positions[j][0] + row_dif
				row_val_before := positions[i][0] - row_dif
				col_dif := positions[j][1] - positions[i][1]
				if col_dif < 0 {
					col_dif = -col_dif
				}
				col_val := max(positions[j][1], positions[i][1]) + col_dif
				col_val_before := min(positions[j][1], positions[i][1]) - col_dif
				if positions[j][1] < positions[i][1] {
					if col_val_before >= 0 && row_val < num_rows {
						antinodes[[2]int{row_val, col_val_before}] = true
					}
					if col_val < num_cols && row_val_before >= 0 {
						antinodes[[2]int{row_val_before, col_val}] = true
					}
				} else {
					if col_val < num_cols && row_val < num_rows {
						antinodes[[2]int{row_val, col_val}] = true
					}
					if col_val_before >= 0 && row_val_before >= 0 {
						antinodes[[2]int{row_val_before, col_val_before}] = true
					}
				}
			}
		}
	}
	fmt.Println(len(antinodes))
}
