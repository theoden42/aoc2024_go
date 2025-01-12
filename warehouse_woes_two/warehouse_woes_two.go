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
		new_line := []byte{}
		for _, c := range line {
			if c == 'O' {
				new_line = append(new_line, '[')
				new_line = append(new_line, ']')
			} else if c == '@' {
				new_line = append(new_line, '@')
				new_line = append(new_line, '.')
			} else {
				new_line = append(new_line, byte(c))
				new_line = append(new_line, byte(c))
			}
		}
		if len(line) == 0 {
			break
		}
		input_grid = append(input_grid, []byte(new_line))
	}
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
	fmt.Println(input_grid)
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
			new_input_grid := input_grid
			var check func(x, y, dx, dy int) bool
			var move func(x, y, dx, dy int)
			check = func(x, y, dx, dy int) bool {
				if dx == 0 {
					if input_grid[x][y] == '@' {
						return check(x, y+dy, dx, dy)
					}
					if input_grid[x][y+2*dy] == '#' {
						return false
					} else {
						if input_grid[x][y+2*dy] == '.' {
							return true
						}
						return check(x, y+2*dy, dx, dy)
					}
				} else {
					if input_grid[x+dx][y] == '.' {
						return true
					} else if input_grid[x+dx][y] == '#' {
						return false
					} else if input_grid[x+dx][y] == ']' {
						return (check(x+dx, y, dx, dy) && check(x+dx, y-1, dx, dy))
					} else {
						return (check(x+dx, y, dx, dy) && check(x+dx, y+1, dx, dy))
					}
				}
			}
			move = func(x, y, dx, dy int) {
				fmt.Println(x, y, input_grid[x][y])
				if input_grid[x][y] == '.' {
					return
				}
				if dx == 0 {
					if input_grid[x][y] == '@' {
						move(x, y+dy, dx, dy)
						new_input_grid[x][y] = '.'
						new_input_grid[x][y+dy] = '@'
						return
					}
					move(x, y+2*dy, dx, dy)
					new_input_grid[x][min(y+dy, y+2*dy)] = '['
					new_input_grid[x][max(y+dy, y+2*dy)] = ']'
					new_input_grid[x][y] = '.'
				} else {
					fmt.Println("here", x, y, x+dx, y)
					if input_grid[x+dx][y] == ']' {
						move(x+dx, y, dx, dy)
						move(x+dx, y-1, dx, dy)
						new_input_grid[x+dx][y] = input_grid[x][y]
						new_input_grid[x][y] = '.'
					} else if input_grid[x+dx][y] == '[' {
						move(x+dx, y, dx, dy)
						move(x+dx, y+1, dx, dy)
						new_input_grid[x+dx][y] = input_grid[x][y]
						new_input_grid[x][y] = '.'
					} else {
						new_input_grid[x+dx][y] = input_grid[x][y]
						new_input_grid[x][y] = '.'
					}

				}
			}
			if !check(ix, iy, dx, dy) {
				continue
			}
			fmt.Println("before move", ix, iy)
			move(ix, iy, dx, dy)
			input_grid = new_input_grid
			fmt.Println(input_grid, dx, dy)
			ix, iy = ix+dx, iy+dy
			fmt.Println(ix, iy, input_grid[ix+1][iy])
		}
	}
	// fmt.Println(input_grid)
	final_ans := 0
	for i := 0; i < len(input_grid); i++ {
		for j := 0; j < len(input_grid[0]); j++ {
			if input_grid[i][j] == '[' {
				fmt.Println(i, j)
				final_ans += 100*i + j
			}
		}
	}
	fmt.Println(final_ans)
}
