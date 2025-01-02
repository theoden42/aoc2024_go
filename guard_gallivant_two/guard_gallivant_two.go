package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func is_in_cycle(num_rows int, num_cols int, place_map [][]rune) bool {
	directions := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // up, right, down, left
	visited := make(map[[2]int]bool)
	visited_with_direction := make(map[[3]int]bool)
	start_pos := [2]int{0, 0}
	start_dir := 0
	for i := 0; i < num_rows; i++ {
		for j := 0; j < num_cols; j++ {
			if place_map[i][j] == '<' {
				start_pos = [2]int{i, j}
				start_dir = 3
				break
			}
			if place_map[i][j] == '>' {
				start_pos = [2]int{i, j}
				start_dir = 1
				break
			}
			if place_map[i][j] == '^' {
				start_pos = [2]int{i, j}
				start_dir = 0
				break
			}
			if place_map[i][j] == 'v' {
				start_pos = [2]int{i, j}
				start_dir = 2
				break
			}
		}
	}
	visited[start_pos] = true
	visited_with_direction[[3]int{start_pos[0], start_pos[1], start_dir}] = true
	for true {
		next_pos := [2]int{start_pos[0] + directions[start_dir][0], start_pos[1] + directions[start_dir][1]}
		if next_pos[0] < 0 || next_pos[0] >= num_rows || next_pos[1] < 0 || next_pos[1] >= num_cols {
			break
		}
		if place_map[next_pos[0]][next_pos[1]] == '#' {
			start_dir = (start_dir + 1) % 4
			continue
		} else if visited_with_direction[[3]int{next_pos[0], next_pos[1], start_dir}] {
			return true
		} else {
			visited[next_pos] = true
			visited_with_direction[[3]int{next_pos[0], next_pos[1], start_dir}] = true
			start_pos = next_pos
		}
	}
	return false
}

func main() {
	file, err := os.Open("guard_gallivant.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	place_map := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		place_map = append(place_map, []rune(line))
	}
	count_pos := 0
	for i := 0; i < len(place_map); i++ {
		for j := 0; j < len(place_map[i]); j++ {
			if place_map[i][j] != '.' {
				continue
			}
			place_map[i][j] = '#'
			if is_in_cycle(len(place_map), len(place_map[0]), place_map) {
				count_pos++
			}
			place_map[i][j] = '.'
		}
	}
	fmt.Println(count_pos)
}
