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
	file, err := os.Open("mull_it_over.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var final_ans int = 0
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line)-5; i++ {
			if line[i:i+3] == "mul" && line[i+3] == '(' {
				numbers_string := ""
				found_end := false
				for j := i + 4; j < len(line) && j < i+15; j++ {
					if line[j] == ')' {
						found_end = true
						break
					}
					numbers_string += string(line[j])
				}
				fmt.Println(numbers_string, found_end)
				if !found_end {
					continue
				}
				fmt.Println(numbers_string)
				numbers := strings.Split(numbers_string, ",")
				fmt.Println(numbers)
				if len(numbers) != 2 {
					continue
				}
				numbers_1, err := strconv.Atoi(numbers[0])
				if err != nil {
					continue
				}
				numbers_2, err := strconv.Atoi(numbers[1])
				if err != nil {
					continue
				}
				final_ans += numbers_1 * numbers_2
				i += 3
			}
		}
	}
	fmt.Print(final_ans)
}
