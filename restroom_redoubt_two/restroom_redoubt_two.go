package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("restroom_redoubt.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	const mx = 101
	const my = 103
	robots := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		var ax, ay, vx, vy int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &ax, &ay, &vx, &vy)
		robots = append(robots, []int{ax, ay, vx, vy})
	}
	for num_moves := 10; num_moves <= 5000; num_moves++ {
		grid := [mx][my]string{}
		for i := 0; i < mx; i++ {
			for j := 0; j < my; j++ {
				grid[i][j] = "."
			}
		}
		for _, robot := range robots {
			ax := robot[0]
			ay := robot[1]
			vx := robot[2]
			vy := robot[3]
			nx := (ax + (num_moves*(vx+mx)%mx)%mx + mx) % mx
			ny := (ay + (num_moves*(vy+my)%my)%my + my) % my
			grid[nx][ny] = "x"
		}
		cnt := 0
		for i := 0; i < mx; i++ {
			for j := 0; j < my; j++ {
				if grid[i][j] == grid[i][my-1-j] || grid[mx-1-i][j] == grid[i][j] {
					cnt += 1
				}
			}
		}

		if cnt > 2000 {
			fmt.Println(num_moves)
			for i := 0; i < mx; i++ {
				for j := 0; j < my; j++ {
					fmt.Print(grid[i][j])
				}
				fmt.Println()
			}
		}
	}
}
