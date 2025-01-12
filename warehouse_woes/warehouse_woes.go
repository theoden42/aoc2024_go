package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("warehouse_woes.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	input_grid := [][]byte{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		input_grid = append(input_grid, []byte(line))
	}
	// fmt.Println(input_grid)
	ix := 0
	iy := 0
	for i := 0; i < len(input_grid); i++ {
		for j := 0; j < len(input_grid[i]); j++ {
			if input_grid[i][j] == '@' {
				ix = i
				iy = j
			}
		}
	}
	dir := make(map[byte][]int)
	dir['>'] = []int{0, 1}
	dir['^'] = []int{-1, 0}
	dir['<'] = []int{0, -1}
	dir['v'] = []int{1, 0}
	for scanner.Scan() {
		line := scanner.Text()
		for _, c := range line {
			dx := dir[byte(c)][0]
			dy := dir[byte(c)][1]
			nx, ny := ix, iy
			for ; nx < len(input_grid) && ny < len(input_grid[0]); nx, ny = nx+dx, ny+dy {
				if input_grid[nx][ny] == '#' || input_grid[nx][ny] == '.' {
					break
				}
			}
			if input_grid[nx][ny] == '#' {
				continue
			}
			for ; nx != ix || ny != iy; nx, ny = nx-dx, ny-dy {
				input_grid[nx][ny] = 'O'
			}
			input_grid[ix][iy] = '.'
			ix, iy = ix+dx, iy+dy
			input_grid[ix][iy] = '@'
			// fmt.Println(input_grid, dx, dy)
		}
	}
	// fmt.Println(input_grid)
	final_ans := 0
	for i := 0; i < len(input_grid); i++ {
		for j := 0; j < len(input_grid[0]); j++ {
			if input_grid[i][j] == 'O' {
				final_ans += 100*i + j
			}
		}
	}
	fmt.Println(final_ans)
}
