package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("plutonian_pebbles.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str_line := strings.Split(scanner.Text(), " ")
		line := []int{}
		for _, str := range str_line {
			integer_value, _ := strconv.Atoi(str)
			line = append(line, integer_value)
		}
		for i := 0; i < 25; i++ {
			var new_line []int
			for _, j := range line {
				num_as_string := strconv.Itoa(j)
				if j == 0 {
					new_line = append(new_line, 1)
				} else if len(num_as_string)%2 == 0 {
					first_conv, _ := strconv.Atoi(num_as_string[0 : len(num_as_string)/2])
					second_conv, _ := strconv.Atoi(num_as_string[len(num_as_string)/2:])
					new_line = append(new_line, first_conv)
					new_line = append(new_line, second_conv)
				} else {
					new_line = append(new_line, j*2024)
				}
			}
			line = new_line
		}
		fmt.Println(line)
		fmt.Println(len(line))
	}

}
