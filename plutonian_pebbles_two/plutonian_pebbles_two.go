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
		count_map := make(map[int]int)
		for _, str := range str_line {
			integer_value, _ := strconv.Atoi(str)
			line = append(line, integer_value)
			count_map[integer_value] += 1
		}
		for i := 0; i < 75; i++ {
			new_map := make(map[int]int)
			for j, cnt := range count_map {
				num_as_string := strconv.Itoa(j)
				if j == 0 {
					new_map[1] += cnt
				} else if len(num_as_string)%2 == 0 {
					first_conv, _ := strconv.Atoi(num_as_string[0 : len(num_as_string)/2])
					second_conv, _ := strconv.Atoi(num_as_string[len(num_as_string)/2:])
					new_map[first_conv] += cnt
					new_map[second_conv] += cnt
				} else {
					new_map[2024*j] += cnt
				}
			}
			count_map = new_map
		}
		final_ans := 0
		for _, v := range count_map {
			final_ans += v
		}
		fmt.Println(final_ans)
	}

}
