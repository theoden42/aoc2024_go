package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("claw_contraption.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	final_ans := int64(0)
	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()
		scanner.Scan()
		empty := scanner.Text()
		var ax, ay, bx, by, px, py int64
		fmt.Println(line1, line2, line3)
		fmt.Sscanf(line1, "Button A: X+%d, Y+%d\n", &ax, &ay)
		fmt.Sscanf(line2, "Button B: X+%d, Y+%d\n", &bx, &by)
		fmt.Sscanf(line3, "Prize: X=%d, Y=%d\n", &px, &py)

		deter1 := ax*by - ay*bx
		deter2 := px*by - py*bx
		deter3 := ax*py - ay*px
		// solve the system of equation
		// x * ax + y * ay = px and x * bx + y * by = py
		// solve using creamer's rule
		if deter1 == 0 {
			// infinite solutions or no solutions, does not occur for this input
			if px%ax == 0 && py%by == 0 {
				final_ans = min(px/ax+py/ay, px/bx+py/by)
			}
		} else if deter2%deter1 == 0 && deter3%deter1 == 0 {
			x := deter2 / deter1
			y := deter3 / deter1
			if x >= 0 && y >= 0 && x <= 100 && y <= 100 {
				final_ans += 3*x + y
			}
		}

		if len(empty) != 0 {
			break
		}
	}
	fmt.Println(final_ans)

}
