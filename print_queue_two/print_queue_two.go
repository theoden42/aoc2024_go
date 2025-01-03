package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("print_queue.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	edges := make(map[[2]string]bool)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		numbers := strings.Split(line, "|")
		edges[[2]string{numbers[0], numbers[1]}] = true
	}
	final_ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		queue := strings.Split(line, ",")
		correct_line_flag := true
		for i := 0; i < len(queue)-1; i++ {
			for j := i + 1; j < len(queue); j++ {
				_, exists := edges[[2]string{queue[j], queue[i]}]
				if exists {
					correct_line_flag = false
					// we can swap the values
					// we also assume that a valid sorting always exits
					queue[i], queue[j] = queue[j], queue[i]
				}
			}
		}
		if correct_line_flag {
			continue
		}
		number, err := strconv.Atoi(queue[len(queue)/2])
		if err != nil {
			log.Fatal(err)
			return
		}
		final_ans += number
	}
	fmt.Println(final_ans)
}
