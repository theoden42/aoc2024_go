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
	for num_moves := 1000; num_moves <= 10000; num_moves++ {
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
				if grid[i][j] == "x" && i < mx-1 && grid[i+1][j] == "x" {
					cnt += 1
				}
				if grid[i][j] == "x" && i > 0 && grid[i-1][j] == "x" {
					cnt += 1
				}
				if grid[i][j] == "x" && j < my-1 && grid[i][j+1] == "x" {
					cnt += 1
				}
				if grid[i][j] == "x" && i > 0 && grid[i-1][j] == "x" {
					cnt += 1
				}
			}
		}
		if cnt > 300 {
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
