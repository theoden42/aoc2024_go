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
	fmt.Println(character_pos)
	antinodes := make(map[[2]int]bool)
	for _, positions := range character_pos {
		sort.Slice(positions, func(i, j int) bool {
			return positions[i][0] < positions[j][0]
		})
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				// calculate equation of line passing through two points
				// (x2 - x1)(y - y1) =  (x - x1)(y2 - y1)
				// this is very slow but easy to implement
				for x := 0; x < num_rows; x++ {
					for y := 0; y < num_cols; y++ {
						left_val := (positions[j][0] - positions[i][0]) * (y - positions[i][1])
						right_val := (x - positions[i][0]) * (positions[j][1] - positions[i][1])
						if left_val == right_val {
							antinodes[[2]int{x, y}] = true
						}
					}
				}
			}
		}
	}
	fmt.Println(len(antinodes))
}
