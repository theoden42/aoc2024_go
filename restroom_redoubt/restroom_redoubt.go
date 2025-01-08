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
	mx := 101
	my := 103
	quads := [4]int{0, 0, 0, 0}
	num_moves := 100
	for scanner.Scan() {
		line := scanner.Text()
		var ax, ay, vx, vy int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &ax, &ay, &vx, &vy)
		fmt.Println(ax, ay, vx, vy)
		nx := (ax + (num_moves*(vx+mx)%mx)%mx + mx) % mx
		ny := (ay + (num_moves*(vy+my)%my)%my + my) % my
		if nx < mx/2 && ny < my/2 {
			quads[0] += 1
		} else if nx < mx/2 && ny > my/2 {
			quads[1] += 1
		} else if nx > mx/2 && ny < my/2 {
			quads[2] += 1
		} else if nx > mx/2 && ny > my/2 {
			quads[3] += 1
		}
		fmt.Println("new", nx, ny)
	}
	fmt.Println(quads)
	ans := quads[0] * quads[1] * quads[2] * quads[3]
	fmt.Println(ans)

}
